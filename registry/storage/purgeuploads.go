package storage

import (
	"context"
	"time"

	storageDriver "github.com/distribution/distribution/v3/registry/storage/driver"
)

type uploadData struct {
	containingDir string
	startedAt     time.Time
}

func newUploadData() uploadData { _ = "STUB: not implemented"; return *new(uploadData) }

func PurgeUploads(ctx context.Context, driver storageDriver.StorageDriver, olderThan time.Time, actuallyDelete bool) ([]string, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getOutstandingUploads(ctx context.Context, driver storageDriver.StorageDriver) (map[string]uploadData, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func uuidFromPath(path string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func readStartedAtFile(ctx context.Context, driver storageDriver.StorageDriver, path string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
