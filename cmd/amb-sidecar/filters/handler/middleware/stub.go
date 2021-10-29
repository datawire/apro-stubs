package middleware

import (
	"context"
	"net/http"

	crd "github.com/datawire/apro/apis/getambassador.io/v1beta2"
	"github.com/datawire/apro/lib/filterapi"
)

func WithRequestID(ctx context.Context, _ string) context.Context {
	return ctx
}

func NewTemplatedErrorResponse(_ *crd.ErrorResponse, _ context.Context, _ int, _ error, _ map[string]interface{}) *filterapi.HTTPResponse {
	return &filterapi.HTTPResponse{
		Header: http.Header{},
		Body:   "",
	}
}
