package inmemory

import (
	"context"
	"io"
	"net/http"
	"sync"

	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/distribution/distribution/v3/registry/storage/driver/base"
	"github.com/distribution/distribution/v3/registry/storage/driver/factory"
)

const driverName = "inmemory"

func init() {
	factory.Register(driverName, &inMemoryDriverFactory{})
}

type inMemoryDriverFactory struct{}

func (factory *inMemoryDriverFactory) Create(ctx context.Context, parameters map[string]any) (storagedriver.StorageDriver, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.StorageDriver), nil
}

type driver struct {
	root  *dir
	mutex sync.RWMutex
}

type baseEmbed struct {
	base.Base
}

type Driver struct {
	baseEmbed
}

var _ storagedriver.StorageDriver = &Driver{}

func New() *Driver { _ = "STUB: not implemented"; return nil }

func (d *driver) Name() string { _ = "STUB: not implemented"; return "" }

func (d *driver) GetContent(ctx context.Context, path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *driver) PutContent(ctx context.Context, p string, contents []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *driver) Reader(ctx context.Context, path string, offset int64) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (d *driver) reader(ctx context.Context, path string, offset int64) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (d *driver) Writer(ctx context.Context, path string, append bool) (storagedriver.FileWriter, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.FileWriter), nil
}

func (d *driver) Stat(ctx context.Context, path string) (storagedriver.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.FileInfo), nil
}

func (d *driver) List(ctx context.Context, path string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *driver) Move(ctx context.Context, sourcePath string, destPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *driver) Delete(ctx context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *driver) RedirectURL(*http.Request, string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *driver) Walk(ctx context.Context, path string, f storagedriver.WalkFn, options ...func(*storagedriver.WalkOptions)) error {
	_ = "STUB: not implemented"
	return nil
}

type writer struct {
	d         *driver
	f         *file
	buffer    []byte
	buffSize  int
	closed    bool
	committed bool
	cancelled bool
}

func (d *driver) newWriter(f *file) storagedriver.FileWriter {
	_ = "STUB: not implemented"
	return *new(storagedriver.FileWriter)
}

func (w *writer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *writer) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (w *writer) Close() error { _ = "STUB: not implemented"; return nil }

func (w *writer) Cancel(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (w *writer) Commit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (w *writer) flush() error { _ = "STUB: not implemented"; return nil }
