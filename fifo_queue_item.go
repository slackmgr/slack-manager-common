package common

import (
	"context"
	"time"
)

// FifoQueueItem represents an item received from a FIFO queue.
type FifoQueueItem struct {
	// MessageID is the unique identifier of the message (as defined by the queue implementation).
	MessageID string

	// SlackChannelID is the ID of the Slack channel to which the message is related.
	SlackChannelID string

	// ReceiveTimestamp is the time when the message was received from the queue.
	ReceiveTimestamp time.Time

	// Body is the body of the message.
	Body string

	// Ack acknowledges the successful processing of the message, effectively removing it from the queue.
	// This function cannot be nil.
	Ack func(ctx context.Context)

	// Nack negatively acknowledges the processing of the message, thus making it available for reprocessing.
	// This function cannot be nil.
	Nack func(ctx context.Context)
}
