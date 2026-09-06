package schema2

import (
	"fmt"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/manifest"
	"github.com/opencontainers/image-spec/specs-go"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

const (
	MediaTypeManifest = "application/vnd.docker.distribution.manifest.v2+json"

	MediaTypeImageConfig = "application/vnd.docker.container.image.v1+json"

	MediaTypePluginConfig = "application/vnd.docker.plugin.v1+json"

	MediaTypeLayer = "application/vnd.docker.image.rootfs.diff.tar.gzip"

	MediaTypeForeignLayer = "application/vnd.docker.image.rootfs.foreign.diff.tar.gzip"

	MediaTypeUncompressedLayer = "application/vnd.docker.image.rootfs.diff.tar"
)

const (
	defaultSchemaVersion = 2
	defaultMediaType     = MediaTypeManifest
)

//nolint:staticcheck // ignore SA1019: manifest.Versioned is deprecated:
var SchemaVersion = manifest.Versioned{
	SchemaVersion: defaultSchemaVersion,
	MediaType:     defaultMediaType,
}

func init() {
	if err := distribution.RegisterManifestSchema(defaultMediaType, unmarshalSchema2); err != nil {
		panic(fmt.Sprintf("Unable to register manifest: %s", err))
	}
}

func unmarshalSchema2(b []byte) (distribution.Manifest, v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), *new(v1.Descriptor), nil
}

type Manifest struct {
	specs.Versioned

	MediaType string `json:"mediaType,omitempty"`

	Config v1.Descriptor `json:"config"`

	Layers []v1.Descriptor `json:"layers"`
}

func (m Manifest) References() []v1.Descriptor { _ = "STUB: not implemented"; return nil }

func (m Manifest) Target() v1.Descriptor { _ = "STUB: not implemented"; return *new(v1.Descriptor) }

type DeserializedManifest struct {
	Manifest

	canonical []byte
}

func FromStruct(m Manifest) (*DeserializedManifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *DeserializedManifest) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

func (m *DeserializedManifest) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m DeserializedManifest) Payload() (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
