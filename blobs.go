package distribution

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/distribution/reference"
	"github.com/opencontainers/go-digest"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

var (
	ErrBlobExists = errors.New("blob exists")

	ErrBlobDigestUnsupported = errors.New("unsupported blob digest")

	ErrBlobUnknown = errors.New("unknown blob")

	ErrBlobUploadUnknown = errors.New("blob upload unknown")

	ErrBlobInvalidLength = errors.New("blob invalid length")
)

type ErrBlobInvalidDigest struct {
	Digest digest.Digest
	Reason error
}

func (err ErrBlobInvalidDigest) Error() string { _ = "STUB: not implemented"; return "" }

type ErrBlobMounted struct {
	From       reference.Canonical
	Descriptor v1.Descriptor
}

func (err ErrBlobMounted) Error() string { _ = "STUB: not implemented"; return "" }

type Descriptor = v1.Descriptor

type BlobStatter interface {
	Stat(ctx context.Context, dgst digest.Digest) (v1.Descriptor, error)
}

type BlobDeleter interface {
	Delete(ctx context.Context, dgst digest.Digest) error
}

type BlobEnumerator interface {
	Enumerate(ctx context.Context, ingester func(dgst digest.Digest) error) error
}

type BlobDescriptorService interface {
	BlobStatter

	SetDescriptor(ctx context.Context, dgst digest.Digest, desc v1.Descriptor) error

	Clear(ctx context.Context, dgst digest.Digest) error
}

type BlobDescriptorServiceFactory interface {
	BlobAccessController(svc BlobDescriptorService) BlobDescriptorService
}

type BlobProvider interface {
	Get(ctx context.Context, dgst digest.Digest) ([]byte, error)

	Open(ctx context.Context, dgst digest.Digest) (io.ReadSeekCloser, error)
}

type BlobServer interface {
	ServeBlob(ctx context.Context, w http.ResponseWriter, r *http.Request, dgst digest.Digest) error
}

type BlobIngester interface {
	Put(ctx context.Context, mediaType string, p []byte) (v1.Descriptor, error)

	Create(ctx context.Context, options ...BlobCreateOption) (BlobWriter, error)

	Resume(ctx context.Context, id string) (BlobWriter, error)
}

type BlobCreateOption interface {
	Apply(any) error
}

type CreateOptions struct {
	Mount struct {
		ShouldMount bool
		From        reference.Canonical

		Stat *v1.Descriptor
	}
}

type BlobWriter interface {
	io.WriteCloser
	io.ReaderFrom

	Size() int64

	ID() string

	StartedAt() time.Time

	Commit(ctx context.Context, provisional v1.Descriptor) (canonical v1.Descriptor, err error)

	Cancel(ctx context.Context) error
}

type BlobService interface {
	BlobStatter
	BlobProvider
	BlobIngester
}

type BlobStore interface {
	BlobService
	BlobServer
	BlobDeleter
}
