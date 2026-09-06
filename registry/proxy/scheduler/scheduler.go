package scheduler

import (
	"context"
	"sync"
	"time"

	"github.com/distribution/distribution/v3/registry/storage/driver"
	"github.com/distribution/reference"
)

type expiryFunc func(reference.Reference) error

const (
	entryTypeBlob = iota
	entryTypeManifest
	indexSaveFrequency = 5 * time.Second
)

type schedulerEntry struct {
	Key       string    `json:"Key"`
	Expiry    time.Time `json:"ExpiryData"`
	EntryType int       `json:"EntryType"`

	timer *time.Timer
}

func New(ctx context.Context, driver driver.StorageDriver, path string) *TTLExpirationScheduler {
	_ = "STUB: not implemented"
	return nil
}

type TTLExpirationScheduler struct {
	sync.Mutex

	entries map[string]*schedulerEntry

	driver          driver.StorageDriver
	ctx             context.Context
	pathToStateFile string

	stopped bool

	onBlobExpire     expiryFunc
	onManifestExpire expiryFunc

	indexDirty bool
	saveTimer  *time.Ticker
	doneChan   chan struct{}
}

func (ttles *TTLExpirationScheduler) OnBlobExpire(f expiryFunc) { _ = "STUB: not implemented"; return }

func (ttles *TTLExpirationScheduler) OnManifestExpire(f expiryFunc) {
	_ = "STUB: not implemented"
	return
}

func (ttles *TTLExpirationScheduler) AddBlob(blobRef reference.Canonical, ttl time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (ttles *TTLExpirationScheduler) AddManifest(manifestRef reference.Canonical, ttl time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (ttles *TTLExpirationScheduler) Start() error { _ = "STUB: not implemented"; return nil }

func (ttles *TTLExpirationScheduler) add(r reference.Reference, ttl time.Duration, eType int) {
	_ = "STUB: not implemented"
	return
}

func (ttles *TTLExpirationScheduler) startTimer(entry *schedulerEntry, ttl time.Duration) *time.Timer {
	_ = "STUB: not implemented"
	return nil
}

func (ttles *TTLExpirationScheduler) Stop() error { _ = "STUB: not implemented"; return nil }

func (ttles *TTLExpirationScheduler) writeState() error { _ = "STUB: not implemented"; return nil }

func (ttles *TTLExpirationScheduler) readState() error { _ = "STUB: not implemented"; return nil }
