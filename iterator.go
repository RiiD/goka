package goka

import (
	"errors"
	"github.com/lovoo/goka/storage"
)

// Iterator allows one to iterate over the keys of a view.
type Iterator interface {
	// Next advances the iterator to the next KV-pair. Err should be called
	// after Next returns false to check whether the iteration finished
	// from exhaustion or was aborted due to an error.
	Next() bool
	Key() (interface{}, error)
	Value() (interface{}, error)
	Release()
	// Err returns the possible iteration error.
	Err() error
	Seek(key interface{}) (bool, error)
}

type iterator struct {
	iter       storage.Iterator
	valueCodec Codec
	keyCodec   Codec
}

// Next advances the iterator to the next key.
func (i *iterator) Next() bool {
	return i.iter.Next()
}

// Key returns the current key.
func (i *iterator) Key() (interface{}, error) {
	keyBytes := i.iter.Key()
	if keyBytes == nil {
		return nil, errors.New("nil key")
	}
	return i.valueCodec.Decode(keyBytes)
}

// Value returns the current value decoded by the codec of the storage.
func (i *iterator) Value() (interface{}, error) {
	data, err := i.iter.Value()
	if err != nil {
		return nil, err
	} else if data == nil {
		return nil, nil
	}
	return i.valueCodec.Decode(data)
}

// Err returns the possible iteration error.
func (i *iterator) Err() error {
	return i.iter.Err()
}

// Releases releases the iterator. The iterator is not usable anymore after calling Release.
func (i *iterator) Release() {
	i.iter.Release()
}

func (i *iterator) Seek(key interface{}) (bool, error) {
	keyBytes, err := i.keyCodec.Encode(key)
	if err != nil {
		return false, err
	}
	return i.iter.Seek(keyBytes), nil
}
