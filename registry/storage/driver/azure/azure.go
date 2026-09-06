package azure

import (
	"bufio"
	"context"
	"errors"
	"io"
	"net/http"
	"sync/atomic"
	"time"

	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/distribution/distribution/v3/registry/storage/driver/base"
	"github.com/distribution/distribution/v3/registry/storage/driver/factory"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

func init() {
	factory.Register(driverName, &azureDriverFactory{})
}

var ErrCorruptedData = errors.New("corrupted data found in the uploaded data")

const (
	driverName   = "azure"
	maxChunkSize = 4 * 1024 * 1024
)

type azureDriverFactory struct{}

func (factory *azureDriverFactory) Create(ctx context.Context, parameters map[string]any) (storagedriver.StorageDriver, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.StorageDriver), nil
}

var _ storagedriver.StorageDriver = &driver{}

type driver struct {
	azClient      *azureClient
	client        *container.Client
	rootDirectory string
	maxRetries    int
	retryDelay    time.Duration
}

type baseEmbed struct {
	base.Base
}

type Driver struct {
	baseEmbed
}

func New(ctx context.Context, params *DriverParameters) (*Driver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *driver) Name() string { _ = "STUB: not implemented"; return "" }

func (d *driver) GetContent(ctx context.Context, path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *driver) PutContent(ctx context.Context, path string, contents []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *driver) uploadBlockBlob(ctx context.Context, blobName string, contents []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *driver) deleteLegacyPutContentBlob(ctx context.Context, blobName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *driver) Reader(ctx context.Context, path string, offset int64) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (d *driver) Writer(ctx context.Context, path string, appendMode bool) (storagedriver.FileWriter, error) {
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

func (d *driver) RedirectURL(req *http.Request, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *driver) signBlobURL(ctx context.Context, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *driver) Walk(ctx context.Context, path string, f storagedriver.WalkFn, options ...func(*storagedriver.WalkOptions)) error {
	_ = "STUB: not implemented"
	return nil
}

func directDescendants(blobs []string, prefix string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (d *driver) listBlobs(ctx context.Context, virtPath string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *driver) blobName(path string) string { _ = "STUB: not implemented"; return "" }

func is404(err error) bool { _ = "STUB: not implemented"; return false }

var _ storagedriver.FileWriter = &writer{}

type writer struct {
	driver    *driver
	path      string
	size      *atomic.Int64
	bw        *bufio.Writer
	closed    bool
	committed bool
	cancelled bool
}

func (d *driver) newWriter(ctx context.Context, path string, size int64, eTag *azcore.ETag) storagedriver.FileWriter {
	_ = "STUB: not implemented"
	return *new(storagedriver.FileWriter)
}

func (w *writer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *writer) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (w *writer) Close() error { _ = "STUB: not implemented"; return nil }

func (w *writer) Cancel(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (w *writer) Commit(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type blockWriter struct {
	client     *container.Client
	path       string
	maxRetries int32
	ctx        context.Context
	size       *atomic.Int64
	eTag       *azcore.ETag
}

func (bw *blockWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (bw *blockWriter) chunkUploadVerify(appendPos int64, chunk []byte) (int64, *azcore.ETag, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}
