package message

import "github.com/google/uuid"

type Message struct {
	// Identity & tracing
	ID         uuid.UUID // globally unique (uuid/ulid)
	TraceID    uuid.UUID // optional: distributed tracing
	ProducedAt int64     // unix millis

	// Routing
	Topic   string            // logical stream/topic
	Key     []byte            // partitioning/ordering key (optional)
	Headers map[string][]byte // extensible metadata

	// Delivery controls
	TTLMillis   int64  // 0 => no TTL
	DelayMillis int64  // 0 => no delay
	Priority    uint8  // optional
	Attempt     uint16 // delivery attempt counter

	// Payload
	ContentType string // e.g. application/json
	Body        []byte // opaque payload bytes

	// Integrity (optional)
	BodyCRC32 uint32
	Version   uint16 // message schema version
}
