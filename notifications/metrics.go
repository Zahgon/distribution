package notifications

import (
	"expvar"
	"sync"

	prometheus "github.com/distribution/distribution/v3/metrics"
	events "github.com/docker/go-events"
	"github.com/docker/go-metrics"
)

var (
	eventsCounter = prometheus.NotificationsNamespace.NewLabeledCounter("events", "The number of total events", "type", "endpoint")

	pendingGauge = prometheus.NotificationsNamespace.NewLabeledGauge("pending", "The gauge of pending events in queue", metrics.Total, "endpoint")

	statusCounter = prometheus.NotificationsNamespace.NewLabeledCounter("status", "The number of status code", "code", "endpoint")
)

var endpoints struct {
	registered []*Endpoint
	mu         sync.Mutex
}

func init() {

	registry := expvar.Get("registry")

	if registry == nil {
		registry = expvar.NewMap("registry")
	}

	var notifications expvar.Map
	notifications.Init()
	notifications.Set("endpoints", expvar.Func(func() any {
		endpoints.mu.Lock()
		defer endpoints.mu.Unlock()

		var names []any
		if len(endpoints.registered) == 0 {
			return names
		}
		names = make([]any, 0, len(endpoints.registered))
		for _, v := range endpoints.registered {
			var epjson struct {
				Name string `json:"name"`
				URL  string `json:"url"`
				EndpointConfig

				Metrics EndpointMetrics
			}

			epjson.Name = v.Name()
			epjson.URL = v.URL()
			epjson.EndpointConfig = v.EndpointConfig

			v.ReadMetrics(&epjson.Metrics)

			names = append(names, epjson)
		}

		return names
	}))

	registry.(*expvar.Map).Set("notifications", &notifications)

	metrics.Register(prometheus.NotificationsNamespace)
}

type EndpointMetrics struct {
	Pending   int
	Events    int
	Successes int
	Failures  int
	Errors    int
	Statuses  map[string]int
}

type safeMetrics struct {
	EndpointName string
	EndpointMetrics
	sync.Mutex
}

func newSafeMetrics(name string) *safeMetrics { _ = "STUB: not implemented"; return nil }

func (sm *safeMetrics) httpStatusListener() httpStatusListener {
	_ = "STUB: not implemented"
	return *new(httpStatusListener)
}

func (sm *safeMetrics) eventQueueListener() eventQueueListener {
	_ = "STUB: not implemented"
	return *new(eventQueueListener)
}

type endpointMetricsHTTPStatusListener struct {
	*safeMetrics
}

var _ httpStatusListener = &endpointMetricsHTTPStatusListener{}

func (emsl *endpointMetricsHTTPStatusListener) success(status int, event events.Event) {
	_ = "STUB: not implemented"
	return
}

func (emsl *endpointMetricsHTTPStatusListener) failure(status int, event events.Event) {
	_ = "STUB: not implemented"
	return
}

func (emsl *endpointMetricsHTTPStatusListener) err(err error, event events.Event) {
	_ = "STUB: not implemented"
	return
}

type endpointMetricsEventQueueListener struct {
	*safeMetrics
}

func (eqc *endpointMetricsEventQueueListener) ingress(event events.Event) {
	_ = "STUB: not implemented"
	return
}

func (eqc *endpointMetricsEventQueueListener) egress(event events.Event) {
	_ = "STUB: not implemented"
	return
}

func register(e *Endpoint) { _ = "STUB: not implemented"; return }
