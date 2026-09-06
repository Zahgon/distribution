package factory

import (
	"context"

	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
)

var driverFactories = make(map[string]StorageDriverFactory)

type StorageDriverFactory interface {
	Create(ctx context.Context, parameters map[string]any) (storagedriver.StorageDriver, error)
}

func Register(name string, factory StorageDriverFactory) { _ = "STUB: not implemented"; return }

func Create(ctx context.Context, name string, parameters map[string]any) (storagedriver.StorageDriver, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.StorageDriver), nil
}

type InvalidStorageDriverError struct {
	Name string
}

func (err InvalidStorageDriverError) Error() string { _ = "STUB: not implemented"; return "" }
