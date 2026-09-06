package notifications

import (
	"net/http"
	"time"

	"github.com/distribution/distribution/v3/configuration"
	events "github.com/docker/go-events"
)

type EndpointConfig struct {
	Headers           http.Header
	Timeout           time.Duration
	Threshold         int
	Backoff           time.Duration
	IgnoredMediaTypes []string
	Transport         *http.Transport `json:"-"`
	Ignore            configuration.Ignore
}

func (ec *EndpointConfig) defaults() { _ = "STUB: not implemented"; return }

type Endpoint struct {
	events.Sink
	url  string
	name string

	EndpointConfig

	metrics *safeMetrics
}

func NewEndpoint(name, url string, config EndpointConfig) *Endpoint {
	_ = "STUB: not implemented"
	return nil
}

func (e *Endpoint) Name() string { _ = "STUB: not implemented"; return "" }

func (e *Endpoint) URL() string { _ = "STUB: not implemented"; return "" }

func (e *Endpoint) ReadMetrics(em *EndpointMetrics) { _ = "STUB: not implemented"; return }
