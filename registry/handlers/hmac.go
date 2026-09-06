package handlers

import (
	"fmt"
	"time"
)

type blobUploadState struct {
	Name string

	UUID string

	Offset int64

	StartedAt time.Time
}

type hmacKey string

var errInvalidSecret = fmt.Errorf("invalid secret")

func (secret hmacKey) unpackUploadState(token string) (blobUploadState, error) {
	_ = "STUB: not implemented"
	return *new(blobUploadState), nil
}

func (secret hmacKey) packUploadState(lus blobUploadState) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
