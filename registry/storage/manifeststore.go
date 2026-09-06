package storage

import (
	"context"

	"github.com/distribution/distribution/v3"
	"github.com/opencontainers/go-digest"
)

type ManifestHandler interface {
	Unmarshal(ctx context.Context, dgst digest.Digest, content []byte) (distribution.Manifest, error)

	Put(ctx context.Context, manifest distribution.Manifest, skipDependencyVerification bool) (digest.Digest, error)
}

func SkipLayerVerification() distribution.ManifestServiceOption {
	_ = "STUB: not implemented"
	return *new(distribution.ManifestServiceOption)
}

type skipLayerOption struct{}

func (o skipLayerOption) Apply(m distribution.ManifestService) error {
	_ = "STUB: not implemented"
	return nil
}

type manifestStore struct {
	repository *repository
	blobStore  *linkedBlobStore
	ctx        context.Context

	skipDependencyVerification bool

	schema2Handler        ManifestHandler
	manifestListHandler   ManifestHandler
	ocischemaHandler      ManifestHandler
	ocischemaIndexHandler ManifestHandler
}

var _ distribution.ManifestService = &manifestStore{}

func (ms *manifestStore) Exists(ctx context.Context, dgst digest.Digest) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ms *manifestStore) Get(ctx context.Context, dgst digest.Digest, options ...distribution.ManifestServiceOption) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}

func (ms *manifestStore) Put(ctx context.Context, manifest distribution.Manifest, options ...distribution.ManifestServiceOption) (digest.Digest, error) {
	_ = "STUB: not implemented"
	return *new(digest.Digest), nil
}

func (ms *manifestStore) Delete(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (ms *manifestStore) Enumerate(ctx context.Context, ingester func(digest.Digest) error) error {
	_ = "STUB: not implemented"
	return nil
}
