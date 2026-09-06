//go:build !noresumabledigest

package storage

import (
	"context"
)

func (bw *blobWriter) resumeDigest(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type hashStateEntry struct {
	offset int64
	path   string
}

func (bw *blobWriter) getStoredHashStates(ctx context.Context) ([]hashStateEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bw *blobWriter) storeHashState(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
