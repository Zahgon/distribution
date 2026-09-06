package storage

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/distribution/distribution/v3"
	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

var errResumableDigestNotAvailable = errors.New("resumable digest not available")

const (
	digestSha256Empty = "sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
)

type blobWriter struct {
	ctx       context.Context
	blobStore *linkedBlobStore

	id        string
	startedAt time.Time
	digester  digest.Digester
	written   int64

	fileWriter storagedriver.FileWriter
	driver     storagedriver.StorageDriver
	path       string

	resumableDigestEnabled bool
	committed              bool
}

var _ distribution.BlobWriter = &blobWriter{}

func (bw *blobWriter) ID() string { _ = "STUB: not implemented"; return "" }

func (bw *blobWriter) StartedAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (bw *blobWriter) Commit(ctx context.Context, desc v1.Descriptor) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (bw *blobWriter) Cancel(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (bw *blobWriter) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (bw *blobWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (bw *blobWriter) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (bw *blobWriter) Close() error { _ = "STUB: not implemented"; return nil }

func (bw *blobWriter) validateBlob(ctx context.Context, desc v1.Descriptor) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (bw *blobWriter) moveBlob(ctx context.Context, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (bw *blobWriter) removeResources(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (bw *blobWriter) Reader() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
