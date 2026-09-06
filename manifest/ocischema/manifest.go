package ocischema

import (
	"fmt"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/manifest"
	"github.com/opencontainers/image-spec/specs-go"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

//nolint:staticcheck // ignore SA1019: manifest.Versioned is deprecated:
var SchemaVersion = manifest.Versioned{
	SchemaVersion: 2,
	MediaType:     v1.MediaTypeImageManifest,
}

func init() {
	if err := distribution.RegisterManifestSchema(v1.MediaTypeImageManifest, unmarshalOCISchema); err != nil {
		panic(fmt.Sprintf("Unable to register manifest: %s", err))
	}
}

func unmarshalOCISchema(b []byte) (distribution.Manifest, v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), *new(v1.Descriptor), nil
}

type Manifest struct {
	specs.Versioned

	MediaType string `json:"mediaType,omitempty"`

	Config v1.Descriptor `json:"config"`

	Layers []v1.Descriptor `json:"layers"`

	Annotations map[string]string `json:"annotations,omitempty"`
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

func (m *DeserializedManifest) Payload() (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func validateManifest(b []byte) error { _ = "STUB: not implemented"; return nil }
