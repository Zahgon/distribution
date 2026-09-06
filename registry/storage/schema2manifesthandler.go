package storage

import (
	"context"
	"errors"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/manifest/schema2"
	"github.com/opencontainers/go-digest"
)

var (
	errMissingURL = errors.New("missing URL on layer")
	errInvalidURL = errors.New("invalid URL on layer")
)

type schema2ManifestHandler struct {
	repository   distribution.Repository
	blobStore    distribution.BlobStore
	ctx          context.Context
	manifestURLs manifestURLs
}

var _ ManifestHandler = &schema2ManifestHandler{}

func (ms *schema2ManifestHandler) Unmarshal(ctx context.Context, dgst digest.Digest, content []byte) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}

func (ms *schema2ManifestHandler) Put(ctx context.Context, manifest distribution.Manifest, skipDependencyVerification bool) (digest.Digest, error) {
	_ = "STUB: not implemented"
	return *new(digest.Digest), nil
}

func (ms *schema2ManifestHandler) verifyManifest(ctx context.Context, mnfst schema2.DeserializedManifest, skipDependencyVerification bool) error {
	_ = "STUB: not implemented"
	return nil
}
