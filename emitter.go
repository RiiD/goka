package goka

import (
	"errors"
	"fmt"
	"github.com/lovoo/goka/codec"
	"sync"
)

var (
	// ErrEmitterAlreadyClosed is returned when Emit is called after the emitter has been finished.
	ErrEmitterAlreadyClosed error = errors.New("emitter already closed")
)

// Emitter emits messages into a specific Kafka topic, first encoding the message with the given codec.
type Emitter struct {
	valueCodec Codec
	keyCodec   Codec
	producer   Producer

	topic string

	wg   sync.WaitGroup
	done chan struct{}
}

// NewEmitter creates a new emitter using passed brokers, topic, codec and possibly options.
func NewEmitter(brokers []string, topic Stream, valueCodec Codec, options ...EmitterOption) (*Emitter, error) {
	return NewEmitterWithCustomKeyCodec(brokers, topic, new(codec.String), valueCodec, options...)
}

// NewEmitterWithCustomKeyCodec same as NewEmitter except it allows to specify custom key codec.
func NewEmitterWithCustomKeyCodec(brokers []string, topic Stream, keyCodec, valudeCodec Codec, options ...EmitterOption) (*Emitter, error) {
	options = append(
		// default options comes first
		[]EmitterOption{
			WithEmitterClientID(fmt.Sprintf("goka-emitter-%s", topic)),
		},

		// user-defined options (may overwrite default ones)
		options...,
	)

	opts := new(eoptions)

	opts.applyOptions(topic, keyCodec, valudeCodec, options...)

	prod, err := opts.builders.producer(brokers, opts.clientID, opts.hasher)
	if err != nil {
		return nil, fmt.Errorf(errBuildProducer, err)
	}

	return &Emitter{
		valueCodec: valudeCodec,
		keyCodec:   keyCodec,
		producer:   prod,
		topic:      string(topic),
		done:       make(chan struct{}),
	}, nil
}

// Emit sends a message for passed key using the emitter's codec.
func (e *Emitter) Emit(key, msg interface{}) (*Promise, error) {
	select {
	case <-e.done:
		return NewPromise().Finish(nil, ErrEmitterAlreadyClosed), nil
	default:
	}

	var (
		err  error
		data []byte
	)

	if msg != nil {
		data, err = e.valueCodec.Encode(msg)
		if err != nil {
			return nil, fmt.Errorf("Error encoding value for key %v in topic %s: %v", key, e.topic, err)
		}
	}

	keyBytes, err := e.keyCodec.Encode(key)
	if err != nil {
		return nil, fmt.Errorf("Error encoding key for key %v in topic %s: %v", key, e.topic, err)
	}

	e.wg.Add(1)
	return e.producer.Emit(e.topic, keyBytes, data).Then(func(err error) {
		e.wg.Done()
	}), nil
}

// EmitSync sends a message to passed topic and key.
func (e *Emitter) EmitSync(key interface{}, msg interface{}) error {
	var (
		err     error
		promise *Promise
	)
	promise, err = e.Emit(key, msg)

	if err != nil {
		return err
	}

	done := make(chan struct{})
	promise.Then(func(asyncErr error) {
		err = asyncErr
		close(done)
	})
	<-done
	return err
}

// Finish waits until the emitter is finished producing all pending messages.
func (e *Emitter) Finish() error {
	close(e.done)
	e.wg.Wait()
	return e.producer.Close()
}
