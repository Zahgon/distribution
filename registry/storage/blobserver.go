package storage

import (
	"context"
	"net/http"
	"time"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/opencontainers/go-digest"
)

const blobCacheControlMaxAge = 365 * 24 * time.Hour

type blobServer struct {
	driver   driver.StorageDriver
	statter  distribution.BlobStatter
	pathFn   func(dgst digest.Digest) (string, error)
	redirect bool
}

func (bs *blobServer) ServeBlob(ctx context.Context, w http.ResponseWriter, r *http.Request, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}
