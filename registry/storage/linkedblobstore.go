package storage

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/reference"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type linkPathFunc func(name string, dgst digest.Digest) (string, error)

type linkedBlobStore struct {
	*blobStore
	registry               *registry
	blobServer             distribution.BlobServer
	blobAccessController   distribution.BlobDescriptorService
	repository             distribution.Repository
	ctx                    context.Context
	deleteEnabled          bool
	resumableDigestEnabled bool

	linkPath linkPathFunc

	linkDirectoryPathSpec pathSpec
}

var _ distribution.BlobStore = &linkedBlobStore{}

func (lbs *linkedBlobStore) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (lbs *linkedBlobStore) Get(ctx context.Context, dgst digest.Digest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lbs *linkedBlobStore) Open(ctx context.Context, dgst digest.Digest) (io.ReadSeekCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeekCloser), nil
}

func (lbs *linkedBlobStore) ServeBlob(ctx context.Context, w http.ResponseWriter, r *http.Request, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (lbs *linkedBlobStore) Put(ctx context.Context, mediaType string, p []byte) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

type optionFunc func(any) error

func (f optionFunc) Apply(v any) error { _ = "STUB: not implemented"; return nil }

func WithMountFrom(ref reference.Canonical) distribution.BlobCreateOption {
	_ = "STUB: not implemented"
	return *new(distribution.BlobCreateOption)
}

func (lbs *linkedBlobStore) Create(ctx context.Context, options ...distribution.BlobCreateOption) (distribution.BlobWriter, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobWriter), nil
}

func (lbs *linkedBlobStore) Resume(ctx context.Context, id string) (distribution.BlobWriter, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobWriter), nil
}

func (lbs *linkedBlobStore) Delete(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (lbs *linkedBlobStore) Enumerate(ctx context.Context, ingestor func(digest.Digest) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (lbs *linkedBlobStore) mount(ctx context.Context, sourceRepo reference.Named, dgst digest.Digest, sourceStat *v1.Descriptor) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (lbs *linkedBlobStore) newBlobUpload(ctx context.Context, uuid, path string, startedAt time.Time, append bool) (distribution.BlobWriter, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobWriter), nil
}

func (lbs *linkedBlobStore) linkBlob(ctx context.Context, canonical v1.Descriptor, aliases ...digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

type linkedBlobStatter struct {
	*blobStore
	repository distribution.Repository

	linkPath linkPathFunc
}

var _ distribution.BlobDescriptorService = &linkedBlobStatter{}

func (lbs *linkedBlobStatter) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (lbs *linkedBlobStatter) Clear(ctx context.Context, dgst digest.Digest) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (lbs *linkedBlobStatter) SetDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func blobLinkPath(name string, dgst digest.Digest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func manifestRevisionLinkPath(name string, dgst digest.Digest) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
