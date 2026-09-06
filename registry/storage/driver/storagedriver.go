package driver

import (
	"context"
	"io"
	"net/http"
	"regexp"
)

type Version string

func (version Version) Major() uint { _ = "STUB: not implemented"; return 0 }

func (version Version) Minor() uint { _ = "STUB: not implemented"; return 0 }

const CurrentVersion Version = "0.1"

type WalkOptions struct {
	StartAfterHint string
}

func WithStartAfterHint(startAfterHint string) func(*WalkOptions) {
	_ = "STUB: not implemented"
	return nil
}

type StorageDriver interface {
	Name() string

	GetContent(ctx context.Context, path string) ([]byte, error)

	PutContent(ctx context.Context, path string, content []byte) error

	Reader(ctx context.Context, path string, offset int64) (io.ReadCloser, error)

	Writer(ctx context.Context, path string, append bool) (FileWriter, error)

	Stat(ctx context.Context, path string) (FileInfo, error)

	List(ctx context.Context, path string) ([]string, error)

	Move(ctx context.Context, sourcePath string, destPath string) error

	Delete(ctx context.Context, path string) error

	RedirectURL(r *http.Request, path string) (string, error)

	Walk(ctx context.Context, path string, f WalkFn, options ...func(*WalkOptions)) error
}

type FileWriter interface {
	io.WriteCloser

	Size() int64

	Cancel(context.Context) error

	Commit(context.Context) error
}

var PathRegexp = regexp.MustCompile(`^(/[A-Za-z0-9._-]+)+$`)

type ErrUnsupportedMethod struct {
	DriverName string
}

func (err ErrUnsupportedMethod) Error() string { _ = "STUB: not implemented"; return "" }

type PathNotFoundError struct {
	Path       string
	DriverName string
}

func (err PathNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

type InvalidPathError struct {
	Path       string
	DriverName string
}

func (err InvalidPathError) Error() string { _ = "STUB: not implemented"; return "" }

type InvalidOffsetError struct {
	Path       string
	Offset     int64
	DriverName string
}

func (err InvalidOffsetError) Error() string { _ = "STUB: not implemented"; return "" }

type Error struct {
	DriverName string
	Detail     error
}

func (err Error) Error() string { _ = "STUB: not implemented"; return "" }

func (err Error) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type Errors struct {
	DriverName string
	Errs       []error
}

var _ error = Errors{}

func (e Errors) Error() string { _ = "STUB: not implemented"; return "" }

func (e Errors) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
