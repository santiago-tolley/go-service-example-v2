package service_test

import (
	"api-go-example/internal/client"
	"api-go-example/internal/entity"
	"api-go-example/internal/mapper"
	"api-go-example/internal/service"
	"context"
	"testing"

	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mmock "api-go-example/internal/mapper/mocks"
	smock "api-go-example/internal/service/mocks"
)

func TestServiceGet(t *testing.T) {
	tc := []struct {
		name       string
		gReqMapper mapper.GetRequestMapper
		gResMapper mapper.GetResponseMapper
		cli        client.CustomClient
		err        error
	}{
		{
			name:       "must execute without errors",
			gReqMapper: mmock.MapGetRqOK,
			gResMapper: mmock.MapGetResOK,
			cli:        smock.ClientOK,
			err:        nil,
		},
		{
			name:       "must fail if request mapper returns an error",
			gReqMapper: mmock.MapGetRqError,
			gResMapper: mmock.MapGetResOK,
			cli:        smock.ClientOK,
			err:        entity.TypedError{},
		},
		{
			name:       "must fail if response mapper returns an error",
			gReqMapper: mmock.MapGetRqOK,
			gResMapper: mmock.MapGetResError,
			cli:        smock.ClientOK,
			err:        entity.TypedError{},
		},
		{
			name:       "must fail if client returns an error",
			gReqMapper: mmock.MapGetRqOK,
			gResMapper: mmock.MapGetResOK,
			cli:        smock.ClientErr,
			err:        entity.TypedError{},
		},
	}

	l := log.NewNopLogger()
	ctx := context.Background()

	for _, c := range tc {

		t.Run(c.name, func(tt *testing.T) {
			svc := service.NewService(l,
				c.gReqMapper,
				c.gResMapper,
				nil,
				nil,
				c.cli)
			res, err := svc.ExecuteGET(ctx, entity.GetRequest{})
			if c.err != nil {
				require.NotNil(t, err)
				assert.Equal(t, c.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
				require.NotNil(t, res)
			}
		})
	}
}

func TestServicePost(t *testing.T) {

	tc := []struct {
		name       string
		pReqMapper mapper.PostRequestMapper
		pResMapper mapper.PostResponseMapper
		cli        client.CustomClient
		err        error
	}{
		{
			name:       "must execute without errors",
			pReqMapper: mmock.MapPostRqOK,
			pResMapper: mmock.MapPostResOK,
			cli:        smock.ClientOK,
			err:        nil,
		},
		{
			name:       "must fail if request mapper returns an error",
			pReqMapper: mmock.MapPostRqError,
			pResMapper: mmock.MapPostResOK,
			cli:        smock.ClientOK,
			err:        entity.TypedError{},
		},
		{
			name:       "must fail if response mapper returns an error",
			pReqMapper: mmock.MapPostRqOK,
			pResMapper: mmock.MapPostResError,
			cli:        smock.ClientOK,
			err:        entity.TypedError{},
		},
		{
			name:       "must fail if client returns an error",
			pReqMapper: mmock.MapPostRqOK,
			pResMapper: mmock.MapPostResOK,
			cli:        smock.ClientErr,
			err:        entity.TypedError{},
		},
	}

	l := log.NewNopLogger()
	ctx := context.Background()

	for _, c := range tc {
		t.Run(c.name, func(tt *testing.T) {
			svc := service.NewService(l,
				nil,
				nil,
				c.pReqMapper,
				c.pResMapper,
				c.cli)

			res, err := svc.ExecutePOST(ctx, entity.PostRequest{})
			if c.err != nil {
				require.NotNil(t, err)
				assert.Equal(t, c.err.Error(), err.Error())
			} else {
				assert.Nil(t, err)
				require.NotNil(t, res)
			}
		})
	}
}
