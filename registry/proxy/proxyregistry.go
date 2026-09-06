package proxy

import (
	"context"
	"net/url"
	"sync"
	"time"

	"github.com/distribution/reference"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/configuration"
	"github.com/distribution/distribution/v3/internal/client/auth"
	"github.com/distribution/distribution/v3/internal/client/auth/challenge"
	"github.com/distribution/distribution/v3/registry/proxy/scheduler"
	"github.com/distribution/distribution/v3/registry/storage/driver"
)

var repositoryTTL = 24 * 7 * time.Hour

type proxyingRegistry struct {
	embedded          distribution.Namespace
	scheduler         *scheduler.TTLExpirationScheduler
	ttl               *time.Duration
	cacheWriteTimeout time.Duration
	remoteURL         url.URL
	authChallenger    authChallenger
	basicAuth         auth.CredentialStore
}

func NewRegistryPullThroughCache(ctx context.Context, registry distribution.Namespace, driver driver.StorageDriver, config configuration.Proxy) (distribution.Namespace, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Namespace), nil
}

func (pr *proxyingRegistry) Scope() distribution.Scope {
	_ = "STUB: not implemented"
	return *new(distribution.Scope)
}

func (pr *proxyingRegistry) Repositories(ctx context.Context, repos []string, last string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (pr *proxyingRegistry) Repository(ctx context.Context, name reference.Named) (distribution.Repository, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Repository), nil
}

func (pr *proxyingRegistry) Blobs() distribution.BlobEnumerator {
	_ = "STUB: not implemented"
	return *new(distribution.BlobEnumerator)
}

func (pr *proxyingRegistry) BlobStatter() distribution.BlobStatter {
	_ = "STUB: not implemented"
	return *new(distribution.BlobStatter)
}

type Closer interface {
	Close() error
}

func (pr *proxyingRegistry) Close() error { _ = "STUB: not implemented"; return nil }

type authChallenger interface {
	tryEstablishChallenges(context.Context) error
	challengeManager() challenge.Manager
	credentialStore() auth.CredentialStore
}

type remoteAuthChallenger struct {
	remoteURL url.URL
	sync.Mutex
	cm challenge.Manager
	cs auth.CredentialStore
}

func (r *remoteAuthChallenger) credentialStore() auth.CredentialStore {
	_ = "STUB: not implemented"
	return *new(auth.CredentialStore)
}

func (r *remoteAuthChallenger) challengeManager() challenge.Manager {
	_ = "STUB: not implemented"
	return *new(challenge.Manager)
}

func (r *remoteAuthChallenger) tryEstablishChallenges(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type proxiedRepository struct {
	blobStore distribution.BlobStore
	manifests distribution.ManifestService
	name      reference.Named
	tags      distribution.TagService
}

func (pr *proxiedRepository) Manifests(ctx context.Context, options ...distribution.ManifestServiceOption) (distribution.ManifestService, error) {
	_ = "STUB: not implemented"
	return *new(distribution.ManifestService), nil
}

func (pr *proxiedRepository) Blobs(ctx context.Context) distribution.BlobStore {
	_ = "STUB: not implemented"
	return *new(distribution.BlobStore)
}

func (pr *proxiedRepository) Named() reference.Named {
	_ = "STUB: not implemented"
	return *new(reference.Named)
}

func (pr *proxiedRepository) Tags(ctx context.Context) distribution.TagService {
	_ = "STUB: not implemented"
	return *new(distribution.TagService)
}
