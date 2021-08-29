package memezis

import (
	"context"
	"strconv"
	"time"

	"github.com/cherya/memezis/internal/app/store"

	"github.com/pkg/errors"
)

// sources
//TODO: make configurable
const (
	SourceMemezisBot = "memezis_bot"
)

// publish channels
//TODO: make configurable
const (
	ChannelShaurmemes = "shaurmemes"
)

//TODO: make configurable
const votesToEnqueue = 3

func (i *Memezis) VotePost(ctx context.Context, postID int64, vote store.VotesCount) (store.PublishStatus, error) {
	post, err := i.store.GetPostByID(ctx, postID)
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			return store.PublishStatusUnknown, errors.Wrap(err, "post not found")
		}
		return store.PublishStatusUnknown, errors.Wrap(err, "VotePost: can't get post from store")
	}

	if post.Source == SourceMemezisBot {
		if vote.Up >= votesToEnqueue {
			err = i.enqueuePost(ctx, postID, time.Now(), ChannelShaurmemes)
			if err != nil {
				return store.PublishStatusUnknown, errors.Wrap(err, "VotePost: can't publish post")
			}
			return store.PublishStatusEnqueued, nil
		}
		if vote.Down >= votesToEnqueue {
			return store.PublishStatusDeclined, nil
		}
	}

	return store.PublishStatusUnknown, nil
}

var sourceToQueue = map[string]string{
	"shaurmemes": "shaurmemes",
}

func (i *Memezis) enqueuePost(ctx context.Context, postID int64, at time.Time, to string) error {
	err := i.store.EnqueuePost(ctx, postID, at, to)
	if err != nil {
		return errors.Wrap(err, "publishPost: can't publish post")
	}

	if q, ok := sourceToQueue[to]; ok {
		err = i.qm.Push(q, strconv.FormatInt(postID, 10))
		if err != nil {
			return errors.Wrap(err, "publishPost: can't publish post to queue")
		}
	}

	return nil
}
