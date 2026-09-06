package storage

import (
	"strings"

	"github.com/opencontainers/go-digest"
)

const (
	storagePathVersion = "v2"
	storagePathRoot    = "/docker/registry/"
)

func pathFor(spec pathSpec) (string, error) { _ = "STUB: not implemented"; return "", nil }

type pathSpec interface {
	pathSpec()
}

type manifestsPathSpec struct {
	name string
}

func (manifestsPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type manifestRevisionsPathSpec struct {
	name string
}

func (manifestRevisionsPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type manifestRevisionPathSpec struct {
	name     string
	revision digest.Digest
}

func (manifestRevisionPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type manifestRevisionLinkPathSpec struct {
	name     string
	revision digest.Digest
}

func (manifestRevisionLinkPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type manifestTagsPathSpec struct {
	name string
}

func (manifestTagsPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type manifestTagPathSpec struct {
	name string
	tag  string
}

func (manifestTagPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type manifestTagCurrentPathSpec struct {
	name string
	tag  string
}

func (manifestTagCurrentPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type manifestTagIndexPathSpec struct {
	name string
	tag  string
}

func (manifestTagIndexPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type manifestTagIndexEntryPathSpec struct {
	name     string
	tag      string
	revision digest.Digest
}

func (manifestTagIndexEntryPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type manifestTagIndexEntryLinkPathSpec struct {
	name     string
	tag      string
	revision digest.Digest
}

func (manifestTagIndexEntryLinkPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type layersPathSpec struct {
	name string
}

func (layersPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type layerLinkPathSpec struct {
	name   string
	digest digest.Digest
}

func (layerLinkPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

var blobAlgorithmReplacer = strings.NewReplacer(
	"+", "/",
	".", "/",
	";", "/",
)

type blobsPathSpec struct{}

func (blobsPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type blobPathSpec struct {
	digest digest.Digest
}

func (blobPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type blobDataPathSpec struct {
	digest digest.Digest
}

func (blobDataPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type uploadDataPathSpec struct {
	name string
	id   string
}

func (uploadDataPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type uploadStartedAtPathSpec struct {
	name string
	id   string
}

func (uploadStartedAtPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type uploadHashStatePathSpec struct {
	name   string
	id     string
	alg    digest.Algorithm
	offset int64
	list   bool
}

func (uploadHashStatePathSpec) pathSpec() { _ = "STUB: not implemented"; return }

type repositoriesRootPathSpec struct{}

func (repositoriesRootPathSpec) pathSpec() { _ = "STUB: not implemented"; return }

func digestPathComponents(dgst digest.Digest, multilevel bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func digestFromPath(digestPath string) (digest.Digest, error) {
	_ = "STUB: not implemented"
	return *new(digest.Digest), nil
}
