package api

import (
	"net/http"

	"github.com/distribution/distribution/v3/health"
)

var updater = health.NewStatusUpdater()

func init() {
	health.Register("manual_http_status", updater)
	http.HandleFunc("/debug/health/down", DownHandler)
	http.HandleFunc("/debug/health/up", UpHandler)
}

func DownHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func UpHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
