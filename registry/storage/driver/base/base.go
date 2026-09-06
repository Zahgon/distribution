package base

import (
	"context"
	"io"
	"net/http"

	prometheus "github.com/distribution/distribution/v3/metrics"
	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/docker/go-metrics"
	"go.opentelemetry.io/otel"
)

var storageAction = prometheus.StorageNamespace.NewLabeledTimer("action", "The number of seconds that the storage action takes", "driver", "action")

var tracer = otel.Tracer("github.com/distribution/distribution/v3/registry/storage/driver/base")

func init() {
	metrics.Register(prometheus.StorageNamespace)
}

type Base struct {
	storagedriver.StorageDriver
}

func (base *Base) setDriverName(e error) error { _ = "STUB: not implemented"; return nil }

func (base *Base) GetContent(ctx context.Context, path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (base *Base) PutContent(ctx context.Context, path string, content []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (base *Base) Reader(ctx context.Context, path string, offset int64) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (base *Base) Writer(ctx context.Context, path string, append bool) (storagedriver.FileWriter, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.FileWriter), nil
}

func (base *Base) Stat(ctx context.Context, path string) (storagedriver.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.FileInfo), nil
}

func (base *Base) List(ctx context.Context, path string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (base *Base) Move(ctx context.Context, sourcePath string, destPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (base *Base) Delete(ctx context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (base *Base) RedirectURL(r *http.Request, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (base *Base) Walk(ctx context.Context, path string, f storagedriver.WalkFn, options ...func(*storagedriver.WalkOptions)) error {
	_ = "STUB: not implemented"
	return nil
}
