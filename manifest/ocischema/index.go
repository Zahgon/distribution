package ocischema

import (
	"fmt"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/manifest"
	"github.com/opencontainers/image-spec/specs-go"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

//nolint:staticcheck // ignore SA1019: manifest.Versioned is deprecated:
var IndexSchemaVersion = manifest.Versioned{
	SchemaVersion: 2,
	MediaType:     v1.MediaTypeImageIndex,
}

func init() {
	if err := distribution.RegisterManifestSchema(v1.MediaTypeImageIndex, unmarshalImageIndex); err != nil {
		panic(fmt.Sprintf("Unable to register OCI Image Index: %s", err))
	}
}

func unmarshalImageIndex(b []byte) (distribution.Manifest, v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), *new(v1.Descriptor), nil
}

type ImageIndex struct {
	specs.Versioned

	MediaType string `json:"mediaType,omitempty"`

	Manifests []v1.Descriptor `json:"manifests"`

	Annotations map[string]string `json:"annotations,omitempty"`
}

func (ii ImageIndex) References() []v1.Descriptor { _ = "STUB: not implemented"; return nil }

type DeserializedImageIndex struct {
	ImageIndex

	canonical []byte
}

func FromDescriptors(descriptors []v1.Descriptor, annotations map[string]string) (*DeserializedImageIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromDescriptorsWithMediaType(descriptors []v1.Descriptor, annotations map[string]string, mediaType string) (_ *DeserializedImageIndex, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *DeserializedImageIndex) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *DeserializedImageIndex) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m DeserializedImageIndex) Payload() (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func validateIndex(b []byte) error { _ = "STUB: not implemented"; return nil }
