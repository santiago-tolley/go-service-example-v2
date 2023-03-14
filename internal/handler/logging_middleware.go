package handler

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func MakeEndpointMiddleware(logger log.Logger, service, path, method string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		logger = log.WithPrefix(logger,
			"path", path,
			"method", method,
			"service", service,
			"caller", log.DefaultCaller)

		return func(ctx context.Context, req interface{}) (response interface{}, err error) {
			now := time.Now()
			level.Info(logger).Log("started", now)
			level.Debug(logger).Log("request", req)

			defer func(begin time.Time) {
				level.Info(logger).Log("took", time.Since(begin))
				if err != nil {
					level.Info(logger).Log("Response", "NOK")
					level.Error(logger).Log("error", err.Error())

				} else {
					level.Info(logger).Log("Response", "OK")
					level.Debug(logger).Log("response", response)
				}
			}(now)

			return next(ctx, req)
		}
	}
}
