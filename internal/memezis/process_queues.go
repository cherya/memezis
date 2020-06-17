package memezis

import (
	"github.com/cherya/memezis/internal/store"
	"github.com/cherya/memezis/pkg/queue"
	"github.com/pkg/errors"
)

func (m *memezis) publishToQueues(post *store.Post) error {

	err := m.qm.Push(queue.EverythingQueue, post.ID)
	if err != nil {
		return errors.Wrapf(err, "publishToQueues: can't add post to queue (queue=%s)", queue.EverythingQueue)
	}

	//if post.Source == SourceMemezisBot {
	//	err := m.qm.Push(queue.MemezisSuggestions, post.ID)
	//	if err != nil {
	//		return errors.Wrapf(err, "publishToQueues: can't add post to queue (queue=%s)", queue.MemezisSuggestions)
	//	}
	//}

	return nil
}
