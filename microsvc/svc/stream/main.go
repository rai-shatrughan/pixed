package stream

import (
	"net/http"

	mw "microsvc/pkg/mware"

	"github.com/gorilla/mux"
)

func StreamHandler(conf mw.Config) http.Handler {

	r := mux.NewRouter()
	r.PathPrefix(conf.GetString("basepath.stream")).Handler(http.StripPrefix(conf.GetString("basepath.stream"), http.FileServer(http.Dir(conf.GetString("svc.stream.dir")))))

	return r
}
