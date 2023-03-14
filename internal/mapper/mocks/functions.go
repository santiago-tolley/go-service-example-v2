package mocks

import (
	"api-go-example/internal/client"
	"api-go-example/internal/entity"
	"net/http"
)

// Requests

func MapGetRqOK(entity.GetRequest) (*client.CustomGetRequest, error) {
	return &client.CustomGetRequest{}, nil
}

func MapGetRqError(entity.GetRequest) (*client.CustomGetRequest, error) {
	return nil, entity.TypedError{
		Type:    entity.SERVICE_ERROR,
		Code:    http.StatusInternalServerError,
		Message: "request mapper error",
	}
}

func MapPostRqOK(entity.PostRequest) (*client.CustomPostRequest, error) {
	return &client.CustomPostRequest{}, nil
}

func MapPostRqError(entity.PostRequest) (*client.CustomPostRequest, error) {
	return nil, entity.TypedError{
		Type:    entity.SERVICE_ERROR,
		Code:    http.StatusInternalServerError,
		Message: "request mapper error",
	}
}

// Responses

func MapGetResOK(*client.CustomGetResponse) (*entity.GetResponse, error) {
	return &entity.GetResponse{}, nil
}

func MapGetResError(*client.CustomGetResponse) (*entity.GetResponse, error) {
	return nil, entity.TypedError{
		Type:    entity.SERVICE_ERROR,
		Code:    http.StatusInternalServerError,
		Message: "response mapper error",
	}
}

func MapPostResOK(*client.CustomPostResponse) (*entity.PostResponse, error) {
	return &entity.PostResponse{}, nil
}

func MapPostResError(*client.CustomPostResponse) (*entity.PostResponse, error) {
	return nil, entity.TypedError{
		Type:    entity.SERVICE_ERROR,
		Code:    http.StatusInternalServerError,
		Message: "response mapper error",
	}
}
