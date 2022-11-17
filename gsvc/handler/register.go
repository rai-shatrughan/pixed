package handler

import (
	"gsvc/handler/mc"
	"gsvc/handler/stream"
	"gsvc/pkg/util"
	"net/http"
	"regexp"
	"strconv"
	"sync"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type ApiHandler struct {
	*util.AppState
}

func (h *ApiHandler) New(st *util.AppState) {
	h.AppState = st
}

func (api *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var h http.Handler

	p := r.URL.Path
	switch {
	case match(p, "/api/mindconnect/v3/exchange/([^/]+)", &api.Params[0]) && r.Method == "POST":
		h = mc.PostMixedTimeseries(api.AppState)
	case match(p, "/api/iottimeseries/v3/timeseries/([^/]+)", &api.Params[0]) && r.Method == "GET":
		h = mc.GetTimeseries(api.AppState)
	case match(p, "/api/v1/stream"):
		h = stream.StreamHandler(api.AppState)
	case match(p, "/metrics"):
		h = promhttp.Handler()
	default:
		http.NotFound(w, r)
		return
	}

	h.ServeHTTP(w, r)
}

func match(path, pattern string, vars ...interface{}) bool {
	regex := mustCompileCached(pattern)
	matches := regex.FindStringSubmatch(path)
	if len(matches) <= 0 {
		return false
	}
	for i, match := range matches[1:] {
		switch p := vars[i].(type) {
		case *string:
			*p = match
		case *int:
			n, err := strconv.Atoi(match)
			if err != nil {
				return false
			}
			*p = n
		default:
			panic("vars must be *string or *int")
		}
	}
	return true
}

var (
	regexen = make(map[string]*regexp.Regexp)
	relock  sync.Mutex
)

func mustCompileCached(pattern string) *regexp.Regexp {
	relock.Lock()
	defer relock.Unlock()

	regex := regexen[pattern]
	if regex == nil {
		regex = regexp.MustCompile("^" + pattern + "$")
		regexen[pattern] = regex
	}
	return regex
}
