package mware

import "net/http"

func ResponseWriter(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}
