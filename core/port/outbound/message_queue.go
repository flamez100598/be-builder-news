package outbound

import (
	"context"
)

type MessageQueue interface {
	ReadMessage(ctx context.Context) (key []byte, value []byte, err error)
}

