package stream

import (
	"net/http"

	"gsvc/pkg/util"
)

func StreamHandler(st util.AppState) http.Handler {
	return http.StripPrefix(st.Conf.GetString("basepath.stream"), http.FileServer(http.Dir(st.Conf.GetString("svc.stream.dir"))))
}
