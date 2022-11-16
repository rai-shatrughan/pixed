package stream

import (
	"net/http"

	"gsvc/pkg/util"
)

func StreamHandler(conf *util.Config) http.Handler {
	return http.StripPrefix(conf.GetString("basepath.stream"), http.FileServer(http.Dir(conf.GetString("svc.stream.dir"))))
}
