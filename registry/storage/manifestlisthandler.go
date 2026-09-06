package storage

import (
	"context"

	"github.com/distribution/distribution/v3"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type manifestListHandler struct {
	repository           distribution.Repository
	blobStore            distribution.BlobStore
	ctx                  context.Context
	validateImageIndexes validateImageIndexes
}

var _ ManifestHandler = &manifestListHandler{}

func (ms *manifestListHandler) Unmarshal(ctx context.Context, dgst digest.Digest, content []byte) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}

func (ms *manifestListHandler) Put(ctx context.Context, manifestList distribution.Manifest, skipDependencyVerification bool) (digest.Digest, error) {
	_ = "STUB: not implemented"
	return *new(digest.Digest), nil
}

func (ms *manifestListHandler) verifyManifest(ctx context.Context, mnfst distribution.Manifest, skipDependencyVerification bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms *manifestListHandler) platformMustExist(descriptor v1.Descriptor) bool {
	_ = "STUB: not implemented"
	return false
}
