package memezis

import (
	"context"
	"encoding/json"
	"github.com/cherya/memezis/internal/store"
	e "github.com/cherya/memezis/pkg/errors"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type upVoteRequest struct {
	UserID string `json:"user_id"`
}

type voteResponse struct {
	Up     int    `json:"up"`
	Down   int    `json:"down"`
	Status string `json:"status,omitempty"`
}

func (m *memezis) UpVote(ctx context.Context, req *http.Request) (interface{}, error) {
	postID, err := strconv.Atoi(mux.Vars(req)["post_id"])
	if err != nil {
		err = errors.Wrapf(err, "UpVote: wrong ")
		return nil, e.WrapMC(err, "UpVote: wrong", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(req.Body)
	var r upVoteRequest
	err = decoder.Decode(&r)
	if err != nil {
		return nil, e.WrapMC(err, "UpVote: can't decode request body", http.StatusBadRequest)
	}

	vote, err := m.store.UpVote(ctx, postID, r.UserID)
	if err != nil {
		if err == store.ErrNotFound {
			return nil, e.WrapC(err, http.StatusNotFound)
		}
		return nil, e.WrapMC(err, "UpVote: can't save vote", http.StatusInternalServerError)
	}

	status, err := m.onPostVote(ctx, postID, *vote)
	if err != nil {
		return nil, err
	}

	return &voteResponse{
		Up:     vote.Up,
		Down:   vote.Down,
		Status: string(status),
	}, nil
}

func (m *memezis) DownVote(ctx context.Context, req *http.Request) (interface{}, error) {
	postID, err := strconv.Atoi(mux.Vars(req)["post_id"])
	if err != nil {
		err = errors.Wrapf(err, "DownVote: invalid post_id (%d)")
		return nil, e.WrapC(err, http.StatusBadRequest)
	}

	decoder := json.NewDecoder(req.Body)
	var r upVoteRequest
	err = decoder.Decode(&r)
	if err != nil {
		return nil, e.WrapMC(err, "DownVote: can't decode request body", http.StatusBadRequest)
	}

	vote, err := m.store.DownVote(ctx, postID, r.UserID)
	if err != nil {
		if err == store.ErrNotFound {
			return nil, e.WrapC(err, http.StatusNotFound)
		}
		return nil, e.WrapMC(err, "DownVote: can't save vote", http.StatusInternalServerError)
	}

	status, err := m.onPostVote(ctx, postID, *vote)
	if err != nil {
		return nil, err
	}

	return &voteResponse{
		Up:     vote.Up,
		Down:   vote.Down,
		Status: string(status),
	}, nil
}
