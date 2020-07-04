package queue

import (
	"context"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

var ErrEmptyQueue = errors.New("queue is empty")

type Manager struct {
	redis *redis.Pool
	ns    string
}

func NewManager(pool *redis.Pool, namespace string) *Manager {
	return &Manager{
		redis: pool,
		ns: namespace,
	}
}

func (m *Manager) Push(queueName string, value string) error {
	conn := m.redis.Get()
	defer conn.Close()

	_, err := conn.Do("ZADD", m.queueKey(queueName), time.Now().Unix() * 1000, value)
	if err != nil {
		return errors.Wrap(err, "QueueManager.SetQueueTimeout: can't push value")
	}

	return nil
}

func (m *Manager) Pop(queueName string) (string, error) {
	conn := m.redis.Get()
	defer conn.Close()

	reply, err := redis.Strings(conn.Do("ZPOPMIN", m.queueKey(queueName)))
	if err != nil {
		return "", errors.Wrap(err, "QueueManager.Pop: can't pop value")
	}
	if len(reply) == 0 {
		return "", ErrEmptyQueue
	}

	err = m.updateQueueLastPopTime(queueName, time.Now().UTC())
	if err != nil {
		return "", errors.Wrap(err, "QueueManager.Pop: can't update pop time")
	}

	return reply[0], nil
}

func (m *Manager) QueueLength(queueName string) (int64, error) {
	conn := m.redis.Get()
	defer conn.Close()
	l, err := redis.Int64(conn.Do("ZCARD", m.queueKey(queueName)))
	if err != nil {
		return 0, errors.Wrap(err, "QueueManager.QueueLength: can't get queue length")
	}
	return l, nil
}

func (m *Manager) QueueLastPopTime(queue string) (time.Time, error) {
	return m.getQueueLastPopTime(queue, time.Now().UTC())
}

func (m *Manager) Consume(ctx context.Context, queue string, pollTimeout time.Duration, handler func(value string)) {
	ticker := time.NewTicker(pollTimeout)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				lt, err := m.getQueueLastPopTime(queue, time.Now().UTC())
				if err != nil {
					log.Error("Consume: can't get queue last pop time")
					break
				}
				t, err := m.GetQueueTimeout(queue)
				if err != nil {
					log.Error("Consume: can't get queue timeout")
					break
				}
				if time.Now().UTC().After(lt.UTC().Add(t)) {
					val, err := m.Pop(queue)
					if err != nil {
						if errors.Cause(err) != ErrEmptyQueue {
							log.Error("Consume: can't pop elemet")
						}
						break
					}
					handler(val)
				}
			}
		}
	}(ctx)
}

func (m *Manager) SetQueueTimeout(queue string, d time.Duration) error {
	conn := m.redis.Get()
	defer conn.Close()

	_, err := conn.Do("SET", m.timeoutKey(queue), d.Seconds())
	if err != nil {
		return errors.Wrap(err, "QueueManager.SetQueueTimeout: can't set timeout")
	}
	return nil
}

func (m *Manager) GetQueueTimeout(queue string) (time.Duration, error) {
	conn := m.redis.Get()
	defer conn.Close()

	t, err := redis.Int64(conn.Do("GET", m.timeoutKey(queue)))
	if err != nil {
		if errors.Cause(err) == redis.ErrNil {
			return time.Minute, nil
		}
		return 0, errors.Wrap(err, "QueueManager.GetQueueTimeout: can't get timeout")
	}

	return time.Second * time.Duration(t), nil
}

func (m *Manager) updateQueueLastPopTime(queue string, t time.Time) error {
	conn := m.redis.Get()
	defer conn.Close()

	_, err := conn.Do("SET", m.lastOperationTimestampKey(queue), t.Unix())
	if err != nil {
		return errors.Wrap(err, "QueueManager.updateQueueLastPopTime: can't set handle time")
	}
	return nil
}

func (m *Manager) getQueueLastPopTime(queue string, dflt time.Time) (time.Time, error) {
	conn := m.redis.Get()
	defer conn.Close()

	_ = conn.Send("MULTI")
	_ = conn.Send("SETNX", m.lastOperationTimestampKey(queue), dflt.Unix())
	_ = conn.Send("GET", m.lastOperationTimestampKey(queue))
	res, err := redis.Int64s(conn.Do("EXEC"))
	if err != nil {
		return time.Time{}, errors.Wrap(err, "can't get handle time")
	}
	if len(res) != 2 {
		return time.Time{}, errors.Wrap(err, "unexpected answer length")
	}

	return time.Unix(res[1], 0).UTC(), nil
}

// KEYS

func (m *Manager) queueKey(queue string) string {
	return fmt.Sprintf("%s:%s", m.ns, queue)
}

func (m *Manager) timeoutKey(queue string) string {
	return fmt.Sprintf("%s:%s:timeout", m.ns, queue)
}

func (m *Manager) lastOperationTimestampKey(queue string) string {
	return fmt.Sprintf("%s:%s:last_pop_time", m.ns, queue)
}

func (m *Manager) queueJobsListKey(queue string) string {
	return fmt.Sprintf("%s:jobs:%s", m.ns, queue)
}

func (m *Manager) queueLockedItemsKey(queue string) string {
	return fmt.Sprintf("%s:jobs:%s:lock", m.ns, queue)
}