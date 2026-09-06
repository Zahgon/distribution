package testutil

import (
	"net/http"
)

type RequestResponseMap []RequestResponseMapping

type RequestResponseMapping struct {
	Request  Request
	Response Response
}

type Request struct {
	Method string

	Route string

	QueryParams map[string][]string

	Body []byte

	Headers http.Header
}

func (r Request) String() string { _ = "STUB: not implemented"; return "" }

type Response struct {
	StatusCode int

	Headers http.Header

	Body []byte
}

type testHandler struct {
	responseMap map[string][]Response
}

func NewHandler(requestResponseMap RequestResponseMap) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (app *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
