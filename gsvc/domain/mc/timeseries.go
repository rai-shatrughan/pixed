package mc

import (
	"encoding/json"
	"net/http"

	"gsvc/pkg/model"
	"gsvc/pkg/util"

	"github.com/gorilla/mux"
)

func PostTimeseries(kf *util.KafkaWriter, logger *util.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body model.TimeseriesArray
		vars := mux.Vars(r)
		assetId, ok := vars["assetId"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("AssetId Missing in Path"))
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		kfmsg, _ := json.Marshal(body)

		err := kf.Write([]byte(assetId), []byte(kfmsg))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not write TS, retry after sometime"))
			logger.Error(err.Error())
			return
		} else {
			w.Write([]byte("{\"TimeseriesUpload\": \"Ok\"}"))
			return
		}
	})
}

func GetTimeseries(kv *util.KV, logger *util.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		assetId, ok := vars["assetId"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("AssetId Missing in Path"))
			return
		}

		res, err := kv.GetFromKeyWithLimit("/"+assetId, 1000)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not Fetch TS, retry after sometime"))
			logger.Error(err.Error())
		} else {
			w.Write([]byte(res))
		}

	})
}
