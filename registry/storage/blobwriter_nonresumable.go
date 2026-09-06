//go:build noresumabledigest

package storage

import (
	"context"
)

func (bw *blobWriter) resumeDigest(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (bw *blobWriter) storeHashState(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
