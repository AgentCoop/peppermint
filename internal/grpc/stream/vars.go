package stream

import "errors"

var (
	ErrEmptyEncryptionKey       = errors.New("stream: encryption key is empty")
	ErrFailedToRetrieveMetadata = errors.New("stream: failed to retrieve metadata")
)
