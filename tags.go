package distribution

import (
	"context"

	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type TagService interface {
	Get(ctx context.Context, tag string) (v1.Descriptor, error)

	Tag(ctx context.Context, tag string, desc v1.Descriptor) error

	Untag(ctx context.Context, tag string) error

	All(ctx context.Context) ([]string, error)

	Lookup(ctx context.Context, digest v1.Descriptor) ([]string, error)

	List(ctx context.Context, limit int, last string) ([]string, error)
}

type TagManifestsProvider interface {
	ManifestDigests(ctx context.Context, tag string) ([]digest.Digest, error)
}
