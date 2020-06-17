package memezis

import (
	"context"
	"net/http"
	"time"

	e "github.com/cherya/memezis/pkg/errors"
	"github.com/gorilla/mux"
	"google.golang.org/appengine/log"
)

type queueInfoResponse struct {
	Length       int64     `json:"length"`
	LastPostTime time.Time `json:"last_post_time"`
	DueTime      time.Time `json:"due_time"`
}

func (m *memezis) GetQueueInfo(ctx context.Context, req *http.Request) (interface{}, error) {
	queueName := mux.Vars(req)["queue_name"]
	length, err := m.qm.QueueLength(queueName)
	if err != nil {
		log.Errorf(ctx, "GetQueueInfo: can't get queue length (err=%s)", err)
		return nil, e.NewC("can't get queue length", http.StatusInternalServerError)
	}

	lastPostTime, err := m.qm.QueueLastJobTime(queueName)
	if err != nil {
		log.Errorf(ctx, "GetQueueInfo: can't get queue last job time (err=%s)", err)
		return nil, e.NewC("can't get queue last job time", http.StatusInternalServerError)
	}
	dueTime := lastPostTime

	if length > 0 {
		timeout, err := m.qm.GetQueueTimeout(queueName)
		if err != nil {
			log.Errorf(ctx, "GetQueueInfo: can't get queue timeout (err=%s)", err)
			return nil, e.NewC("can't get queue timeout", http.StatusInternalServerError)
		}

		dueTime = dueTime.Add(timeout * time.Duration(length))
	}


	return &queueInfoResponse{
		Length: length,
		DueTime: dueTime,
		LastPostTime: lastPostTime,
	}, nil
}
