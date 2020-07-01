package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func (s *store) AddPost(
	ctx context.Context,
	media []*Media,
	postTags []string,
	originalCreatedAt time.Time,
	source, submittedBy, text string) (*Post, error) {
	tagsIDs := make([]int64, 0)

	if len(postTags) != 0 {
		tags := make([]*Tag, len(postTags))
		for i, t := range postTags {
			tags[i] = &Tag{Text: t}
		}
		_, err := s.db.NamedExecContext(ctx, `insert into tags(text) values(:text) on conflict do nothing`, tags)
		if err != nil {
			return nil, errors.Wrap(err, "store.AddPost: can't add tags")
		}

		err = s.db.SelectContext(ctx, &tagsIDs, `select id from tags where text = any($1)`, pq.Array(postTags))
		if err != nil {
			return nil, errors.Wrap(err, "store.AddPost: can't save tags")
		}
	}
	hasMedia := len(media) > 0
	post := new(Post)
	err := s.db.GetContext(ctx, post, `insert into posts(source, submitted_by, text, original_created_at, tags, has_media) values($1, $2, $3, $4, $5, $6) returning *`,
		source, submittedBy, text, originalCreatedAt, pq.Array(&tagsIDs), hasMedia)
	if err != nil {
		return nil, errors.Wrap(err, "store.AddPost: can't add post")
	}

	for _, m := range media {
		m.PostID = post.ID
	}

	_, err = s.db.NamedExecContext(ctx, `insert into media(url, post_id, type, source_id, sha1) values(:url, :post_id, :type, :source_id, :sha1)`, media)
	if err != nil {
		return nil, errors.Wrap(err, "store.AddPost: can't save media")
	}

	return post, nil
}

func (s *store) GetPostByID(ctx context.Context, postID int64) (*Post, error) {
	var post Post
	err := s.db.GetContext(ctx, &post, `select 
       id, source, submitted_by, text, tags, created_at, original_created_at, has_media
       from posts where id = $1`, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.Wrapf(err, "GetPostByID: can't get post (postID=%d)", postID)
	}

	err = s.db.GetContext(ctx, &post.Votes, `select * from votes_count where post_id = $1`, postID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrapf(err, "GetPostByID: can't get post votes (postID=%d)", postID)
		}
	}

	err = s.db.SelectContext(ctx, &post.Media, `select * from media where post_id = $1`, postID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrapf(err, "GetPostByID: can't get post media (postID=%d)", postID)
		}
	}

	return &post, nil
}

func (s *store) GetTagsByIDs(ctx context.Context, tagsIDs []int64) ([]string, error) {
	tags := make([]string, 0, len(tagsIDs))
	err := s.db.SelectContext(ctx, &tags, `select text from tags where id = any($1)`, pq.Array(tagsIDs))
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrapf(err, "store.GetTagsByIDs: can't get tags (ids=%v)", tagsIDs)
		}
	}
	return tags, nil
}

func (s *store) vote(ctx context.Context, postID int64, userID string, isUp bool) (*VotesCount, error) {
	votes := new(VotesCount)
	q := `insert into votes_count(post_id, %s) values($1, 1) 
			on conflict(post_id)
			do update set (up, down) = (
				(select count(*) from votes where post_id = $1 and is_up = true),
				(select count(*) from votes where post_id = $1 and is_up = false)
			)
			where votes_count.post_id = $1 
			returning *`

	if isUp {
		q = fmt.Sprintf(q, "up")
	} else {
		q = fmt.Sprintf(q, "down")
	}

	err := WithTransaction(ctx, s.db, func(ctx context.Context, tx *sqlx.Tx) error {
		_, err := tx.ExecContext(ctx, `
			insert into votes(post_id, user_id, is_up) values($1, $2, $3) 
			on conflict(post_id, user_id) 
			do update set is_up = $3 where votes.post_id = $1 and votes.user_id = $2`, postID, userID, isUp)
		if err != nil {
			return errors.Wrapf(err, "can't create vote (postID=%d) (userID=%s)", postID, userID)
		}

		err = tx.GetContext(ctx, votes, q, postID)
		if err != nil {
			return errors.Wrapf(err, "can't count votes (postID=%d, userID=%s)", postID, userID)
		}

		return nil
	})
	return votes, err
}

func (s *store) UpVote(ctx context.Context, postID int64, userID string) (*VotesCount, error) {
	votes, err := s.vote(ctx, postID, userID, true)
	if err != nil {
		return nil, errors.Wrap(err, "store.UpVote: can't add vote")
	}

	return votes, nil
}

func (s *store) DownVote(ctx context.Context, postID int64, userID string) (*VotesCount, error) {
	votes, err := s.vote(ctx, postID, userID, false)
	if err != nil {
		return nil, errors.Wrap(err, "store.DownVote: can't add vote")
	}

	return votes, nil
}

func (s *store) GetRandomPost(ctx context.Context) (*Post, error) {
	var postID int64
	err := s.db.GetContext(ctx, &postID, `SELECT id FROM posts OFFSET floor(random() * (SELECT count(*) FROM posts)) LIMIT 1`)
	if err != nil {
		return nil, errors.Wrap(err, "store.GetRandomPost: can't get posts count")
	}

	post, err := s.GetPostByID(ctx, postID)
	if err != nil {
		return nil, errors.Wrap(err, "store.GetRandomPost: can't get post")
	}

	return post, nil
}

func (s *store) EnqueuePost(ctx context.Context, postID int64, publishedAt time.Time, to string) error {
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO publish(post_id, status, published_at, published_to) VALUES ($1, $2, $3, $4)`,
		postID, PublishStatusEnqueued, publishedAt, to)

	if err != nil {
		return errors.Wrap(err, "store.EnqueuePost: can't save post publish status")
	}

	return nil
}

func (s *store) CheckDuplicate(ctx context.Context, postID int64) ([]int64, error) {
	var media []Media

	err := s.db.SelectContext(ctx, &media, `select * from media where post_id = $1`, postID)
	if err != nil {
		return nil, errors.Wrap(err, "store.CheckDuplicate: can't get post media")
	}

	var sha, sources []string
	for _, m := range media {
		sha = append(sha, m.SHA1)
		sources = append(sources, m.SourceID)
	}
	var posts []int64
	err = s.db.SelectContext(ctx, &posts, `
		SELECT post_id from media WHERE (source_id = ANY($1) OR sha1 = ANY($2)) AND post_id != $3
	`, pq.Array(sources), pq.Array(sha), postID)
	if err != nil {
		return nil, errors.Wrap(err, "store.CheckDuplicate: can't get media")
	}

	return posts, nil
}
