package goka

import "errors"

var (
	errBuildConsumer = "error creating Kafka consumer: %v"
	errBuildProducer = "error creating Kafka producer: %v"
	errApplyOptions  = "error applying options: %v"

	ErrNoKeyCodec   = errors.New("no key codec")
	ErrNoValueCodec = errors.New("no value codec")
)
