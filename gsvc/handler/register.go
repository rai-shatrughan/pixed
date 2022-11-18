package handler

import (
	"gsvc/handler/mc"
	"gsvc/handler/stream"
	"gsvc/pkg/util"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ApiHandler struct {
	*util.AppState
}

func (h *ApiHandler) New(st *util.AppState) {
	h.AppState = st
}

func (api *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := sanitizeNsplit(r.URL.Path)
	n := len(p)

	var h http.Handler

	switch {

	// /api/.../exchange/{assetId}
	case n > 0 && p[0] == "api" && p[n-2] == "exchange" && r.Method == "POST":
		h = mc.PostMixedTimeseries(api.AppState, p[n-1])

	// /api/.../timeseries/{assetId}
	case n > 0 && p[0] == "api" && p[n-2] == "timeseries" && r.Method == "GET":
		h = mc.GetTimeseries(api.AppState, p[n-1])

	// /api/.../stream
	case n > 0 && p[0] == "api" && p[n-1] == "stream":
		h = stream.StreamHandler(api.AppState)

	// /metrics
	case n > 0 && p[0] == "metrics":
		h = promhttp.Handler()

	// none of the above
	default:
		http.NotFound(w, r)
		return
	}

	h.ServeHTTP(w, r)
}

func sanitizeNsplit(path string) []string {
	p := strings.TrimPrefix(path, "/")
	p = strings.TrimSuffix(p, "/")
	return strings.Split(p, "/")
}
