package checks

import (
	"net/http"
	"time"

	"github.com/distribution/distribution/v3/health"
)

func FileChecker(f string) health.Checker { _ = "STUB: not implemented"; return *new(health.Checker) }

func HTTPChecker(r string, statusCode int, timeout time.Duration, headers http.Header) health.Checker {
	_ = "STUB: not implemented"
	return *new(health.Checker)
}

func TCPChecker(addr string, timeout time.Duration) health.Checker {
	_ = "STUB: not implemented"
	return *new(health.Checker)
}
