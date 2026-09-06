package storage

import (
	"context"

	"github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/opencontainers/go-digest"
)

func NewVacuum(ctx context.Context, driver driver.StorageDriver) Vacuum {
	_ = "STUB: not implemented"
	return *new(Vacuum)
}

type Vacuum struct {
	driver driver.StorageDriver
	ctx    context.Context
}

func (v Vacuum) RemoveBlob(dgst string) error { _ = "STUB: not implemented"; return nil }

func (v Vacuum) RemoveManifest(name string, dgst digest.Digest, tags []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (v Vacuum) RemoveRepository(repoName string) error { _ = "STUB: not implemented"; return nil }

func (v Vacuum) RemoveLayer(repoName string, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}
