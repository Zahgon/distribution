package storage

import (
	"bufio"
	"context"
	"io"

	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
)

const fileReaderBufferSize = 4 * 1024 * 1024

type fileReader struct {
	driver storagedriver.StorageDriver

	ctx context.Context

	path string
	size int64

	rc     io.ReadCloser
	brd    *bufio.Reader
	offset int64
	err    error
}

func newFileReader(ctx context.Context, driver storagedriver.StorageDriver, path string, size int64) (*fileReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fr *fileReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (fr *fileReader) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (fr *fileReader) Close() error { _ = "STUB: not implemented"; return nil }

func (fr *fileReader) reader() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (fr *fileReader) reset() { _ = "STUB: not implemented"; return }

func (fr *fileReader) closeWithErr(err error) error { _ = "STUB: not implemented"; return nil }
