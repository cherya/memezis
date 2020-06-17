package memezis

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/cherya/memezis/internal/store"
	e "github.com/cherya/memezis/pkg/errors"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type addPostRequest struct {
	Media     []mediaRequest `json:"media"`
	AddedBy   string         `json:"added_by"`
	Text      string         `json:"text"`
	Tags      []string       `json:"tags"`
	CreatedAt time.Time      `json:"created_at"`
}

type mediaRequest struct {
	URL      string `json:"url"`
	Type     string `json:"type"`
	SourceID string `json:"source_id"`
	SHA1     string `json:"sha_1"`
}

type addPostResponse struct {
	ID         int64   `json:"id"`
	Duplicates []int64 `json:"duplicates"`
}

func (m *memezis) AddPost(ctx context.Context, req *http.Request) (interface{}, error) {
	decoder := json.NewDecoder(req.Body)
	var r addPostRequest
	err := decoder.Decode(&r)
	if err != nil {
		return nil, e.WrapMC(err, "AddPost: can't decode request body", http.StatusBadRequest)
	}

	wg := sync.WaitGroup{}
	errs := make(chan error)
	wgDone := make(chan bool)


	media := make([]*store.Media, 0, len(r.Media))
	for _, u := range r.Media {
		if u.URL == "" && u.SourceID == "" {
			return nil, e.WrapC(errors.Errorf("AddPost: empty media", u), http.StatusBadRequest)
		}
		if u.URL != "" {
			wg.Add(1)
			go func(mr mediaRequest) {
				if !m.fs.IsTempObjExists(u.URL) {
					errs <- e.WrapC(errors.Errorf("AddPost: object %s does not exists", u), http.StatusBadRequest)
				}
				wg.Done()
			}(u)
		}

		media = append(media, &store.Media{
			Key:      u.URL,
			Type:     u.Type,
			SourceID: u.SourceID,
			SHA1:     u.SHA1,
		})
	}

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		// carry on
		break
	case err := <-errs:
		close(errs)
		return nil, err
	}

	client := req.Context().Value("client").(*Client)
	post, err := m.store.AddPost(ctx, media, r.Tags, r.CreatedAt, client.Name, r.AddedBy, r.Text)
	if err != nil {
		return nil, e.WrapMC(err, "AddPost: can't save post to store", http.StatusInternalServerError)
	}


	wgDone = make(chan bool)

	for _, u := range r.Media {
		if u.URL != "" {
			wg.Add(1)
			go func(url string) {
				err := m.fs.MakeObjPermanent(u.URL)
				if err != nil {
					err = errors.Wrapf(err, "AddPost: can't make object permanent (key=%s)", m)
					errs <- e.WrapC(err, http.StatusInternalServerError)
				}
				wg.Done()
			}(u.URL)
		}
	}

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		// carry on
		break
	case err := <-errs:
		close(errs)
		return nil, err
	}

	err = m.publishToQueues(post)
	if err != nil {
		log.Println("AddPost: can't add posts to queues", err)
	}

	duplicates, err := m.store.CheckDuplicate(ctx, post.ID)
	if err != nil {
		log.Println("AddPost: can't check duplicates", err)
	}
	return &addPostResponse{
		ID:         post.ID,
		Duplicates: duplicates,
	}, nil
}

type mediaResponse struct {
	URL      string `json:"url"`
	Type     string `json:"type"`
	SourceID string `json:"source_id"`
}

type getPostResponse struct {
	ID      int64            `json:"id"`
	Media   []*mediaResponse `json:"media"`
	AddedBy string           `json:"added_by"`
	Source  string           `json:"source"`
	Text    string           `json:"text"`
	Votes   *voteResponse    `json:"votes"`
	Tags    []string         `json:"tags"`
}

