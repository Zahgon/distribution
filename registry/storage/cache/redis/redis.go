package redis

import (
	"context"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/storage/cache"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/redis/go-redis/v9"
)

type redisBlobDescriptorService struct {
	pool redis.UniversalClient
}

var _ distribution.BlobDescriptorService = &redisBlobDescriptorService{}

func NewRedisBlobDescriptorCacheProvider(pool redis.UniversalClient) cache.BlobDescriptorCacheProvider {
	_ = "STUB: not implemented"
	return *new(cache.BlobDescriptorCacheProvider)
}

func (rbds *redisBlobDescriptorService) RepositoryScoped(repo string) (distribution.BlobDescriptorService, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobDescriptorService), nil
}

func (rbds *redisBlobDescriptorService) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (rbds *redisBlobDescriptorService) Clear(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (rbds *redisBlobDescriptorService) stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (rbds *redisBlobDescriptorService) SetDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (rbds *redisBlobDescriptorService) setDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (rbds *redisBlobDescriptorService) blobDescriptorHashKey(dgst digest.Digest) string {
	_ = "STUB: not implemented"
	return ""
}

type repositoryScopedRedisBlobDescriptorService struct {
	repo     string
	upstream *redisBlobDescriptorService
}

var _ distribution.BlobDescriptorService = &repositoryScopedRedisBlobDescriptorService{}

func (rsrbds *repositoryScopedRedisBlobDescriptorService) Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (rsrbds *repositoryScopedRedisBlobDescriptorService) Clear(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (rsrbds *repositoryScopedRedisBlobDescriptorService) SetDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (rsrbds *repositoryScopedRedisBlobDescriptorService) setDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (rsrbds *repositoryScopedRedisBlobDescriptorService) blobDescriptorHashKey(dgst digest.Digest) string {
	_ = "STUB: not implemented"
	return ""
}

func (rsrbds *repositoryScopedRedisBlobDescriptorService) repositoryBlobSetKey(repo string) string {
	_ = "STUB: not implemented"
	return ""
}
