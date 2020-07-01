package queue

import (
	"fmt"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"

)

const namespace = "QueueTest"

func TestManager_Push(t *testing.T) {
	a := assert.New(t)
	r := getRedisPool()

	m := &Manager{
		redis: r,
		ns:    namespace,
	}

	type args struct {
		queueName string
		values     []string
	}
	tests := []struct {
		name    string
		args    args
		expected []string
		wantErr bool
	}{
		{
			"one value",
			args{
				"queue",
				[]string{"10"},
			},
			[]string{"10"},
			false,
		},
		{
			"several values",
			args{
				"queue",
				[]string{"10", "beer", "heh", "lol"},
			},
			[]string{"10", "beer", "heh", "lol"},
			false,
		},
		{
			"repeated values",
			args{
				"queue",
				[]string{"kek", "kek", "kek", "kek"},
			},
			[]string{"kek"},
			false,
		},
		{
			"no values",
			args{
				"queue",
				[]string{},
			},
			[]string{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.values {
				err := m.Push(tt.args.queueName, v)
				if err != nil {
					if !tt.wantErr {
						t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}
			res := make([]string, 0)
			for range tt.expected {
				v, err := zpopmin(r, fmt.Sprintf("%s:%s", namespace, tt.args.queueName))
				if err != nil {
					t.Errorf("zpopmin() error = %v", err)
				}
				res = append(res, v)
			}
			a.Equal(tt.expected, res, tt.name)
			if err := flush(r, namespace); err != nil {
				t.Fatal("can't flush redis")
			}
		})
	}
}

func TestManager_Pop(t *testing.T) {
	a := assert.New(t)
	r := getRedisPool()

	m := &Manager{
		redis: r,
		ns:    namespace,
	}

	type args struct {
		queueName string
		values []string
	}
	tests := []struct {
		name    string
		args    args
		expected []string
		wantErr bool
	}{
		{
			"one value",
			args{
				"queue",
				[]string{"10"},
			},
			[]string{"10"},
			false,
		},
		{
			"several values",
			args{
				"queue",
				[]string{"10", "kek", "pek"},
			},
			[]string{"10", "kek", "pek"},
			false,
		},
		{
			"no values",
			args{
				"queue",
				[]string{},
			},
			[]string{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, v := range tt.args.values {
				err := zadd(r, fmt.Sprintf("%s:%s", namespace, tt.args.queueName), v, i)
				if err != nil {
					t.Errorf("zadd() error = %v", err)
				}
			}
			res := make([]string, 0)
			for range tt.expected {
				v, err := m.Pop(tt.args.queueName)
				if err != nil {
					if !tt.wantErr {
						t.Errorf("Pop() error = %v", err)
					}
				}
				res = append(res, v)
			}
			a.Equal(tt.expected, res, tt.name)
			if err := flush(r, namespace); err != nil {
				t.Fatal("can't flush redis")
			}
		})
	}
}

func TestManager_QueueLength(t *testing.T) {
	a := assert.New(t)
	r := getRedisPool()

	m := &Manager{
		redis: r,
		ns:    namespace,
	}

	type args struct {
		queueName string
		values []string
	}
	tests := []struct {
		name    string
		args    args
		expected int64
		wantErr bool
	}{
		{
			"one value",
			args{
				"queue",
				[]string{"10"},
			},
			1,
			false,
		},
		{
			"several values",
			args{
				"queue",
				[]string{"10", "kek", "pek"},
			},
			3,
			false,
		},
		{
			"no values",
			args{
				"queue",
				[]string{},
			},
			0,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.values {
				err := m.Push(tt.args.queueName, v)
				if err != nil {
					t.Errorf("Push() error = %v", err)
				}
			}
			l, err := m.QueueLength(tt.args.queueName)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("QueueLength() error = %v", err)
				}
			}

			a.Equal(tt.expected, l, tt.name)
			if err := flush(r, namespace); err != nil {
				t.Fatal("can't flush redis")
			}
		})
	}
}

func TestManager_getQueueLastPopTime(t *testing.T) {
	a := assert.New(t)
	r := getRedisPool()

	m := &Manager{
		redis: r,
		ns:    namespace,
	}

	queue := "queue"

	tName := "initial get"
	t.Run(tName, func(t *testing.T) {
		now := time.Now().UTC()
		tm, err := m.getQueueLastPopTime(queue, now)
		if err != nil {
			t.Errorf("getQueueLastPopTime() error = %v", err)
		}
		a.Equal(tm.Unix(), now.Unix(), tName)
		if err := flush(r, namespace); err != nil {
			t.Fatal("can't flush redis")
		}
	})

	tName = "second get"
	t.Run(tName, func(t *testing.T) {
		now := time.Now().UTC()
		_, err := m.getQueueLastPopTime(queue, now)
		if err != nil {
			t.Errorf("getQueueLastPopTime() error = %v", err)
		}

		err = m.updateQueueLastPopTime(queue, now.Add(time.Hour))
		if err != nil {
			t.Errorf("updateQueueLastPopTime() error = %v", err)
		}

		tm, err := m.getQueueLastPopTime(queue, now)
		if err != nil {
			t.Errorf("getQueueLastPopTime() error = %v", err)
		}

		a.Equal(tm.Unix(), now.Add(time.Hour).Unix(), tName)
		if err := flush(r, namespace); err != nil {
			t.Fatal("can't flush redis")
		}
	})
}

func getRedisPool() *redis.Pool {
	pool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")
	if err != nil {
		panic("redis not available")
	}

	return pool
}

func flush(r *redis.Pool, namespace string) error {
	conn := r.Get()
	defer conn.Close()
	keys, err := redis.Strings(conn.Do("KEYS", namespace+"*"))
	if err != nil {
		return err
	}
	args := make([]interface{}, 0, len(keys))
	for _, k := range keys {
		args = append(args, k)
	}
	if len(args) != 0 {
		_, err = conn.Do("DEL", args...)
		if err != nil {
			return err
		}
	}
	return nil
}

func zpopmin(r *redis.Pool, set string) (string, error) {
	conn := r.Get()
	defer conn.Close()
	v, err := redis.Strings(conn.Do("ZPOPMIN", set))
	return v[0], err
}

func zadd(r *redis.Pool, set, val string, score int) error {
	conn := r.Get()
	defer conn.Close()
	_, err := conn.Do("ZADD", set, score, val)
	return err
}