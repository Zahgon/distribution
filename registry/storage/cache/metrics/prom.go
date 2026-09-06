package metrics

import (
	"context"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/storage/cache"
	"github.com/docker/go-metrics"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type prometheusCacheProvider struct {
	cache.BlobDescriptorCacheProvider
	latencyTimer metrics.LabeledTimer
}

func NewPrometheusCacheProvider(wrap cache.BlobDescriptorCacheProvider, name, help string) cache.BlobDescriptorCacheProvider {
	_ = "STUB: not implemented"
	return *new(cache.BlobDescriptorCacheProvider)
}

func (p *prometheusCacheProvider) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (p *prometheusCacheProvider) SetDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

type prometheusRepoCacheProvider struct {
	distribution.BlobDescriptorService
	latencyTimer metrics.LabeledTimer
}

func (p *prometheusRepoCacheProvider) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (p *prometheusRepoCacheProvider) SetDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *prometheusCacheProvider) RepositoryScoped(repo string) (distribution.BlobDescriptorService, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobDescriptorService), nil
}
