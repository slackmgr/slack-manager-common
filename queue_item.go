package common

import (
	"context"
	"time"
)

type QueueItem struct {
	MessageID         string
	GroupID           string
	ReceiveTimestamp  time.Time
	VisibilityTimeout time.Duration
	Body              string
	Ack               func(ctx context.Context) error
	Extend            func(ctx context.Context) error
}
