package storagemiddleware

import (
	"context"

	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
)

type InitFunc func(ctx context.Context, storageDriver storagedriver.StorageDriver, options map[string]any) (storagedriver.StorageDriver, error)

var storageMiddlewares map[string]InitFunc

func Register(name string, initFunc InitFunc) error { _ = "STUB: not implemented"; return nil }

func Get(ctx context.Context, name string, options map[string]any, storageDriver storagedriver.StorageDriver) (storagedriver.StorageDriver, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.StorageDriver), nil
}
