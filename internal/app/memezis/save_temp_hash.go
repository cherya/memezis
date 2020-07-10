package memezis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

func (i *Memezis) saveHash(filename, hash string) error {
	conn := i.redis.Get()
	defer conn.Close()

	ex := (time.Hour * 24).Seconds()
	key := hashKey(filename)
	_, err := conn.Do("SET", key, hash, "EX", ex)
	if err != nil {
		return errors.Wrapf(err, "can't set [%s] – [%s] to redis", key, hash)
	}

	return nil
}

func (i *Memezis) getHash(filename string) (string, error) {
	conn := i.redis.Get()
	defer conn.Close()

	key := hashKey(filename)

	hash, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", errors.Wrapf(err, "can't get [%s] from redis", key)
	}

	return hash, nil
}

func hashKey(filename string) string {
	return fmt.Sprintf("hash:%s", filename)
}