package redis

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/lovoo/goka/storage"

	"gopkg.in/redis.v5"
)

var (
	offsetKey = []byte("__offset")
)

type redisStorage struct {
	client *redis.Client
	hash   string
}

var _ storage.Storage = &redisStorage{}

// New creates a new Storage backed by Redis.
func New(client *redis.Client, hash string) (storage.Storage, error) {
	if client == nil {
		return nil, errors.New("invalid redis client")
	}
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &redisStorage{
		client: client,
		hash:   hash,
	}, nil
}

func (s *redisStorage) Has(key []byte) (bool, error) {
	return s.client.HExists(s.hash, string(key)).Result()
}

func (s *redisStorage) Get(key []byte) ([]byte, error) {
	has, err := s.client.HExists(s.hash, string(key)).Result()
	if err != nil {
		return nil, fmt.Errorf("error checking for existence in redis (key %s): %v", key, err)
	} else if !has {
		return nil, nil
	}
	value, err := s.client.HGet(s.hash, string(key)).Bytes()
	if err != nil {
		return nil, fmt.Errorf("error getting from redis (key %s): %v", key, err)
	}
	return value, nil
}

func (s *redisStorage) GetOffset(defValue int64) (int64, error) {
	data, err := s.Get(offsetKey)
	if err != nil {
		return 0, err
	}
	if data == nil {
		return defValue, nil
	}

	value, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error decoding redis offset (%s): %v", string(data), err)
	}
	return value, nil
}

func (s *redisStorage) Set(key, value []byte) error {
	err := s.client.HSet(s.hash, string(key), value).Err()
	if err != nil {
		return fmt.Errorf("error setting to redis (key %s): %v", key, err)
	}
	return nil
}

func (s *redisStorage) SetOffset(offset int64) error {
	return s.Set(offsetKey, []byte(strconv.FormatInt(offset, 10)))
}

func (s *redisStorage) Delete(key []byte) error {
	return s.client.HDel(s.hash, string(key)).Err()
}

func (s *redisStorage) Iterator() (storage.Iterator, error) {
	var current uint64
	var keys []string
	var err error

	keys, current, err = s.client.HScan(s.hash, current, "", 0).Result()
	if err != nil {
		return nil, err
	}
	return &redisIterator{
		current: current,
		keys:    keys,
		client:  s.client,
		hash:    s.hash,
	}, nil
}

func (s *redisStorage) IteratorWithRange(start, _ []byte) (storage.Iterator, error) {
	var current uint64
	var keys []string
	var err error

	keys, current, err = s.client.HScan(s.hash, current, string(start), 0).Result()
	if err != nil {
		return nil, err
	}
	return &redisIterator{
		current: current,
		keys:    keys,
		client:  s.client,
		hash:    s.hash,
	}, nil
}

func (s *redisStorage) Recovered() bool {
	return false
}

func (s *redisStorage) MarkRecovered() error {
	return nil
}

func (s *redisStorage) Open() error {
	return nil
}

func (s *redisStorage) Close() error {
	return nil
}

type redisIterator struct {
	current uint64
	keys    []string
	client  *redis.Client
	hash    string
}

func (i *redisIterator) exhausted() bool {
	return uint64(len(i.keys)) <= i.current
}

func (i *redisIterator) Next() bool {
	i.current++
	if bytes.Equal(i.Key(), offsetKey) {
		i.current++
	}
	return !i.exhausted()
}

func (i *redisIterator) Key() []byte {
	if i.exhausted() {
		return nil
	}
	key := i.keys[i.current]
	return []byte(key)
}

func (i *redisIterator) Err() error {
	return nil
}

func (i *redisIterator) Value() ([]byte, error) {
	if i.exhausted() {
		return nil, nil
	}
	key := i.keys[i.current]
	return i.client.HGet(i.hash, key).Bytes()
}

func (i *redisIterator) Release() {
	i.current = uint64(len(i.keys))
}

func (i *redisIterator) Seek(_ []byte) bool {
	return !i.exhausted()
}
