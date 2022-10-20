package timeseries

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	md "microsvc/pkg/model"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
)

var errMissingAssetId = errors.New("AssetId missing in request path")

func MakeHandler(ts Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	putTimeseriesHandler := kithttp.NewServer(
		makePutTimeseriesEndpoint(ts),
		decodePutTimeseriesRequest,
		encodeResponse,
		opts...,
	)
	getTimeseriesHandler := kithttp.NewServer(
		makeGetTimeseriesEndpoint(ts),
		decodeGetTimeseriesRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/api/v1/timeseries/{assetId}", putTimeseriesHandler).Methods("PUT")
	r.Handle("/api/v1/timeseries/{assetId}", getTimeseriesHandler).Methods("GET")

	return r
}

func decodePutTimeseriesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body md.TimeseriesArray
	vars := mux.Vars(r)
	id, ok := vars["assetId"]
	if !ok {
		return nil, errMissingAssetId
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return putTimeseriesRequest{assetId: id, tsa: &body}, nil
}

func decodeGetTimeseriesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["assetId"]
	if !ok {
		return nil, errMissingAssetId
	}
	return getTimeseriesRequest{assetId: id}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
