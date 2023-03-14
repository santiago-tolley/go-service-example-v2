package handler

import (
	"api-go-example/internal/entity"
	"context"
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func ErrorToHttpCode(err error) (int, ErrorMessage) {
	switch e := err.(type) {
	case entity.TypedError:
		switch e.Type {
		case entity.INPUT_ERROR:
			return http.StatusBadRequest, ErrorMessage{
				Message: e.Error(),
			}
		case entity.CLIENT_ERROR, entity.SERVICE_ERROR:
			return http.StatusInternalServerError, ErrorMessage{
				Message: e.Error(),
			}
		default:
			return e.Code, ErrorMessage{
				Message: e.Error(),
			}
		}
	default:
		return http.StatusInternalServerError, ErrorMessage{
			Message: err.Error(),
		}
	}
}

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	code, msg := ErrorToHttpCode(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(msg)
}
