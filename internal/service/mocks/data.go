package mocks

import (
	"api-go-example/internal/client"
	cmock "api-go-example/internal/client/mocks"
	"api-go-example/internal/entity"
	"context"
	"net/http"

	"github.com/stretchr/testify/mock"
)

var (
	ClientOK *cmock.CustomClient

	ClientErr *cmock.CustomClient
)

func init() {
	ClientOK = new(cmock.CustomClient)
	ClientOK.On("ExecuteGET", mock.Anything, mock.Anything).Return(
		func(context.Context, *client.CustomGetRequest) *client.CustomGetResponse {
			return &client.CustomGetResponse{}
		},
		func(context.Context, *client.CustomGetRequest) error {
			return nil
		},
		// &client.CustomGetResponse{}, nil
	)
	ClientOK.On("ExecutePOST", mock.Anything, mock.Anything).Return(
		func(context.Context, *client.CustomPostRequest) *client.CustomPostResponse {
			return &client.CustomPostResponse{}
		},
		func(context.Context, *client.CustomPostRequest) error {
			return nil
		},
		// &client.CustomPostResponse{}, nil
	)

	ClientErr = new(cmock.CustomClient)
	ClientErr.On("ExecuteGET", mock.Anything, mock.Anything).Return(
		func(context.Context, *client.CustomGetRequest) *client.CustomGetResponse {
			return nil
		},
		func(context.Context, *client.CustomGetRequest) error {
			return &entity.TypedError{
				Type:    entity.CLIENT_ERROR,
				Code:    http.StatusBadGateway,
				Message: "bad gateway",
			}
		},
		// nil, &entity.TypedError{
		// 	Type:    entity.CLIENT_ERROR,
		// 	Code:    http.StatusBadGateway,
		// 	Message: "bad gateway",
		// }
	)
	ClientErr.On("ExecutePOST", mock.Anything, mock.Anything).Return(
		func(context.Context, *client.CustomPostRequest) *client.CustomPostResponse {
			return nil
		},
		func(context.Context, *client.CustomPostRequest) error {
			return &entity.TypedError{
				Type:    entity.CLIENT_ERROR,
				Code:    http.StatusBadGateway,
				Message: "bad gateway",
			}
		},
		// nil, &entity.TypedError{
		// 	Type:    entity.CLIENT_ERROR,
		// 	Code:    http.StatusBadGateway,
		// 	Message: "bad gateway",
		// }
	)
}
