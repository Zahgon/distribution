package proxy

import (
	"context"
	"time"

	"github.com/opencontainers/go-digest"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/proxy/scheduler"
	"github.com/distribution/reference"
)

type proxyManifestStore struct {
	ctx             context.Context
	localManifests  distribution.ManifestService
	remoteManifests distribution.ManifestService
	repositoryName  reference.Named
	scheduler       *scheduler.TTLExpirationScheduler
	ttl             *time.Duration
	authChallenger  authChallenger
}

var _ distribution.ManifestService = &proxyManifestStore{}

func (pms proxyManifestStore) Exists(ctx context.Context, dgst digest.Digest) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (pms proxyManifestStore) Get(ctx context.Context, dgst digest.Digest, options ...distribution.ManifestServiceOption) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}

func (pms proxyManifestStore) Put(ctx context.Context, manifest distribution.Manifest, options ...distribution.ManifestServiceOption) (digest.Digest, error) {
	_ = "STUB: not implemented"
	return *new(digest.Digest), nil
}

func (pms proxyManifestStore) Delete(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}
