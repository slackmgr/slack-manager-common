package types

import "encoding/json"

// MoveMapping represents an issue that has been moved from one channel to another, based on the correlation ID.
// It is used to ensure that new alerts with the same correlation ID are not processed in the original channel,
// but instead are processed in the new channel.
//
// The actual implementation of this interface is internal to the Slack Manager, and may change without notice.
// The database implementation should only store a JSON representation of the move mapping, and should not depend
// on any specific fields or structure of the move mapping.
type MoveMapping interface {
	json.Marshaler

	// ChannelID returns the Slack channel ID that this move mapping belongs to (i.e. the channel where the move was initiated).
	ChannelID() string

	// UniqueID returns a unique and deterministic ID for this move mapping, for database/storage purposes.
	// The ID is based on the original channel and the correlation ID, and is base64 encoded to ensure it is safe for use in URLs and as a database key.
	UniqueID() string

	// GetCorrelationID returns the correlation ID that this move mapping is associated with, for the given channel.
	// It is not URL safe, and should thus be encoded before being used in URLs or as part of a database key.
	GetCorrelationID() string
}
