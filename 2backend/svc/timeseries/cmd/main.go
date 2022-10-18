package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	md "mware/model"
)

func main() {
	const port = "8002"
	const host = "0.0.0.0"

	router := mux.NewRouter()
	router.
		HandleFunc("/api/v1/timeseries/{id}", TSHandler).
		Methods("GET", "POST", "PUT")

	srv := &http.Server{
		Handler: router,
		Addr:    host + ":" + port,
	}
	log.Printf("Starting server on %s:%v\n", host, port)
	log.Fatal(srv.ListenAndServe())
}

func TSHandler(w http.ResponseWriter, r *http.Request) {
	var tsa md.TimeseriesArray

	err := json.NewDecoder(r.Body).Decode(&tsa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, _ := json.Marshal(tsa)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))

}
