package manifestlist

import (
	"fmt"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/manifest"
	"github.com/opencontainers/image-spec/specs-go"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

const (
	MediaTypeManifestList = "application/vnd.docker.distribution.manifest.list.v2+json"
)

//nolint:staticcheck // ignore SA1019: manifest.Versioned is deprecated:
var SchemaVersion = manifest.Versioned{
	SchemaVersion: 2,
	MediaType:     MediaTypeManifestList,
}

func init() {
	if err := distribution.RegisterManifestSchema(MediaTypeManifestList, unmarshalManifestList); err != nil {
		panic(fmt.Sprintf("Unable to register manifest: %s", err))
	}
}

func unmarshalManifestList(b []byte) (distribution.Manifest, v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), *new(v1.Descriptor), nil
}

type PlatformSpec struct {
	Architecture string `json:"architecture"`

	OS string `json:"os"`

	OSVersion string `json:"os.version,omitempty"`

	OSFeatures []string `json:"os.features,omitempty"`

	Variant string `json:"variant,omitempty"`

	Features []string `json:"features,omitempty"`
}

type ManifestDescriptor struct {
	v1.Descriptor

	Platform PlatformSpec `json:"platform"`
}

type ManifestList struct {
	specs.Versioned

	MediaType string `json:"mediaType,omitempty"`

	Manifests []ManifestDescriptor `json:"manifests"`
}

func (m ManifestList) References() []v1.Descriptor { _ = "STUB: not implemented"; return nil }

type DeserializedManifestList struct {
	ManifestList

	canonical []byte
}

func FromDescriptors(descriptors []ManifestDescriptor) (*DeserializedManifestList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromDescriptorsWithMediaType(descriptors []ManifestDescriptor, mediaType string) (*DeserializedManifestList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *DeserializedManifestList) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *DeserializedManifestList) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m DeserializedManifestList) Payload() (string, []byte, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func validateManifestList(b []byte) error { _ = "STUB: not implemented"; return nil }
