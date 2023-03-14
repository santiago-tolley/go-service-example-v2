package endpoint

import (
	"api-go-example/internal/entity"
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-kit/kit/endpoint"
	validator "github.com/go-playground/validator/v10"
)

var (
	v *validator.Validate
)

func ValidatorMiddleware(nxt endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		if err := validate(in); err != nil {
			return nil, err
		}
		return nxt(ctx, in)
	}
}

func validate(req interface{}) error {
	v = validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("fieldName"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := v.Struct(req); err != nil {
		reqError := make([]string, 0)
		oneError := make([]string, 0)
		maxError := make([]string, 0)
		defaultError := make([]string, 0)

		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				reqError = append(reqError, e.Field())
			case "oneof":
				oneError = append(oneError, e.Field())
			case "min", "len", "lte", "max":
				maxError = append(maxError, e.Field())
			default:
				defaultError = append(defaultError, e.Field())
			}
		}
		if len(reqError) > 0 {
			return entity.TypedError{
				Type:    entity.INPUT_ERROR,
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("the following fields are required: %v", strings.Join(reqError, ", ")),
			}
		}
		if len(oneError) > 0 {
			return entity.TypedError{
				Type:    entity.INPUT_ERROR,
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("the following fields do not match the enumerated: %v", strings.Join(oneError, ", ")),
			}
		}
		if len(maxError) > 0 {
			return entity.TypedError{
				Type:    entity.INPUT_ERROR,
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("the following fields do not match the required length: %v", strings.Join(maxError, ", ")),
			}
		}
		if len(defaultError) > 0 {
			return entity.TypedError{
				Type:    entity.INPUT_ERROR,
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("the following fields are invalid: %v", strings.Join(defaultError, ", ")),
			}
		}
	}
	return nil
}
