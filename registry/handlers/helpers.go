package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/distribution/distribution/v3/registry/api/errcode"
)

func closeResources(handler http.Handler, closers ...io.Closer) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func copyFullPayload(ctx context.Context, responseWriter http.ResponseWriter, r *http.Request, destWriter io.Writer, limit int64, action string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseContentRange(cr string) (start int64, end int64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func toErrcodeErrors(err error) errcode.Errors {
	_ = "STUB: not implemented"
	return *new(errcode.Errors)
}
