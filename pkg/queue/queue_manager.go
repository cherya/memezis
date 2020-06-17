package queue

import (
	"fmt"
	"time"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

const (
	EverythingQueue  = "everything"

	MemezisNamespace = "memezis"
)

type Manager struct {
	redis *redis.Pool
}

func NewManager(pool *redis.Pool) *Manager {
	return &Manager{
		redis: pool,
	}
}

func (m *Manager) Push(queue string, postID int64) error {
	enqueuer := work.NewEnqueuer(MemezisNamespace, m.redis)
	_, err := enqueuer.Enqueue(queue, work.Q{"postID": postID})
	if err != nil {
		return errors.Wrap(err, "QueueManager.Push: can't enqueue job")
	}
	return nil
}

func (m *Manager) PushWithDelay(queue string, delay time.Duration, postID int64) error {
	enqueuer := work.NewEnqueuer(MemezisNamespace, m.redis)
	_, err := enqueuer.EnqueueIn(queue, int64(delay.Seconds()), work.Q{"postID": postID})
	if err != nil {
		return errors.Wrap(err, "QueueManager.Push: can't enqueue job")
	}
	return nil
}

func (m *Manager) QueueLength(queue string) (int64, error) {
	conn := m.redis.Get()
	defer conn.Close()

	llen, err := redis.Int64(conn.Do("LLEN", queueJobsListKey(queue)))
	if err != nil {
		return 0, errors.Wrap(err, "QueueManager.QueueLength: can't get queue length")
	}

	locked, err := redis.Int64(conn.Do("GET", queueLockedItemsKey(queue)))
	if err != nil {
		return 0, errors.Wrap(err, "QueueManager.QueueLength: can't get locked items count")
	}

	return llen + locked, nil
}

func (m *Manager) QueueLastJobTime(queue string) (time.Time, error) {
	return m.getQueueLastHandleTime(queue)
}

type Context struct{}

func (m *Manager) ConsumeWithDelay(queue string, handler func(job *work.Job) error) {
	pool := work.NewWorkerPool(Context{}, 2, MemezisNamespace, m.redis)
	ticker := time.NewTicker(time.Second * 10)

	pool.JobWithOptions(queue, work.JobOptions{MaxConcurrency: 1}, func(job *work.Job) error {
		for range ticker.C {
			t, err := m.getQueueLastHandleTime(queue)
			if err != nil {
				return errors.Wrap(err, "ConsumeWithDelay: can't get queue last handle time")
			}
			timeout, err := m.GetQueueTimeout(queue)
			if err != nil {
				return errors.Wrap(err, "ConsumeWithDelay: can't get queue timeout")
			}
			if t.UTC().Add(timeout).Before(time.Now().UTC()) {
				err = handler(job)
				if err == nil {
					rerr := m.updateQueueLastHandleTime(queue)
					if rerr != nil {
						return errors.Wrap(rerr, "ConsumeWithDelay: can't set handle time")
					}
				}
				return err
			}
			job.Checkin(fmt.Sprintf("waiting to process %s", time.Now().UTC().Sub(t.UTC()).Truncate(time.Second)))
		}

		return nil
	})

	pool.Start()
}

func (m *Manager) Consume(queue string, handler func(job *work.Job) error) {
	pool := work.NewWorkerPool(Context{}, 10, MemezisNamespace, m.redis)

	pool.JobWithOptions(queue, work.JobOptions{MaxConcurrency: 1}, func(job *work.Job) error {
		err := handler(job)
		if err != nil {
			return errors.Wrap(err, "Consume: error handling job")
		}
		return err
	})

	pool.Start()
}

func timeoutKey(queue string) string {
	return fmt.Sprintf("%s:%s:timeout", MemezisNamespace, queue)
}

func lastOperationTimestampKey(queue string) string {
	return fmt.Sprintf("%s:%s:last_operation_timeout", MemezisNamespace, queue)
}

func queueJobsListKey(queue string) string {
	return fmt.Sprintf("%s:jobs:%s", MemezisNamespace, queue)
}

func queueLockedItemsKey(queue string) string {
	return fmt.Sprintf("%s:jobs:%s:lock", MemezisNamespace, queue)
}

func (m *Manager) SetQueueTimeout(queue string, d time.Duration) error {
	conn := m.redis.Get()
	defer conn.Close()

	_, err := conn.Do("SET", timeoutKey(queue), d.Seconds())
	if err != nil {
		return errors.Wrap(err, "QueueManager.SetQueueTimeout: can't set timeout")
	}
	return nil
}

func (m *Manager) GetQueueTimeout(queue string) (time.Duration, error) {
	conn := m.redis.Get()
	defer conn.Close()

	t, err := redis.Int64(conn.Do("GET", timeoutKey(queue)))
	if err != nil {
		if errors.Cause(err) == redis.ErrNil {
			return time.Minute, nil
		}
		return 0, errors.Wrap(err, "QueueManager.GetQueueTimeout: can't get timeout")
	}

	return time.Second * time.Duration(t), nil
}

func (m *Manager) updateQueueLastHandleTime(queue string) error {
	conn := m.redis.Get()
	defer conn.Close()

	_, err := conn.Do("SET", lastOperationTimestampKey(queue), time.Now().UTC().Unix())
	if err != nil {
		return errors.Wrap(err, "QueueManager.SetQueueTimeout: can't set handle time")
	}
	return nil
}

func (m *Manager) getQueueLastHandleTime(queue string) (time.Time, error) {
	conn := m.redis.Get()
	defer conn.Close()

	res, err := redis.Int64(conn.Do("GET", lastOperationTimestampKey(queue)))
	if err != nil {
		if errors.Cause(err) == redis.ErrNil {
			t := time.Now().UTC()
			_, err = conn.Do("SETNX", lastOperationTimestampKey(queue), t.Unix())
			if err != nil {
				return time.Time{}, errors.Wrap(err, "QueueManager.SetQueueTimeout: can't set handle time")
			}
			return t, nil
		}
		return time.Time{}, errors.Wrap(err, "QueueManager.SetQueueTimeout: can't get handle time")
	}

	return time.Unix(res, 0), nil
}
