// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: memezis.proto

package memezis

import (
	"context"
	"time"

	desc "github.com/cherya/memezis/pkg/memezis"

	"github.com/pkg/errors"
)

func (i *Memezis) GetQueueInfo(ctx context.Context, req *desc.GetQueueInfoRequest) (*desc.GetQueueInfoResponse, error) {
	queueName := req.GetQueue()
	length, err := i.qm.QueueLength(queueName)
	if err != nil {
		return nil, errors.New("can't get queue length")
	}

	lastPostTime, err := i.qm.QueueLastPopTime(queueName)
	if err != nil {
		return nil, errors.New("can't get queue last job time")
	}
	dueTime := lastPostTime

	if length > 0 {
		timeout, err := i.qm.GetQueueTimeout(queueName)
		if err != nil {
			return nil, errors.New("can't get queue timeout")
		}

		dueTime = dueTime.Add(timeout * time.Duration(length))
	}

	return &desc.GetQueueInfoResponse{
		Length:       length,
		DueTime:      toProtoTime(dueTime),
		LastPostTime: toProtoTime(lastPostTime),
	}, nil
}
