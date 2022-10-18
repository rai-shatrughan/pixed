package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/gorilla/mux"

	md "microsvc/pkg/model"
	mw "microsvc/pkg/mware"
)

var (
	topic   = "ts"
	groupID = "ts-consumer-group"
	kf      = mw.KafkaWriter{}
	brokers []string
)

func init() {
	conf := mw.Config{}
	conf.New()

	brokers = conf.GetStringSlice("kafka.brokers")
	kf.GroupID = &groupID
	kf.Topic = &topic
	kf.Brokers = brokers

	kf.New()
}

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

	shutDownServer(srv)

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
	kfmsg, _ := json.Marshal(tsa)

	params := mux.Vars(r)
	id := params["id"]
	kf.Write([]byte(id), []byte(kfmsg))

	w.Header().Set("Content-Type", "application/json")
	resp := "{\"TimeseriesUpload\":\"OK\"}"

	w.Write([]byte(resp))

}

func shutDownServer(server *http.Server) {
	var wg sync.WaitGroup
	wg.Add(1)

	// This goroutine is running in parallels to the main one
	go func() {
		// creating a channel to listen for signals, like SIGINT
		stop := make(chan os.Signal, 1)
		// subscribing to interruption signals
		signal.Notify(stop, os.Interrupt)
		// this blocks until the signal is received
		<-stop
		// initiating the shutdown
		err := server.Shutdown(context.Background())
		// can't do much here except for logging any errors
		if err != nil {
			log.Printf("error during shutdown: %v\n", err)
		}
		// notifying the main goroutine that we are done
		wg.Done()
	}()

}