func (m *memezis) GetPostByID(ctx context.Context, req *http.Request) (interface{}, error) {
	postID, err := strconv.Atoi(mux.Vars(req)["post_id"])
	if err != nil {
		err = errors.Wrapf(err, "GetPostByID: invalid post_id (%d)", postID)
		return nil, e.WrapC(err, http.StatusBadRequest)
	}
	post, err := m.store.GetPostByID(ctx, int64(postID))
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			return nil, e.WrapC(err, http.StatusNotFound)
		}
		log.Println("GetPostByID: can't get post from store", err)
		return nil, errors.Wrap(err, "GetPostByID: can't get post from store")
	}

	tags, err := m.store.GetTagsByIDs(ctx, post.Tags)
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			tags = []string{}
		}
		log.Println("GetPostByID: can't get tags from store", err)
		return nil, errors.Wrap(err, "GetPostByID: can't get tags from store")
	}

	resp := m.postToResponse(post, tags)

	return resp, nil
}

func (m *memezis) postToResponse(post *store.Post, tags []string) *getPostResponse {
	votes := &voteResponse{
		Up:   post.Votes.Up,
		Down: post.Votes.Down,
	}

	media := make([]*mediaResponse, len(post.Media))
	for i, mm := range post.Media {
		media[i] = &mediaResponse{
			URL:      m.fs.GetObjAbsoluteURL(mm.Key),
			Type:     mm.Type,
			SourceID: mm.SourceID,
		}
	}

	return &getPostResponse{
		ID:      post.ID,
		Media:   media,
		Votes:   votes,
		AddedBy: post.SubmittedBy,
		Source:  post.Source,
		Text:    post.Text,
		Tags:    tags,
	}
}

func (m *memezis) GetRandomPost(ctx context.Context, req *http.Request) (interface{}, error) {
	post, err := m.store.GetRandomPost(ctx)
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			return nil, e.WrapC(err, http.StatusNotFound)
		}
		log.Println("GetRandomPost: can't get post from store", err)
		return nil, errors.Wrap(err, "GetRandomPost: can't get post from store")
	}

	tags, err := m.store.GetTagsByIDs(ctx, post.Tags)
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			tags = []string{}
		}
		log.Println("GetPostByID: can't get tags from store", err)
		return nil, errors.Wrap(err, "GetPostByID: can't get tags from store")
	}

	resp := m.postToResponse(post, tags)

	return resp, nil
}

type publishPostRequest struct {
	PublishedAt int64  `json:"published_at"`
	PublishedTo string `json:"published_to"`
}

func (m *memezis) PublishPost(ctx context.Context, req *http.Request) (interface{}, error) {
	postID, err := strconv.Atoi(mux.Vars(req)["post_id"])
	if err != nil {
		err = errors.Wrapf(err, "EnqueuePost: invalid post_id (%d)", postID)
		return nil, e.WrapC(err, http.StatusBadRequest)
	}

	decoder := json.NewDecoder(req.Body)
	var r publishPostRequest
	err = decoder.Decode(&r)
	if err != nil {
		return nil, e.WrapMC(err, "EnqueuePost: can't decode request body", http.StatusBadRequest)
	}

	err = m.enqueuePost(ctx, int64(postID), time.Unix(r.PublishedAt, 0), r.PublishedTo)
	if err != nil {
		return nil, errors.Wrap(err, "EnqueuePost: can't publish post")
	}

	return struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}, nil
}

var sourceToQueue = map[string]string{
	"shaurmemes": "shaurmemes",
}

func (m *memezis) enqueuePost(ctx context.Context, postID int64, at time.Time, to string) error {
	err := m.store.EnqueuePost(ctx, postID, at, to)
	if err != nil {
		return errors.Wrap(err, "publishPost: can't publish post")
	}

	if q, ok := sourceToQueue[to]; ok {
		err = m.qm.Push(q, postID)
		if err != nil {
			return errors.Wrap(err, "publishPost: can't publish post to queue")
		}
	}

	return nil
}
