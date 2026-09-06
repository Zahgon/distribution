package testutil

import (
	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/manifest/manifestlist"
	"github.com/opencontainers/go-digest"
)

func MakeManifestList(blobstatter distribution.BlobStatter, manifestDigests []digest.Digest) (*manifestlist.DeserializedManifestList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MakeSchema2Manifest(repository distribution.Repository, digests []digest.Digest) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}

func MakeOCIManifest(repository distribution.Repository, digests []digest.Digest) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}
