package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const apiPath = "/"
	const songsDir = "/data/asset"
	const port = 8003
	const host = "0.0.0.0"

	http.Handle(apiPath, addHeaders(http.FileServer(http.Dir(songsDir))))
	fmt.Printf("Starting server on %s:%v\n", host, port)
	log.Printf("Serving %s on HTTP port: %s:%v\n", songsDir, host, port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", host, port), nil))
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
