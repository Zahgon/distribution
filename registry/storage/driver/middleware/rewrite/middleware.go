package middleware

import (
	"context"
	"net/http"

	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	storagemiddleware "github.com/distribution/distribution/v3/registry/storage/driver/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := storagemiddleware.Register("rewrite", newRewriteStorageMiddleware); err != nil {
		logrus.Errorf("failed to register rewrite storage middleware: %v", err)
	}
}

type rewriteStorageMiddleware struct {
	storagedriver.StorageDriver
	overrideScheme string
	overrideHost   string
	trimPathPrefix string
}

var _ storagedriver.StorageDriver = &rewriteStorageMiddleware{}

func getStringOption(key string, options map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func newRewriteStorageMiddleware(ctx context.Context, sd storagedriver.StorageDriver, options map[string]any) (storagedriver.StorageDriver, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.StorageDriver), nil
}

func (r *rewriteStorageMiddleware) RedirectURL(req *http.Request, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
