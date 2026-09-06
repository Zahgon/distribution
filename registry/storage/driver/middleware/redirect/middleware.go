package middleware

import (
	"context"
	"net/http"

	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	storagemiddleware "github.com/distribution/distribution/v3/registry/storage/driver/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := storagemiddleware.Register("redirect", newRedirectStorageMiddleware); err != nil {
		logrus.Errorf("failed to register redirect storage middleware: %v", err)
	}
}

type redirectStorageMiddleware struct {
	storagedriver.StorageDriver
	scheme   string
	host     string
	basePath string
}

var _ storagedriver.StorageDriver = &redirectStorageMiddleware{}

func newRedirectStorageMiddleware(ctx context.Context, sd storagedriver.StorageDriver, options map[string]any) (storagedriver.StorageDriver, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.StorageDriver), nil
}

func (r *redirectStorageMiddleware) RedirectURL(_ *http.Request, urlPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
