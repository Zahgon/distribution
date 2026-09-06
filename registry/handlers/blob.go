package handlers

import (
	"net/http"

	"github.com/opencontainers/go-digest"
)

func blobDispatcher(ctx *Context, r *http.Request) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type blobHandler struct {
	*Context

	Digest digest.Digest
}

func (bh *blobHandler) GetBlob(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (bh *blobHandler) DeleteBlob(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
