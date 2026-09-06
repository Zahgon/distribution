package proxy

import (
	"expvar"

	prometheus "github.com/distribution/distribution/v3/metrics"
	"github.com/docker/go-metrics"
)

var (
	requests = prometheus.ProxyNamespace.NewLabeledCounter("requests", "The number of total incoming proxy request received", "type")

	hits = prometheus.ProxyNamespace.NewLabeledCounter("hits", "The number of total proxy request hits", "type")

	misses = prometheus.ProxyNamespace.NewLabeledCounter("misses", "The number of total proxy request misses", "type")

	pulledBytes = prometheus.ProxyNamespace.NewLabeledCounter("pulled_bytes", "The size of total bytes pulled from the upstream", "type")

	pushedBytes = prometheus.ProxyNamespace.NewLabeledCounter("pushed_bytes", "The size of total bytes pushed to the client", "type")
)

type Metrics struct {
	Requests    uint64
	Hits        uint64
	Misses      uint64
	BytesPulled uint64
	BytesPushed uint64
}

type proxyMetricsCollector struct {
	blobMetrics     Metrics
	manifestMetrics Metrics
}

var proxyMetrics = &proxyMetricsCollector{}

func init() {
	registry := expvar.Get("registry")
	if registry == nil {
		registry = expvar.NewMap("registry")
	}

	pm := registry.(*expvar.Map).Get("proxy")
	if pm == nil {
		pm = &expvar.Map{}
		pm.(*expvar.Map).Init()
		registry.(*expvar.Map).Set("proxy", pm)
	}

	pm.(*expvar.Map).Set("blobs", expvar.Func(func() any {
		return proxyMetrics.blobMetrics
	}))

	pm.(*expvar.Map).Set("manifests", expvar.Func(func() any {
		return proxyMetrics.manifestMetrics
	}))

	metrics.Register(prometheus.ProxyNamespace)
	initPrometheusMetrics("blob")
	initPrometheusMetrics("manifest")
}

func initPrometheusMetrics(value string) { _ = "STUB: not implemented"; return }

func (pmc *proxyMetricsCollector) BlobPull(bytesPulled uint64) { _ = "STUB: not implemented"; return }

func (pmc *proxyMetricsCollector) BlobPush(bytesPushed uint64, isHit bool) {
	_ = "STUB: not implemented"
	return
}

func (pmc *proxyMetricsCollector) ManifestPull(bytesPulled uint64) {
	_ = "STUB: not implemented"
	return
}

func (pmc *proxyMetricsCollector) ManifestPush(bytesPushed uint64, isHit bool) {
	_ = "STUB: not implemented"
	return
}
