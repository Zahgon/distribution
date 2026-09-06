package ocischema

import (
	"context"

	"github.com/distribution/distribution/v3"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type Builder struct {
	bs distribution.BlobService

	configJSON []byte

	layers []v1.Descriptor

	annotations map[string]string

	mediaType string
}

func NewManifestBuilder(bs distribution.BlobService, configJSON []byte, annotations map[string]string) *Builder {
	_ = "STUB: not implemented"
	return nil
}

func (mb *Builder) SetMediaType(mediaType string) error { _ = "STUB: not implemented"; return nil }

func (mb *Builder) Build(ctx context.Context) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}

func (mb *Builder) AppendReference(ref v1.Descriptor) error { _ = "STUB: not implemented"; return nil }

func (mb *Builder) References() []v1.Descriptor { _ = "STUB: not implemented"; return nil }
