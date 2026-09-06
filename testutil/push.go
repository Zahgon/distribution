package testutil

import (
	"context"
	"io"

	"github.com/distribution/distribution/v3"
	"github.com/opencontainers/go-digest"
)

func PushBlob(ctx context.Context, repository distribution.Repository, blobReader io.ReadSeeker, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}
