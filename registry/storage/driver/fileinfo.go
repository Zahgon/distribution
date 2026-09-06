package driver

import "time"

type FileInfo interface {
	Path() string

	Size() int64

	ModTime() time.Time

	IsDir() bool
}

type FileInfoFields struct {
	Path string

	Size int64

	ModTime time.Time

	IsDir bool
}

type FileInfoInternal struct {
	FileInfoFields
}

var (
	_ FileInfo = FileInfoInternal{}
	_ FileInfo = &FileInfoInternal{}
)

func (fi FileInfoInternal) Path() string { _ = "STUB: not implemented"; return "" }

func (fi FileInfoInternal) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (fi FileInfoInternal) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (fi FileInfoInternal) IsDir() bool { _ = "STUB: not implemented"; return false }
