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

func (s *Store) AddPost(
	ctx context.Context,
	media []*Media,
	postTags []string,
	originalCreatedAt time.Time,
	source, submittedBy, text, sourceURL string) (*Post, error) {
	tagsIDs := make([]int64, 0)

	if len(postTags) != 0 {
		tags := make([]*Tag, len(postTags))
		for i, t := range postTags {
			tags[i] = &Tag{Text: t}
		}
		_, err := s.db.NamedExecContext(ctx, `INSERT INTO tags(text) VALUES(:text) ON CONFLICT DO NOTHING`, tags)
		if err != nil {
			return nil, errors.Wrap(err, "store.AddPost: can't add tags")
		}

		err = s.db.SelectContext(ctx, &tagsIDs, `SELECT id FROM tags WHERE text = ANY($1)`, pq.Array(postTags))
		if err != nil {
			return nil, errors.Wrap(err, "store.AddPost: can't save tags")
		}
	}
	hasMedia := len(media) > 0
	post := new(Post)
	err := s.db.GetContext(ctx, post, `
		INSERT INTO posts(source, submitted_by, text, original_created_at, tags, has_media, source_url) 
				    VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		source, submittedBy, text, originalCreatedAt, pq.Array(&tagsIDs), hasMedia, sourceURL)
	if err != nil {
		return nil, errors.Wrap(err, "store.AddPost: can't add post")
	}

	for _, m := range media {
		m.PostID = post.ID
	}

	_, err = s.db.NamedExecContext(ctx, `INSERT INTO media(url, post_id, type, source_id, phash) VALUES(:url, :post_id, :type, :source_id, :phash)`, media)
	if err != nil {
		return nil, errors.Wrap(err, "store.AddPost: can't save media")
	}

	return post, nil
}

func (s *Store) GetPostByID(ctx context.Context, postID int64) (*Post, error) {
	var post Post
	err := s.db.GetContext(ctx, &post, `
		SELECT 
       		id, source, submitted_by, text, tags, created_at, original_created_at, has_media, source_url, source
        FROM posts WHERE id = $1
       `, postID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, errors.Wrapf(err, "GetPostByID: can't get post (postID=%d)", postID)
	}

	err = s.db.GetContext(ctx, &post.Votes, `SELECT * FROM votes_count WHERE post_id = $1`, postID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrapf(err, "GetPostByID: can't get post votes (postID=%d)", postID)
		}
	}

	err = s.db.SelectContext(ctx, &post.Media, `SELECT * FROM media WHERE post_id = $1`, postID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrapf(err, "GetPostByID: can't get post media (postID=%d)", postID)
		}
	}

	err = s.db.SelectContext(ctx, &post.Publish, `SELECT * FROM publish WHERE post_id = $1`, postID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrapf(err, "GetPostByID: can't get post publish (postID=%d)", postID)
		}
	}

	return &post, nil
}

func (s *Store) GetTagsByIDs(ctx context.Context, tagsIDs []int64) ([]string, error) {
	tags := make([]string, 0, len(tagsIDs))
	err := s.db.SelectContext(ctx, &tags, `SELECT text FROM tags WHERE id = ANY($1)`, pq.Array(tagsIDs))
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, errors.Wrapf(err, "store.GetTagsByIDs: can't get tags (ids=%v)", tagsIDs)
		}
	}
	return tags, nil
}

func (s *Store) vote(ctx context.Context, postID int64, userID string, isUp bool) (*VotesCount, error) {
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
			INSERT INTO votes(post_id, user_id, is_up) VALUES ($1, $2, $3) 
			ON CONFLICT (post_id, user_id) 
			DO UPDATE SET is_up = $3 WHERE votes.post_id = $1 AND votes.user_id = $2`, postID, userID, isUp)
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

func (s *Store) UpVote(ctx context.Context, postID int64, userID string) (*VotesCount, error) {
	votes, err := s.vote(ctx, postID, userID, true)
	if err != nil {
		return nil, errors.Wrap(err, "store.UpVote: can't add vote")
	}

	return votes, nil
}

func (s *Store) DownVote(ctx context.Context, postID int64, userID string) (*VotesCount, error) {
	votes, err := s.vote(ctx, postID, userID, false)
	if err != nil {
		return nil, errors.Wrap(err, "store.DownVote: can't add vote")
	}

	return votes, nil
}

func (s *Store) GetRandomPost(ctx context.Context) (*Post, error) {
	var postID int64
	err := s.db.GetContext(ctx, &postID,
		`
			SELECT p.id
			FROM posts p
					 JOIN media m ON p.id = m.post_id
			WHERE m.type = 'photo'
			ORDER BY random()
			LIMIT 1;
		`)
	if err != nil {
		return nil, errors.Wrap(err, "store.GetRandomPost: can't get post")
	}

	post, err := s.GetPostByID(ctx, postID)
	if err != nil {
		return nil, errors.Wrap(err, "store.GetRandomPost: can't get post")
	}

	return post, nil
}
// EnqueuePost mark post as enqueued
// TODO: maybe publishedAt is not necessary
func (s *Store) EnqueuePost(ctx context.Context, postID int64, publishedAt time.Time, to string) error {
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO publish(post_id, status, published_at, published_to) VALUES ($1, $2, $3, $4)`,
		postID, PublishStatusEnqueued, publishedAt, to)

	if err != nil {
		return errors.Wrap(err, "store.EnqueuePost: can't save post publish status")
	}

	return nil
}

// PublishPost mark post as published
func (s *Store) PublishPost(ctx context.Context, postID int64, publishedAt time.Time, to, url string) error {
	res, err := s.db.ExecContext(ctx,
		`UPDATE publish SET status=$1, published_at=$2, url=$3 WHERE published_to=$4 AND post_id = $5`,
		PublishStatusPublished, publishedAt, url, to, postID)

	if err != nil {
		return errors.Wrap(err, "store.PublishPost: can't update publish status")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "store.PublishPost: can't get rows affected")
	}
	if rowsAffected != 0 {
		return nil
	}

	_, err = s.db.ExecContext(ctx,
		`INSERT INTO publish(post_id, status, published_at, published_to, url) VALUES ($1, $2, $3, $4, $5)`,
		postID, PublishStatusPublished, publishedAt, to, url)

	if err != nil {
		return errors.Wrap(err, "store.PublishPost: can't save post publish status")
	}

	return nil
}

func (s *Store) GetPostsByMediaHashes(ctx context.Context, hashes []string) ([]Post, error) {
	posts := make([]Post, 0)
	err := s.db.SelectContext(ctx, &posts, `
		SELECT 
       		id, source, submitted_by, text, tags, created_at, original_created_at, has_media, source_url, source
        FROM posts WHERE id IN (
            SELECT post_id FROM media WHERE phash = ANY($1) 
		)
    `, pq.Array(hashes))
	if err != nil {
		return nil, errors.Wrap(err, "store.GetPostsByMediaHashes: can't select posts")
	}

	return posts, nil
}

func (s *Store) GetHashes(ctx context.Context) ([]string, error) {
	hashes := make([]string, 0)
	err := s.db.SelectContext(ctx, &hashes, `SELECT phash FROM media WHERE phash != ''`)
	if err != nil {
		return nil, errors.Wrap(err, "GetHashes: can't select hashes")
	}
	return hashes, nil
}

func (s *Store) GetMediaByIDs(ctx context.Context, ids []int64) ([]Media, error) {
	media := make([]Media, 0)
	err := s.db.SelectContext(ctx, &media, `SELECT * FROM media WHERE id = ANY($1)`, pq.Array(ids))
	if err != nil {
		return nil, errors.Wrap(err, "GetMediaByIDs: can't select media")
	}
	return media, nil
}