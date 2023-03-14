package service

import (
	"api-go-example/internal/client"
	"api-go-example/internal/entity"
	"api-go-example/internal/mapper"
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type Service interface {
	ExecuteGET(context.Context, entity.GetRequest) (*entity.GetResponse, error)
	ExecutePOST(context.Context, entity.PostRequest) (*entity.PostResponse, error)
}

type service struct {
	logger        log.Logger
	getReqMapper  mapper.GetRequestMapper
	getResMapper  mapper.GetResponseMapper
	postReqMapper mapper.PostRequestMapper
	postResMapper mapper.PostResponseMapper
	customClient  client.CustomClient
}

func NewService(l log.Logger,
	gmReq mapper.GetRequestMapper,
	gmRes mapper.GetResponseMapper,
	pmReq mapper.PostRequestMapper,
	pmRes mapper.PostResponseMapper,
	c client.CustomClient,
) Service {
	return &service{
		logger:        l,
		getReqMapper:  gmReq,
		getResMapper:  gmRes,
		postReqMapper: pmReq,
		postResMapper: pmRes,
		customClient:  c,
	}
}

func (s *service) ExecuteGET(ctx context.Context, in entity.GetRequest) (*entity.GetResponse, error) {
	level.Debug(s.logger).Log("api_request", in)
	getReq, err := s.getReqMapper(in)
	if err != nil {
		level.Error(s.logger).Log("error generating GET request", err.Error())
		return nil, err
	}

	level.Debug(s.logger).Log("get_request", getReq)
	clientRes, err := s.customClient.ExecuteGET(ctx, getReq)
	if err != nil {
		level.Error(s.logger).Log("error calling GET client", err.Error())
		return nil, err
	}

	level.Debug(s.logger).Log("get_response", clientRes)
	res, err := s.getResMapper(clientRes)
	if err != nil {
		level.Error(s.logger).Log("error generating GET response", err.Error())
		return nil, err
	}

	level.Debug(s.logger).Log("api_response", res)
	return res, nil
}
func (s *service) ExecutePOST(ctx context.Context, in entity.PostRequest) (*entity.PostResponse, error) {
	level.Debug(s.logger).Log("api_request", in)
	postReq, err := s.postReqMapper(in)
	if err != nil {
		level.Error(s.logger).Log("error generating POST request", err.Error())
		return nil, err
	}

	level.Debug(s.logger).Log("post_request", postReq)
	clientRes, err := s.customClient.ExecutePOST(ctx, postReq)
	if err != nil {
		level.Error(s.logger).Log("error calling POST client", err.Error())
		return nil, err
	}

	level.Debug(s.logger).Log("post_response", clientRes)
	res, err := s.postResMapper(clientRes)
	if err != nil {
		level.Error(s.logger).Log("error generating POST response", err.Error())
		return nil, err
	}

	level.Debug(s.logger).Log("api_response", res)
	return res, nil
}
