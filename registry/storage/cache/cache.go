package cache

import (
	"github.com/distribution/distribution/v3"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type BlobDescriptorCacheProvider interface {
	distribution.BlobDescriptorService

	RepositoryScoped(repo string) (distribution.BlobDescriptorService, error)
}

func ValidateDescriptor(desc v1.Descriptor) error { _ = "STUB: not implemented"; return nil }
