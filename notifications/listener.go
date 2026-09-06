package notifications

import (
	"context"
	"io"
	"net/http"

	"github.com/distribution/distribution/v3"

	"github.com/distribution/reference"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type ManifestListener interface {
	ManifestPushed(repo reference.Named, sm distribution.Manifest, options ...distribution.ManifestServiceOption) error
	ManifestPulled(repo reference.Named, sm distribution.Manifest, options ...distribution.ManifestServiceOption) error
	ManifestDeleted(repo reference.Named, dgst digest.Digest) error
}

type BlobListener interface {
	BlobPushed(repo reference.Named, desc v1.Descriptor) error
	BlobPulled(repo reference.Named, desc v1.Descriptor) error
	BlobMounted(repo reference.Named, desc v1.Descriptor, fromRepo reference.Named) error
	BlobDeleted(repo reference.Named, desc digest.Digest) error
}

type RepoListener interface {
	TagDeleted(repo reference.Named, tag string) error
	RepoDeleted(repo reference.Named) error
}

type Listener interface {
	ManifestListener
	BlobListener
	RepoListener
}

type repositoryListener struct {
	distribution.Repository
	listener Listener
}

type removerListener struct {
	distribution.RepositoryRemover
	listener Listener
}

func Listen(repo distribution.Repository, remover distribution.RepositoryRemover, listener Listener) (distribution.Repository, distribution.RepositoryRemover) {
	_ = "STUB: not implemented"
	return *new(distribution.Repository), *new(distribution.RepositoryRemover)
}

func (nl *removerListener) Remove(ctx context.Context, name reference.Named) error {
	_ = "STUB: not implemented"
	return nil
}

func (rl *repositoryListener) Manifests(ctx context.Context, options ...distribution.ManifestServiceOption) (distribution.ManifestService, error) {
	_ = "STUB: not implemented"
	return *new(distribution.ManifestService), nil
}

func (rl *repositoryListener) Blobs(ctx context.Context) distribution.BlobStore {
	_ = "STUB: not implemented"
	return *new(distribution.BlobStore)
}

type manifestServiceListener struct {
	distribution.ManifestService
	parent *repositoryListener
}

func (msl *manifestServiceListener) Delete(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (msl *manifestServiceListener) Get(ctx context.Context, dgst digest.Digest, options ...distribution.ManifestServiceOption) (distribution.Manifest, error) {
	_ = "STUB: not implemented"
	return *new(distribution.Manifest), nil
}

func (msl *manifestServiceListener) Put(ctx context.Context, sm distribution.Manifest, options ...distribution.ManifestServiceOption) (digest.Digest, error) {
	_ = "STUB: not implemented"
	return *new(digest.Digest), nil
}

type blobServiceListener struct {
	distribution.BlobStore
	parent *repositoryListener
}

var _ distribution.BlobStore = &blobServiceListener{}

func (bsl *blobServiceListener) Get(ctx context.Context, dgst digest.Digest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (bsl *blobServiceListener) Open(ctx context.Context, dgst digest.Digest) (io.ReadSeekCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadSeekCloser), nil
}

func (bsl *blobServiceListener) ServeBlob(ctx context.Context, w http.ResponseWriter, r *http.Request, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (bsl *blobServiceListener) Put(ctx context.Context, mediaType string, p []byte) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

func (bsl *blobServiceListener) Create(ctx context.Context, options ...distribution.BlobCreateOption) (distribution.BlobWriter, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobWriter), nil
}

func (bsl *blobServiceListener) Delete(ctx context.Context, dgst digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

func (bsl *blobServiceListener) Resume(ctx context.Context, id string) (distribution.BlobWriter, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobWriter), nil
}

func (bsl *blobServiceListener) decorateWriter(wr distribution.BlobWriter) distribution.BlobWriter {
	_ = "STUB: not implemented"
	return *new(distribution.BlobWriter)
}

type blobWriterListener struct {
	distribution.BlobWriter
	parent *blobServiceListener
}

func (bwl *blobWriterListener) Commit(ctx context.Context, desc v1.Descriptor) (v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(v1.Descriptor), nil
}

type tagServiceListener struct {
	distribution.TagService
	parent *repositoryListener
}

func (rl *repositoryListener) Tags(ctx context.Context) distribution.TagService {
	_ = "STUB: not implemented"
	return *new(distribution.TagService)
}

func (tagSL *tagServiceListener) Untag(ctx context.Context, tag string) error {
	_ = "STUB: not implemented"
	return nil
}
