package client_test

import (
	"api-go-example/internal/client"
	"api-go-example/internal/client/mocks"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientGet(t *testing.T) {

	t.Run("GET: must respond correctly", func(tt *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mocks.ClientHttpGETResponseOK))
		}))

		defer server.Close()

		cli := client.NewCustomClient(log.NewNopLogger(), server.URL, time.Second*5)
		res, err := cli.ExecuteGET(context.Background(), mocks.ClientGetRequestOK)
		assert.Nil(t, err)
		require.NotNil(t, res)
		assert.Equal(t, "sunny day", res.StringValue)
		assert.Equal(t, int32(10), res.IntValue)
	})

	t.Run("GET: must throw an error if client has an error", func(tt *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusBadGateway)
		}))

		defer server.Close()

		cli := client.NewCustomClient(log.NewNopLogger(), server.URL, time.Second*5)
		_, err := cli.ExecuteGET(context.Background(), mocks.ClientGetRequestOK)
		require.NotNil(t, err)
		assert.Equal(t, "error! 502 Bad Gateway (502): ", err.Error())
	})

	t.Run("GET: must throw an error if response is not json", func(tt *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mocks.ClientHttpGETResponseError))
		}))

		defer server.Close()

		cli := client.NewCustomClient(log.NewNopLogger(), server.URL, time.Second*5)
		_, err := cli.ExecuteGET(context.Background(), mocks.ClientGetRequestOK)
		require.NotNil(t, err)
		assert.Equal(t, "invalid character '<' looking for beginning of value", err.Error())
	})
}
func TestClientPost(t *testing.T) {

	t.Run("POST: must respond correctly", func(tt *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mocks.ClientHttpPOSTResponseOK))
		}))

		defer server.Close()

		cli := client.NewCustomClient(log.NewNopLogger(), server.URL, time.Second*5)
		res, err := cli.ExecutePOST(context.Background(), mocks.ClientPostRequestOK)
		assert.Nil(t, err)
		require.NotNil(t, res)
		assert.Equal(t, "no errors found", res.Notes)
		require.Equal(t, 2, len(res.Results))
		assert.Equal(t, int32(22), res.Results[0].DataPoint1)
		assert.Equal(t, int32(33), res.Results[0].DataPoint2)
		assert.Equal(t, int32(44), res.Results[0].DataPoint3)
		assert.Equal(t, int32(52), res.Results[1].DataPoint1)
		assert.Equal(t, int32(63), res.Results[1].DataPoint2)
		assert.Equal(t, int32(74), res.Results[1].DataPoint3)
	})

	t.Run("POST: must throw an error if client has an error", func(tt *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusBadGateway)
		}))

		defer server.Close()

		cli := client.NewCustomClient(log.NewNopLogger(), server.URL, time.Second*5)
		_, err := cli.ExecutePOST(context.Background(), mocks.ClientPostRequestOK)
		require.NotNil(t, err)
		assert.Equal(t, "error! 502 Bad Gateway (502): ", err.Error())
	})

	t.Run("POST: must throw an error if response is not json", func(tt *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mocks.ClientHttpPOSTResponseError))
		}))

		defer server.Close()

		cli := client.NewCustomClient(log.NewNopLogger(), server.URL, time.Second*5)
		_, err := cli.ExecutePOST(context.Background(), mocks.ClientPostRequestOK)
		require.NotNil(t, err)
		assert.Equal(t, "invalid character 'e' looking for beginning of object key string", err.Error())
	})
}
