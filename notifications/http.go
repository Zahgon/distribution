package notifications

import (
	"net/http"
	"sync"
	"time"

	events "github.com/docker/go-events"
)

type httpSink struct {
	url string

	mu        sync.Mutex
	closed    bool
	client    *http.Client
	listeners []httpStatusListener
}

func newHTTPSink(u string, timeout time.Duration, headers http.Header, transport *http.Transport, listeners ...httpStatusListener) *httpSink {
	_ = "STUB: not implemented"
	return nil
}

type httpStatusListener interface {
	success(status int, event events.Event)
	failure(status int, events events.Event)
	err(err error, events events.Event)
}

func (hs *httpSink) Write(event events.Event) error { _ = "STUB: not implemented"; return nil }

func (hs *httpSink) Close() error { _ = "STUB: not implemented"; return nil }

func (hs *httpSink) String() string { _ = "STUB: not implemented"; return "" }

type headerRoundTripper struct {
	*http.Transport
	headers http.Header
}

func (hrt *headerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
