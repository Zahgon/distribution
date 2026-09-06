package cache

import (
	"context"

	"github.com/distribution/distribution/v3"
	prometheus "github.com/distribution/distribution/v3/metrics"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type cachedBlobStatter struct {
	cache   distribution.BlobDescriptorService
	backend distribution.BlobDescriptorService
}

var (
	cacheRequestCount = prometheus.StorageNamespace.NewCounter("cache_requests", "The number of cache request received")

	cacheHitCount = prometheus.StorageNamespace.NewCounter("cache_hits", "The number of cache request received")

	cacheErrorCount = prometheus.StorageNamespace.NewCounter("cache_errors", "The number of cache request errors")
)

func NewCachedBlobStatter(cache distribution.BlobDescriptorService, backend distribution.BlobDescriptorService) distribution.BlobDescriptorService {
	_ = "STUB: not implemented"
	return *new(distribution.BlobDescriptorService)
}

func (cbds *cachedBlobStatter) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (cbds *cachedBlobStatter) Clear(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (cbds *cachedBlobStatter) SetDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}
