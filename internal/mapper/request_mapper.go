package mapper

import (
	"api-go-example/internal/client"
	"api-go-example/internal/entity"
	"strconv"
)

type GetRequestMapper func(entity.GetRequest) (*client.CustomGetRequest, error)

func MapGetRequest(in entity.GetRequest) (*client.CustomGetRequest, error) {
	return &client.CustomGetRequest{
		PathVar:   in.ID,
		HeaderVar: in.Info,
		QueryVar1: in.Data.Value,
		QueryVar2: strconv.FormatInt(int64(in.Data.IValue), 10),
	}, nil
}

type PostRequestMapper func(entity.PostRequest) (*client.CustomPostRequest, error)

func MapPostRequest(in entity.PostRequest) (*client.CustomPostRequest, error) {
	return &client.CustomPostRequest{
		StringValue: in.Str,
		IntValue:    in.Integer,
	}, nil
}
