package timeseries

import (
	"context"

	md "microsvc/pkg/model"

	"github.com/go-kit/kit/endpoint"
)

type putTimeseriesRequest struct {
	tsa     *md.TimeseriesArray
	assetId string
}

type getTimeseriesRequest struct {
	assetId string
}

func makePutTimeseriesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(putTimeseriesRequest)
		resp := s.PutTimeseries(req.assetId, req.tsa)
		return resp, nil
	}
}

func makeGetTimeseriesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getTimeseriesRequest)
		resp := s.GetTimeseries(req.assetId)
		return resp, nil
	}
}
