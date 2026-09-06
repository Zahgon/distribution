package storage

import (
	"context"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/manifest/ocischema"
	"github.com/opencontainers/go-digest"
)

type ocischemaManifestHandler struct {
	repository   distribution.Repository
	blobStore    distribution.BlobStore
	ctx          context.Context
	manifestURLs manifestURLs
}

var _ ManifestHandler = &ocischemaManifestHandler{}

func (ms *ocischemaManifestHandler) Unmarshal(ctx context.Context, dgst digest.Digest, content []byte) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}

func (ms *ocischemaManifestHandler) Put(ctx context.Context, manifest distribution.Manifest, skipDependencyVerification bool) (digest.Digest, error) {
	_ = "STUB: not implemented"
	return *new(digest.Digest), nil
}

func (ms *ocischemaManifestHandler) verifyManifest(ctx context.Context, mnfst ocischema.DeserializedManifest, skipDependencyVerification bool) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck // ignore A1019: v1.MediaTypeImageLayerNonDistributable is deprecated: Non-distributable layers are deprecated, and not recommended for future use.
