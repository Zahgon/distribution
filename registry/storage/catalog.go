package storage

import (
	"context"

	"github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/distribution/reference"
)

func (reg *registry) Repositories(ctx context.Context, repos []string, last string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (reg *registry) Enumerate(ctx context.Context, ingester func(string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (reg *registry) Remove(ctx context.Context, name reference.Named) error {
	_ = "STUB: not implemented"
	return nil
}

func lessPath(a, b string) bool { _ = "STUB: not implemented"; return false }

func compareReplaceInline(s1, s2 string, old, new byte) int { _ = "STUB: not implemented"; return 0 }

func handleRepository(fileInfo driver.FileInfo, root, last string, fn func(repoPath string) error) error {
	_ = "STUB: not implemented"
	return nil
}
