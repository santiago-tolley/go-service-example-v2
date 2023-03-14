package handler

import (
	"api-go-example/internal/endpoint"
	"api-go-example/internal/entity"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	ep "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func NewHttpHandler(l log.Logger, endpoints endpoint.Set) http.Handler {

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(ErrorEncoder),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(l)),
	}

	m := mux.NewRouter()

	makeGETEndpoint(l, endpoints.GetEndpoint, m, options)
	makePOSTEndpoint(l, endpoints.PostEndpoint, m, options)

	return nil
}

func makeGETEndpoint(l log.Logger, e ep.Endpoint, m *mux.Router, o []httptransport.ServerOption) {
	handler := httptransport.NewServer(
		MakeEndpointMiddleware(l, "get_service", "/", "GET")(e),
		decodeGetRequest,
		encodeResponse,
		o...,
	)

	m.Methods("GET").Path("/").Handler(handler).Queries()
	m.Methods("GET").Path("/").Handler(handler)
}

func makePOSTEndpoint(l log.Logger, e ep.Endpoint, m *mux.Router, o []httptransport.ServerOption) {
	handler := httptransport.NewServer(
		MakeEndpointMiddleware(l, "post_service", "/", "POST")(e),
		decodePostRequest,
		encodeResponse,
		o...,
	)

	m.Methods("POST").Path("/").Handler(handler)
}

func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	intiv, err := strconv.ParseInt(vars["value2"], 10, 32)
	if err != nil {
		return nil, entity.TypedError{
			Type:    entity.INPUT_ERROR,
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("ivalue must be an integer %v", err),
		}
	}

	return &entity.GetRequest{
		ID:   vars["id"],
		Info: r.Header.Get(""),
		Data: entity.GetData{
			Value:  vars["value1"],
			IValue: int32(intiv),
		},
	}, nil
}

func decodePostRequest(ctx context.Context, r *http.Request) (interface{}, error) {

	req := &entity.PostRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if f, ok := response.(ep.Failer); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}
