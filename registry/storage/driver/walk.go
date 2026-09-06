package driver

import (
	"context"
	"errors"
)

var ErrSkipDir = errors.New("skip this directory")

var ErrFilledBuffer = errors.New("we have enough entries")

type WalkFn func(fileInfo FileInfo) error

func WalkFallback(ctx context.Context, driver StorageDriver, from string, f WalkFn, options ...func(*WalkOptions)) error {
	_ = "STUB: not implemented"
	return nil
}

func doWalkFallback(ctx context.Context, driver StorageDriver, from string, startAfterHint string, f WalkFn) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
