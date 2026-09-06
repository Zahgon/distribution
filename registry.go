package distribution

import (
	"context"

	"github.com/distribution/reference"
)

type Scope interface {
	Contains(name string) bool
}

type fullScope struct{}

func (f fullScope) Contains(string) bool { _ = "STUB: not implemented"; return false }

var GlobalScope = Scope(fullScope{})

type Namespace interface {
	Scope() Scope

	Repository(ctx context.Context, name reference.Named) (Repository, error)

	Repositories(ctx context.Context, repos []string, last string) (n int, err error)

	Blobs() BlobEnumerator

	BlobStatter() BlobStatter
}

type RepositoryEnumerator interface {
	Enumerate(ctx context.Context, ingester func(string) error) error
}

type RepositoryRemover interface {
	Remove(ctx context.Context, name reference.Named) error
}

type ManifestServiceOption interface {
	Apply(ManifestService) error
}

func WithTag(tag string) ManifestServiceOption {
	_ = "STUB: not implemented"
	return *new(ManifestServiceOption)
}

type WithTagOption struct{ Tag string }

func (o WithTagOption) Apply(m ManifestService) error { _ = "STUB: not implemented"; return nil }

func WithManifestMediaTypes(mediaTypes []string) ManifestServiceOption {
	_ = "STUB: not implemented"
	return *new(ManifestServiceOption)
}

type WithManifestMediaTypesOption struct{ MediaTypes []string }

func (o WithManifestMediaTypesOption) Apply(m ManifestService) error {
	_ = "STUB: not implemented"
	return nil
}

type Repository interface {
	Named() reference.Named

	Manifests(ctx context.Context, options ...ManifestServiceOption) (ManifestService, error)

	Blobs(ctx context.Context) BlobStore

	Tags(ctx context.Context) TagService
}
