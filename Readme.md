```txt
mq/
├── cmd/
│   ├── broker/                 # broker binary entrypoint
│   │   └── main.go
│   └── mqctl/                  # CLI tool (admin + debug)
│       └── main.go
├── internal/
│   ├── broker/                 # orchestrates server + core engine
│   │   ├── broker.go
│   │   ├── lifecycle.go
│   │   └── wiring.go
│   ├── protocol/               # wire protocols (tcp, http, grpc)
│   │   ├── tcp/
│   │   │   ├── server.go
│   │   │   ├── codec.go
│   │   │   └── handlers.go
│   │   ├── http/
│   │   └── common/             # shared request/response models
│   ├── core/                   # pure queue semantics (no networking)
│   │   ├── api.go              # core interfaces
│   │   ├── pubsub/
│   │   │   ├── router.go
│   │   │   ├── subscription.go
│   │   │   └── dispatcher.go
│   │   ├── topic/
│   │   │   ├── topic.go
│   │   │   └── partition.go
│   │   ├── consumer/
│   │   │   ├── group.go
│   │   │   ├── offset_store.go
│   │   │   └── ack.go
│   │   └── errors.go
│   ├── storage/                # durability primitives
│   │   ├── log/
│   │   │   ├── wal.go
│   │   │   ├── segment.go
│   │   │   ├── index.go
│   │   │   └── retention.go
│   │   ├── kv/                 # metadata/offset store (inmem -> boltdb/pebble later)
│   │   └── fs/                 # filesystem helpers, atomic rename, sync
│   ├── metadata/               # topics, configs, broker state
│   │   ├── registry.go
│   │   └── schema.go
│   ├── auth/                   # authn/authz (token, mTLS later)
│   ├── observability/
│   │   ├── metrics.go
│   │   ├── logging.go
│   │   └── tracing.go
│   └── config/
│       ├── config.go
│       └── defaults.go
├── pkg/                        # optional: public client library
│   └── client/
│       ├── client.go
│       ├── producer.go
│       ├── consumer.go
│       └── protocol.go
├── api/                        # protocol definitions
│   ├── proto/                  # gRPC/protobuf (if used)
│   └── spec/                   # command frames/spec docs
├── test/
│   ├── integration/
│   └── chaos/
├── configs/
│   └── broker.yaml
├── scripts/
├── Makefile
├── go.mod
└── README.md

```
