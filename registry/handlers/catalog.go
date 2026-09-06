package handlers

import (
	"net/http"
)

const defaultReturnedEntries = 100

func catalogDispatcher(ctx *Context, r *http.Request) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type catalogHandler struct {
	*Context
}

type catalogAPIResponse struct {
	Repositories []string `json:"repositories"`
}

func (ch *catalogHandler) GetCatalog(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func createLinkEntry(origURL string, maxEntries int, lastEntry string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
