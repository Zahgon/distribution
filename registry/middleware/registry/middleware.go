package middleware

import (
	"context"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/registry/storage"
	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
)

type InitFunc func(ctx context.Context, registry distribution.Namespace, driver storagedriver.StorageDriver, options map[string]any) (distribution.Namespace, error)

var (
	middlewares     map[string]InitFunc
	registryoptions []storage.RegistryOption
)

func Register(name string, initFunc InitFunc) error { _ = "STUB: not implemented"; return nil }

func Get(ctx context.Context, name string, options map[string]any, registry distribution.Namespace, driver storagedriver.StorageDriver) (distribution.Namespace, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Namespace), nil
}

func RegisterOptions(options ...storage.RegistryOption) error {
	_ = "STUB: not implemented"
	return nil
}

func GetRegistryOptions() []storage.RegistryOption { _ = "STUB: not implemented"; return nil }
