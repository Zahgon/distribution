package distribution

import (
	"context"

	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type Manifest interface {
	References() []v1.Descriptor

	Payload() (mediaType string, payload []byte, err error)
}

type ManifestService interface {
	Exists(ctx context.Context, dgst digest.Digest) (bool, error)

	Get(ctx context.Context, dgst digest.Digest, options ...ManifestServiceOption) (Manifest, error)

	Put(ctx context.Context, manifest Manifest, options ...ManifestServiceOption) (digest.Digest, error)

	Delete(ctx context.Context, dgst digest.Digest) error
}

type ManifestEnumerator interface {
	Enumerate(ctx context.Context, ingester func(digest.Digest) error) error
}

type Describable interface {
	Descriptor() v1.Descriptor
}

func ManifestMediaTypes() (mediaTypes []string) { _ = "STUB: not implemented"; return nil }

type UnmarshalFunc func([]byte) (Manifest, v1.Descriptor, error)

var mappings = make(map[string]UnmarshalFunc)

func UnmarshalManifest(ctHeader string, p []byte) (Manifest, v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(Manifest), *new(v1.Descriptor), nil
}

func RegisterManifestSchema(mediaType string, u UnmarshalFunc) error {
	_ = "STUB: not implemented"
	return nil
}
