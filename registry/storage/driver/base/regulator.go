package base

import (
	"context"
	"io"
	"net/http"
	"sync"

	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
)

type regulator struct {
	storagedriver.StorageDriver
	*sync.Cond

	available uint64
}

func GetLimitFromParameter(param any, min, def uint64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func NewRegulator(driver storagedriver.StorageDriver, limit uint64) storagedriver.StorageDriver {
	_ = "STUB: not implemented"
	return *new(storagedriver.StorageDriver)
}

func (r *regulator) enter() { _ = "STUB: not implemented"; return }

func (r *regulator) exit() { _ = "STUB: not implemented"; return }

func (r *regulator) Name() string { _ = "STUB: not implemented"; return "" }

func (r *regulator) GetContent(ctx context.Context, path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *regulator) PutContent(ctx context.Context, path string, content []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *regulator) Reader(ctx context.Context, path string, offset int64) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (r *regulator) Writer(ctx context.Context, path string, append bool) (storagedriver.FileWriter, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.FileWriter), nil
}

func (r *regulator) Stat(ctx context.Context, path string) (storagedriver.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.FileInfo), nil
}

func (r *regulator) List(ctx context.Context, path string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *regulator) Move(ctx context.Context, sourcePath string, destPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *regulator) Delete(ctx context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *regulator) RedirectURL(req *http.Request, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
