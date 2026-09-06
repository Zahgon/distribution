package storage

import (
	"context"
	"io"

	"github.com/distribution/distribution/v3/registry/storage/driver"
)

const (
	maxBlobGetSize = 4 * 1024 * 1024
)

func getContent(ctx context.Context, driver driver.StorageDriver, p string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readAllLimited(r io.Reader, limit int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func limitReader(r io.Reader, n int64) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

type limitedReader struct {
	r   io.Reader
	n   int64
	err error
}

func (l *limitedReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
