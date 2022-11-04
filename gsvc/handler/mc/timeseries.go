package mc

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"gsvc/mware"
	"gsvc/pkg/model"
	"gsvc/pkg/util"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var (
	tracer           = otel.Tracer("mux-server")
	assetErrMsg      = "assetId missing in path"
	processErrMsg    = "error in processing, retry after sometime"
	uploadSuccessMsg = "{\"TimeseriesUpload\": \"Ok\"}"
)

func PostMixedTimeseries(kf *util.KafkaWriter, logger *util.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tsa model.TimeseriesArray
		vars := mux.Vars(r)
		assetId, ok := vars["assetId"]

		_, span := tracer.Start(r.Context(), "postTS", oteltrace.WithAttributes(attribute.String("assetId", assetId)))
		defer span.End()

		if !ok {
			mware.ResponseWriter(w, assetErrMsg, http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Sugar().Errorln("Error in reading Multipart TS body", err)
			mware.ResponseWriter(w, "", 500)
		}

		jsonWmp := strings.Split(string(body), "Content-Type: application/json")[1]
		jsonBody := strings.Split(jsonWmp, "--")[0]

		if err := json.Unmarshal([]byte(jsonBody), &tsa); err != nil {
			mware.ResponseWriter(w, processErrMsg, http.StatusInternalServerError)
			logger.Error(err.Error())
		}

		kfmsg := []byte(jsonBody)

		if err := kf.Write([]byte(assetId), []byte(kfmsg)); err != nil {
			mware.ResponseWriter(w, processErrMsg, http.StatusInternalServerError)
			logger.Error(err.Error())
			return
		} else {
			mware.ResponseWriter(w, uploadSuccessMsg, http.StatusOK)
			return
		}
	})
}

func PostTimeseries(kf *util.KafkaWriter, logger *util.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body model.TimeseriesArray
		vars := mux.Vars(r)
		assetId, ok := vars["assetId"]

		_, span := tracer.Start(r.Context(), "postTS", oteltrace.WithAttributes(attribute.String("assetId", assetId)))
		defer span.End()

		if !ok {
			mware.ResponseWriter(w, assetErrMsg, http.StatusBadRequest)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			mware.ResponseWriter(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := body.Validate(); err != nil {
			mware.ResponseWriter(w, err.Error(), http.StatusBadRequest)
			return
		}

		kfmsg, _ := json.Marshal(body)

		if err := kf.Write([]byte(assetId), []byte(kfmsg)); err != nil {
			mware.ResponseWriter(w, processErrMsg, http.StatusInternalServerError)
			logger.Error(err.Error())
			return
		} else {
			mware.ResponseWriter(w, uploadSuccessMsg, http.StatusOK)
			return
		}
	})
}

func GetTimeseries(kv *util.KV, logger *util.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		assetId, ok := vars["assetId"]
		if !ok {
			mware.ResponseWriter(w, assetErrMsg, http.StatusBadRequest)
			return
		}

		res, err := kv.GetKeyWithLimit(assetId, 1000)
		// res, err := kv.GetKeyDefLimit(assetId)
		// res, err := kv.Get(assetId + "_" + "2022-10-01T08:30:03.780Z")
		if err != nil {
			mware.ResponseWriter(w, processErrMsg, http.StatusInternalServerError)
			logger.Error(err.Error())
		} else {
			mware.ResponseWriter(w, res, http.StatusOK)
		}

	})
}
