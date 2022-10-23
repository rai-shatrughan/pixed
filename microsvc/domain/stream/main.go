package stream

import (
	"net/http"

	"microsvc/pkg/util"

	"github.com/gorilla/mux"
)

func StreamHandler(conf *util.Config) http.Handler {

	r := mux.NewRouter()
	r.PathPrefix(conf.GetString("basepath.stream")).Handler(http.StripPrefix(conf.GetString("basepath.stream"), http.FileServer(http.Dir(conf.GetString("svc.stream.dir")))))

	return r
}
