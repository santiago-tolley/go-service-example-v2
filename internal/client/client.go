package client

import (
	"api-go-example/internal/entity"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

const (
	PATH_VAR    = "pathVar"
	QUERY_VAR_1 = "queryVar1"
	QUERY_VAR_2 = "queryVar2"
	HEADER_VAR  = "headerVar"
)

type CustomClient interface {
	ExecuteGET(context.Context, *CustomGetRequest) (*CustomGetResponse, error)
	ExecutePOST(context.Context, *CustomPostRequest) (*CustomPostResponse, error)
}

type customClient struct {
	epGet  endpoint.Endpoint
	epPost endpoint.Endpoint
}

func NewCustomClient(l log.Logger, serviceURI string, t time.Duration) CustomClient {
	return &customClient{
		epGet:  makeCustomClientGetEp(l, serviceURI, t),
		epPost: makeCustomClientPostEp(l, serviceURI, t),
	}
}

///////////////////////////
//// GET functionality ////
///////////////////////////
func makeCustomClientGetEp(logger log.Logger, uri string, t time.Duration) endpoint.Endpoint {
	u, _ := url.Parse(uri)
	queryList := []string{}
	for i, v := range u.Query() {
		queryList = append(queryList, i, v[0])
	}
	route := mux.NewRouter().
		Methods("GET").
		Schemes(u.Scheme).
		Host(u.Host).
		Path(u.Path).
		Queries(queryList...)

	return httptransport.NewClient(
		"GET",
		u,
		makeEncodeGetRequest(logger, route),
		makeDecodeGetResponse(logger),
		httptransport.SetClient(&http.Client{Timeout: t}),
	).Endpoint()
}

func makeEncodeGetRequest(l log.Logger, route *mux.Route) func(_ context.Context, r *http.Request, request interface{}) error {
	return func(_ context.Context, r *http.Request, request interface{}) (err error) {
		req := request.(*CustomGetRequest)
		if r.URL, err = route.URL(
			PATH_VAR, req.PathVar,
			QUERY_VAR_1, req.QueryVar1,
			QUERY_VAR_2, req.QueryVar2,
		); err != nil {
			return entity.TypedError{
				Type:    entity.CLIENT_ERROR,
				Code:    http.StatusBadRequest,
				Message: "",
			} //TODO errors!
		}
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add(HEADER_VAR, req.HeaderVar)
		return nil
	}
}

func makeDecodeGetResponse(l log.Logger) func(_ context.Context, r *http.Response) (interface{}, error) {
	return func(_ context.Context, r *http.Response) (interface{}, error) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		if r.StatusCode != http.StatusOK {
			responseBody := buf.String()
			// TODO use the logger
			return nil, entity.TypedError{
				Type:    entity.CLIENT_ERROR,
				Code:    r.StatusCode,
				Message: responseBody,
			}
		}

		response := new(CustomGetResponse)
		if err := json.Unmarshal(buf.Bytes(), response); err != nil {
			return nil, err
		}

		return response, nil
	}
}

func (c customClient) ExecuteGET(ctx context.Context, in *CustomGetRequest) (*CustomGetResponse, error) {
	res, err := c.epGet(ctx, in)
	if err != nil {
		return nil, err
	}
	response := res.(*CustomGetResponse)
	return response, nil
}

///////////////////////////
//// POST functionality ///
///////////////////////////
func makeCustomClientPostEp(logger log.Logger, uri string, t time.Duration) endpoint.Endpoint {
	u, _ := url.Parse(uri)

	return httptransport.NewClient(
		"GET",
		u,
		makeEncodePostRequest(logger),
		makeDecodePostResponse(logger),
		httptransport.SetClient(&http.Client{Timeout: t}),
	).Endpoint()
}

func makeEncodePostRequest(logger log.Logger) func(_ context.Context, r *http.Request, request interface{}) error {
	return func(_ context.Context, r *http.Request, request interface{}) error {
		r.Header.Add("Content-Type", "application/json")

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(request); err != nil {
			// TODO use logger
			return err
		}
		r.Body = ioutil.NopCloser(&buf)
		return nil
	}
}

func makeDecodePostResponse(l log.Logger) func(_ context.Context, r *http.Response) (interface{}, error) {
	return func(_ context.Context, r *http.Response) (interface{}, error) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)

		if r.StatusCode != http.StatusOK {
			responseBody := buf.String()
			// TODO use the logger

			return nil, entity.TypedError{
				Type:    entity.CLIENT_ERROR,
				Code:    r.StatusCode,
				Message: responseBody,
			}
		}

		response := new(CustomPostResponse)
		if err := json.Unmarshal(buf.Bytes(), response); err != nil {
			return nil, err
		}

		return response, nil
	}
}

func (c customClient) ExecutePOST(ctx context.Context, in *CustomPostRequest) (*CustomPostResponse, error) {
	res, err := c.epPost(ctx, in)
	if err != nil {
		return nil, err
	}
	response := res.(*CustomPostResponse)
	return response, nil
}
