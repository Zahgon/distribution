package storage

import (
	"context"

	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/distribution/distribution/v3"
	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
)

var _ distribution.TagService = &tagStore{}

type tagStore struct {
	repository       *repository
	blobStore        *blobStore
	concurrencyLimit int
	deleteEnabled    bool
}

func (ts *tagStore) All(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ts *tagStore) Tag(ctx context.Context, tag string, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (ts *tagStore) Get(ctx context.Context, tag string) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (ts *tagStore) Untag(ctx context.Context, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ts *tagStore) linkedBlobStore(ctx context.Context, tag string) *linkedBlobStore {
	_ = "STUB: not implemented"
	return nil
}

func (ts *tagStore) Lookup(ctx context.Context, desc v1.Descriptor) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ts *tagStore) ManifestDigests(ctx context.Context, tag string) ([]digest.Digest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ts *tagStore) List(ctx context.Context, limit int, last string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func handleTag(fileInfo storagedriver.FileInfo, root, last string, fn func(tagPath string) error) error {
	_ = "STUB: not implemented"
	return nil
}
