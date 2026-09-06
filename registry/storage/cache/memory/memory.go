package memory

import (
	"context"
	"math"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/storage/cache"
	"github.com/hashicorp/golang-lru/arc/v2"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

const (
	DefaultSize = 10000

	UnlimitedSize = math.MaxInt
)

type descriptorCacheKey struct {
	digest digest.Digest
	repo   string
}

type inMemoryBlobDescriptorCacheProvider struct {
	lru *arc.ARCCache[descriptorCacheKey, v1.Descriptor]
}

func NewInMemoryBlobDescriptorCacheProvider(size int) cache.BlobDescriptorCacheProvider {
	_ = "STUB: not implemented"
	return *new(cache.BlobDescriptorCacheProvider)
}

func (imbdcp *inMemoryBlobDescriptorCacheProvider) RepositoryScoped(repo string) (distribution.BlobDescriptorService, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobDescriptorService), nil
}

func (imbdcp *inMemoryBlobDescriptorCacheProvider) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (imbdcp *inMemoryBlobDescriptorCacheProvider) Clear(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (imbdcp *inMemoryBlobDescriptorCacheProvider) SetDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

type repositoryScopedInMemoryBlobDescriptorCache struct {
	repo   string
	parent *inMemoryBlobDescriptorCacheProvider
}

func (rsimbdcp *repositoryScopedInMemoryBlobDescriptorCache) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (rsimbdcp *repositoryScopedInMemoryBlobDescriptorCache) Clear(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (rsimbdcp *repositoryScopedInMemoryBlobDescriptorCache) SetDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}
