package endpoint

import (
	"api-go-example/internal/entity"
	"api-go-example/internal/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GetEndpoint  endpoint.Endpoint
	PostEndpoint endpoint.Endpoint
}

func NewServiceEndpoint(svc service.Service) Set {
	return Set{
		GetEndpoint:  ValidatorMiddleware(makeGetEndpoint(svc)),
		PostEndpoint: ValidatorMiddleware(makePostEndpoint(svc)),
	}
}

func makeGetEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(entity.GetRequest)
		return svc.ExecuteGET(ctx, req)
	}
}

func makePostEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(entity.PostRequest)
		return svc.ExecutePOST(ctx, req)
	}
}
