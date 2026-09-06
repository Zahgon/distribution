package proxy

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/proxy/scheduler"
	"github.com/distribution/reference"
)

type proxyBlobStore struct {
	localStore        distribution.BlobStore
	remoteStore       distribution.BlobService
	scheduler         *scheduler.TTLExpirationScheduler
	ttl               *time.Duration
	cacheWriteTimeout time.Duration
	repositoryName    reference.Named
	authChallenger    authChallenger
}

var _ distribution.BlobStore = &proxyBlobStore{}

var inflight = make(map[digest.Digest]struct{})

var mu sync.Mutex

func setResponseHeaders(h http.Header, length int64, mediaType string, digest digest.Digest) {
	_ = "STUB: not implemented"
	return
}

func (pbs *proxyBlobStore) copyContent(ctx context.Context, dgst digest.Digest, writer io.Writer, h http.Header) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (pbs *proxyBlobStore) serveLocal(ctx context.Context, w http.ResponseWriter, r *http.Request, dgst digest.Digest) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (pbs *proxyBlobStore) ServeBlob(ctx context.Context, w http.ResponseWriter, r *http.Request, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (pbs *proxyBlobStore) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (pbs *proxyBlobStore) Get(ctx context.Context, dgst digest.Digest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pbs *proxyBlobStore) Put(ctx context.Context, mediaType string, p []byte) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (pbs *proxyBlobStore) Create(ctx context.Context, options ...distribution.BlobCreateOption) (distribution.BlobWriter, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobWriter), nil
}

func (pbs *proxyBlobStore) Resume(ctx context.Context, id string) (distribution.BlobWriter, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobWriter), nil
}

func (pbs *proxyBlobStore) Mount(ctx context.Context, sourceRepo reference.Named, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (pbs *proxyBlobStore) Open(ctx context.Context, dgst digest.Digest) (io.ReadSeekCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeekCloser), nil
}

func (pbs *proxyBlobStore) Delete(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}
