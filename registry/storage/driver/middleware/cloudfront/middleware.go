package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	storagemiddleware "github.com/distribution/distribution/v3/registry/storage/driver/middleware"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := storagemiddleware.Register("cloudfront", newCloudFrontStorageMiddleware); err != nil {
		logrus.Errorf("failed to register cloudfront middleware: %v", err)
	}
}

type cloudFrontStorageMiddleware struct {
	storagedriver.StorageDriver
	awsIPs    *awsIPs
	urlSigner *sign.URLSigner
	baseURL   string
	duration  time.Duration
}

var _ storagedriver.StorageDriver = &cloudFrontStorageMiddleware{}

func newCloudFrontStorageMiddleware(ctx context.Context, storageDriver storagedriver.StorageDriver, options map[string]any) (storagedriver.StorageDriver, error) {
	_ = "STUB: not implemented"
	return *new(storagedriver.StorageDriver), nil
}

type S3BucketKeyer interface {
	S3BucketKey(path string) string
}

func (lh *cloudFrontStorageMiddleware) RedirectURL(r *http.Request, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
