package storage

import (
	"context"
	"io"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type blobStore struct {
	driver  driver.StorageDriver
	statter distribution.BlobStatter
}

var _ distribution.BlobProvider = &blobStore{}

func (bs *blobStore) Get(ctx context.Context, dgst digest.Digest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bs *blobStore) Open(ctx context.Context, dgst digest.Digest) (io.ReadSeekCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeekCloser), nil
}

func (bs *blobStore) Put(ctx context.Context, mediaType string, p []byte) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (bs *blobStore) Enumerate(ctx context.Context, ingester func(dgst digest.Digest) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *blobStore) path(dgst digest.Digest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (bs *blobStore) link(ctx context.Context, path string, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *blobStore) readlink(ctx context.Context, path string) (digest.Digest, error) {
	_ = "STUB: not implemented"
	return *new(digest.Digest), nil
}

type blobStatter struct {
	driver driver.StorageDriver
}

var _ distribution.BlobDescriptorService = &blobStatter{}

func (bs *blobStatter) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (bs *blobStatter) Clear(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *blobStatter) SetDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}
