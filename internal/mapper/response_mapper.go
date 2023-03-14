package mapper

import (
	"api-go-example/internal/client"
	"api-go-example/internal/entity"
	"net/http"
)

type GetResponseMapper func(*client.CustomGetResponse) (*entity.GetResponse, error)

func MapGetResponse(in *client.CustomGetResponse) (*entity.GetResponse, error) {
	if in == nil {
		return nil, entity.TypedError{
			Type:    entity.CLIENT_ERROR,
			Code:    http.StatusBadRequest,
			Message: "client response is empty",
		}
	}
	return &entity.GetResponse{
		Str:     in.StringValue,
		Integer: in.IntValue,
	}, nil
}

type PostResponseMapper func(*client.CustomPostResponse) (*entity.PostResponse, error)

func MapPostResponse(in *client.CustomPostResponse) (*entity.PostResponse, error) {
	if in == nil {
		return nil, entity.TypedError{
			Type:    entity.CLIENT_ERROR,
			Code:    http.StatusBadRequest,
			Message: "client response is empty",
		}
	}
	dp := map[int][]int32{}
	for i, v := range in.Results {
		dp[i+1] = append(dp[i+1], v.DataPoint1)
		dp[i+1] = append(dp[i+1], v.DataPoint2)
		dp[i+1] = append(dp[i+1], v.DataPoint3)
	}
	return &entity.PostResponse{
		DataPoints: dp,
	}, nil
}
