// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: memezis.proto

package memezis

import (
	"context"

	"github.com/cherya/memezis/internal/app/store"
	desc "github.com/cherya/memezis/pkg/memezis"

	"github.com/pkg/errors"
)

func (i *Memezis) UpVote(ctx context.Context, req *desc.VoteRequest) (*desc.Vote, error) {
	postID := req.GetPostID()

	vote, err := i.store.UpVote(ctx, postID, req.GetUserID())
	if err != nil {
		if errors.Cause(err) == store.ErrNotFound {
			return nil, errors.Wrap(err, "UpVote: post not found")
		}
		if errors.Cause(err) == store.ErrOwnPostVoting {
			return &desc.Vote{
				Accepted: false,
			}, nil
		}
		return nil, errors.Wrap(err, "UpVote: can't save vote")
	}

	status, err := i.VotePost(ctx, postID, *vote)
	if err != nil {
		return nil, err
	}

	return &desc.Vote{
		Up:       int64(vote.Up),
		Down:     int64(vote.Down),
		Status:   string(status),
		Accepted: true,
	}, nil
}
