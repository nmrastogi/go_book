# go_book

A Go project providing an in-memory, thread-safe commit log.

## Structure

```
.
├── go.mod
├── internal/
│   └── server/
│       └── log.go   # Log implementation
└── README.md
```

## Package: server

The `server` package implements a simple append-only log with concurrent safe access.

### Types

- **`Log`** – In-memory log holding a slice of records and a mutex for synchronization.
- **`Record`** – A single log entry with `Value` ([]byte) and `Offset` (uint64), with JSON tags.

### API

| Symbol | Description |
|--------|-------------|
| `NewLog() *Log` | Constructs a new empty log. |
| `(*Log) Append(record Record) (uint64, error)` | Appends a record, assigns its offset, and returns the offset. |
| `(*Log) Read(offset uint64) (Record, error)` | Returns the record at the given offset, or `ErrOffsetNotFound` if out of range. |
| `ErrOffsetNotFound` | Error returned when `Read` is called with an invalid offset. |

### Example

```go
package main

import (
	"fmt"
	"github.com/nmrastogi/go_book/internal/server"
)

func main() {
	log := server.NewLog()

	offset, _ := log.Append(server.Record{Value: []byte("hello")})
	fmt.Println("appended at", offset) // 0

	rec, err := log.Read(0)
	if err != nil {
		panic(err)
	}
	fmt.Println("read:", string(rec.Value)) // hello
}
```

## Requirements

- Go 1.25.5 or later

## Build

```bash
go build ./...
```

## License

Unspecified.
