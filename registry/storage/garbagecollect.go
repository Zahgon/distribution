package storage

import (
	"context"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/opencontainers/go-digest"
)

func emit(format string, a ...any) { _ = "STUB: not implemented"; return }

type GCOpts struct {
	DryRun         bool
	RemoveUntagged bool
	Quiet          bool
}

type ManifestDel struct {
	Name   string
	Digest digest.Digest
	Tags   []string
}

func MarkAndSweep(ctx context.Context, storageDriver driver.StorageDriver, registry distribution.Namespace, opts GCOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func unmarkReferencedManifest(manifestArr []ManifestDel, markSet map[digest.Digest]struct{}, quietOutput bool) []ManifestDel {
	_ = "STUB: not implemented"
	return nil
}

func markManifestReferences(dgst digest.Digest, manifestService distribution.ManifestService, ctx context.Context, ingester func(digest.Digest) bool) error {
	_ = "STUB: not implemented"
	return nil
}
