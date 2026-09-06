package handlers

import (
	"net/http"
)

func basicAuth(r *http.Request) (username, password string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}
