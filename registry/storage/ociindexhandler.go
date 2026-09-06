package storage

import (
	"context"

	"github.com/distribution/distribution/v3"
	"github.com/opencontainers/go-digest"
)

type ocischemaIndexHandler struct {
	*manifestListHandler
}

var _ ManifestHandler = &manifestListHandler{}

func (ms *ocischemaIndexHandler) Unmarshal(ctx context.Context, dgst digest.Digest, content []byte) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}
