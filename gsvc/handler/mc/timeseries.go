package mc

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"gsvc/mware"
	"gsvc/pkg/model"
	"gsvc/pkg/util"

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

func PostMixedTimeseries(st *util.AppState) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tsa model.TimeseriesArray
		// assetId := strings.TrimPrefix(r.URL.Path, st.Conf.GetString("basepath.timeseries"))
		assetId := st.Params[0]

		_, span := tracer.Start(r.Context(), "postTS", oteltrace.WithAttributes(attribute.String("assetId", assetId)))
		defer span.End()

		if assetId == "" {
			mware.ResponseWriter(w, assetErrMsg, http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			st.Logger.Sugar().Errorln("Error in reading Multipart TS body", err)
			mware.ResponseWriter(w, "", 500)
		}

		jsonWmp := strings.Split(string(body), "Content-Type: application/json")[1]
		jsonBody := strings.Split(jsonWmp, "--")[0]

		if err := json.Unmarshal([]byte(jsonBody), &tsa); err != nil {
			mware.ResponseWriter(w, processErrMsg, http.StatusInternalServerError)
			st.Logger.Error(err.Error())
		}

		kfmsg := []byte(jsonBody)

		if err := st.Kfw.Write([]byte(assetId), []byte(kfmsg)); err != nil {
			mware.ResponseWriter(w, processErrMsg, http.StatusInternalServerError)
			st.Logger.Error(err.Error())
			return
		} else {
			mware.ResponseWriter(w, uploadSuccessMsg, http.StatusOK)
			return
		}
	})
}

func GetTimeseries(st *util.AppState) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// assetId := strings.TrimPrefix(r.URL.Path, st.Conf.GetString("basepath.timeseries"))
		assetId := st.Params[0]
		if assetId == "" {
			mware.ResponseWriter(w, assetErrMsg, http.StatusBadRequest)
			return
		}
		res, err := st.Kv.Get(assetId)
		if err != nil {
			mware.ResponseWriter(w, processErrMsg, http.StatusInternalServerError)
			st.Logger.Error(err.Error())
		} else {
			mware.ResponseWriter(w, res, http.StatusOK)
		}
	})
}
