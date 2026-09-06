package filesystem

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"os"
	"time"

	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/distribution/distribution/v3/registry/storage/driver/base"
	"github.com/distribution/distribution/v3/registry/storage/driver/factory"
)

const (
	driverName           = "filesystem"
	defaultRootDirectory = "/var/lib/registry"
	defaultMaxThreads    = uint64(100)

	minThreads = uint64(25)
)

type DriverParameters struct {
	RootDirectory string
	MaxThreads    uint64
}

func init() {
	factory.Register(driverName, &filesystemDriverFactory{})
}

type filesystemDriverFactory struct{}

func (factory *filesystemDriverFactory) Create(ctx context.Context, parameters map[string]any) (storagedriver.StorageDriver, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.StorageDriver), nil
}

type driver struct {
	rootDirectory string
}

type baseEmbed struct {
	base.Base
}

type Driver struct {
	baseEmbed
}

func FromParameters(parameters map[string]any) (*Driver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fromParametersImpl(parameters map[string]any) (*DriverParameters, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(params DriverParameters) *Driver { _ = "STUB: not implemented"; return nil }

func (d *driver) Name() string { _ = "STUB: not implemented"; return "" }

func (d *driver) GetContent(ctx context.Context, path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *driver) PutContent(ctx context.Context, subPath string, contents []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func syncDir(dir string) (retErr error) { _ = "STUB: not implemented"; return nil }

func (d *driver) Reader(ctx context.Context, path string, offset int64) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (d *driver) Writer(ctx context.Context, subPath string, append bool) (storagedriver.FileWriter, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.FileWriter), nil
}

func (d *driver) Stat(ctx context.Context, subPath string) (storagedriver.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.FileInfo), nil
}

func (d *driver) List(ctx context.Context, subPath string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *driver) Move(ctx context.Context, sourcePath string, destPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *driver) Delete(ctx context.Context, subPath string) error {
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

func (d *driver) fullPath(subPath string) string { _ = "STUB: not implemented"; return "" }

type fileInfo struct {
	os.FileInfo
	path string
}

var _ storagedriver.FileInfo = fileInfo{}

func (fi fileInfo) Path() string { _ = "STUB: not implemented"; return "" }

func (fi fileInfo) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (fi fileInfo) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (fi fileInfo) IsDir() bool { _ = "STUB: not implemented"; return false }

type fileWriter struct {
	file      *os.File
	size      int64
	bw        *bufio.Writer
	closed    bool
	committed bool
	cancelled bool
}

func newFileWriter(file *os.File, size int64) *fileWriter { _ = "STUB: not implemented"; return nil }

func (fw *fileWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (fw *fileWriter) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (fw *fileWriter) Close() (retErr error) { _ = "STUB: not implemented"; return nil }

func (fw *fileWriter) Cancel(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (fw *fileWriter) Commit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
