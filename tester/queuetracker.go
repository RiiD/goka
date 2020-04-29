package tester

// QueueTracker tracks message offsets for each topic for convenient
// 'expect message x to be in topic y' in unit tests
type QueueTracker struct {
	t          T
	topic      string
	nextOffset int64
	tester     *Tester
}

func newQueueTracker(tester *Tester, t T, topic string) *QueueTracker {
	return &QueueTracker{
		t:      t,
		topic:  topic,
		tester: tester,
	}
}

// Next returns the next message since the last time this
// function was called (or MoveToEnd)
// It uses the known codec for the topic to decode the message
func (mt *QueueTracker) Next() (interface{}, interface{}, bool) {

	key, msgRaw, hasNext := mt.NextRaw()

	decodedValue, err := mt.tester.valueCodecForTopic(mt.topic).Decode(msgRaw)
	if err != nil {
		mt.t.Fatalf("Error decoding message value: %v", err)
	}

	decodedKey, err := mt.tester.keyCodecForTopic(mt.topic).Decode(key)
	if err != nil {
		mt.t.Fatalf("Error decoding message key: %v", err)
	}

	return decodedKey, decodedValue, hasNext
}

// NextRaw returns the next message similar to Next(), but without the decoding
func (mt *QueueTracker) NextRaw() ([]byte, []byte, bool) {
	q := mt.tester.queueForTopic(mt.topic)
	if int(mt.nextOffset) >= q.size() {
		return nil, nil, false
	}
	msg := q.message(int(mt.nextOffset))

	mt.nextOffset++
	return msg.key, msg.value, true
}

// Seek moves the index pointer of the queue tracker to passed offset
func (mt *QueueTracker) Seek(offset int64) {
	mt.nextOffset = offset
}

// Hwm returns the tracked queue's hwm value
func (mt *QueueTracker) Hwm() int64 {
	return mt.tester.queueForTopic(mt.topic).hwm
}

// NextOffset returns the tracker's next offset
func (mt *QueueTracker) NextOffset() int64 {
	return mt.nextOffset
}
