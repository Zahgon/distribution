package handlers

import (
	"net/http"

	"github.com/distribution/distribution/v3"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func blobUploadDispatcher(ctx *Context, r *http.Request) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type blobUploadHandler struct {
	*Context

	UUID string

	Upload distribution.BlobWriter

	State blobUploadState
}

func (buh *blobUploadHandler) StartBlobUpload(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (buh *blobUploadHandler) GetUploadStatus(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (buh *blobUploadHandler) PatchBlobData(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (buh *blobUploadHandler) PutBlobUploadComplete(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (buh *blobUploadHandler) CancelBlobUpload(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (buh *blobUploadHandler) ResumeBlobUpload(ctx *Context, r *http.Request) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (buh *blobUploadHandler) blobUploadResponse(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (buh *blobUploadHandler) createBlobMountOption(fromRepo, mountDigest string) (distribution.BlobCreateOption, error) {
	_ = "STUB: not implemented"
	return *new(distribution.BlobCreateOption), nil
}

func (buh *blobUploadHandler) writeBlobCreatedHeaders(w http.ResponseWriter, desc v1.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}
