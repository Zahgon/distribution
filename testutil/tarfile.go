package testutil

import (
	"io"

	"github.com/distribution/distribution/v3"
	"github.com/opencontainers/go-digest"
)

func CreateRandomTarFile() (rs io.ReadSeeker, dgst digest.Digest, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeeker), *new(digest.Digest), nil
}

func CreateRandomLayers(n int) (map[digest.Digest]io.ReadSeeker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UploadBlobs(repository distribution.Repository, layers map[digest.Digest]io.ReadSeeker) error {
	_ = "STUB: not implemented"
	return nil
}
