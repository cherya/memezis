package memezis

import (
	"context"
	"net/http"
	"time"

	"github.com/cherya/memezis/internal/store"
	e "github.com/cherya/memezis/pkg/errors"
	"github.com/pkg/errors"
)

// move to config
const votesToEnqueue = 4

func (m *memezis) onPostVote(ctx context.Context, postID int, vote store.VotesCount) (store.PublishStatus, error) {
	post, err := m.store.GetPostByID(ctx, int64(postID))
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			return store.PublishStatusUnknown, e.WrapC(err, http.StatusNotFound)
		}
		return store.PublishStatusUnknown, errors.Wrap(err, "onPostVote: can't get post from store")
	}

	if post.Source == SourceMemezisBot {
		if vote.Up >= votesToEnqueue {
			err = m.enqueuePost(ctx, int64(postID), time.Now(), "shaurmemes")
			if err != nil {
				return store.PublishStatusUnknown, errors.Wrap(err, "onPostVote: can't publish post")
			}
			return store.PublishStatusEnqueued, nil
		}
		if vote.Down >= votesToEnqueue {
			return store.PublishStatusDeclined, nil
		}
	}

	return store.PublishStatusUnknown, nil
}