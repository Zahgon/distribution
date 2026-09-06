package handlers

import (
	"net/http"

	"github.com/distribution/distribution/v3"
	"github.com/opencontainers/go-digest"
)

const (
	defaultArch         = "amd64"
	defaultOS           = "linux"
	maxManifestBodySize = 4 * 1024 * 1024
	imageClass          = "image"
)

type storageType int

const (
	manifestSchema2 storageType = iota
	manifestlistSchema
	ociSchema
	ociImageIndexSchema
	numStorageTypes
)

func manifestDispatcher(ctx *Context, r *http.Request) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type manifestHandler struct {
	*Context

	Tag    string
	Digest digest.Digest
}

func (imh *manifestHandler) GetManifest(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func etagMatch(r *http.Request, etag string) bool { _ = "STUB: not implemented"; return false }

func (imh *manifestHandler) PutManifest(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (imh *manifestHandler) applyResourcePolicy(manifest distribution.Manifest) error {
	_ = "STUB: not implemented"
	return nil
}

func (imh *manifestHandler) DeleteManifest(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
