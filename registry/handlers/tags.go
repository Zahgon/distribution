package handlers

import (
	"net/http"
)

func tagsDispatcher(ctx *Context, r *http.Request) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type tagsHandler struct {
	*Context
}

type tagsAPIResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func (th *tagsHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
