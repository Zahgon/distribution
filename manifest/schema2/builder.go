package schema2

import (
	"context"

	"github.com/distribution/distribution/v3"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type Builder struct {
	configDescriptor v1.Descriptor

	configJSON []byte

	dependencies []v1.Descriptor
}

func NewManifestBuilder(configDescriptor v1.Descriptor, configJSON []byte) *Builder {
	_ = "STUB: not implemented"
	return nil
}

func (mb *Builder) Build(ctx context.Context) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}

func (mb *Builder) AppendReference(ref v1.Descriptor) error { _ = "STUB: not implemented"; return nil }

func (mb *Builder) References() []v1.Descriptor { _ = "STUB: not implemented"; return nil }
