package inmemory

import (
	"fmt"
	"io"
	"time"
)

var (
	errExists    = fmt.Errorf("exists")
	errNotExists = fmt.Errorf("notexists")
	errIsNotDir  = fmt.Errorf("notdir")
	errIsDir     = fmt.Errorf("isdir")
)

type node interface {
	name() string
	path() string
	isdir() bool
	modtime() time.Time
}

type dir struct {
	common

	children map[string]node
}

var _ node = &dir{}

func (d *dir) isdir() bool { _ = "STUB: not implemented"; return false }

func (d *dir) add(n node) { _ = "STUB: not implemented"; return }

func (d *dir) find(q string) node { _ = "STUB: not implemented"; return *new(node) }

func (d *dir) list(p string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *dir) mkfile(p string) (*file, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *dir) mkdirs(p string) (*dir, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *dir) mkdir(name string) (*dir, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *dir) move(src, dst string) error { _ = "STUB: not implemented"; return nil }

func (d *dir) delete(p string) error { _ = "STUB: not implemented"; return nil }

func (d *dir) String() string { _ = "STUB: not implemented"; return "" }

type file struct {
	common
	data []byte
}

var _ node = &file{}

func (f *file) isdir() bool { _ = "STUB: not implemented"; return false }

func (f *file) truncate() { _ = "STUB: not implemented"; return }

func (f *file) sectionReader(offset int64) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (f *file) ReadAt(p []byte, offset int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

const reallocExponent = 1.25

func (f *file) WriteAt(p []byte, offset int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *file) String() string { _ = "STUB: not implemented"; return "" }

type common struct {
	p   string
	mod time.Time
}

func (c *common) name() string { _ = "STUB: not implemented"; return "" }

func (c *common) path() string { _ = "STUB: not implemented"; return "" }

func (c *common) modtime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func normalize(p string) string { _ = "STUB: not implemented"; return "" }
