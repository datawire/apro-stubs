package middleware

import (
	"context"

	crd "github.com/datawire/apro/apis/getambassador.io/v1beta2"
	"github.com/datawire/apro/lib/filterapi"
)

func WithRequestID(_ context.Context, _ string) context.Context {
	return nil
}

func NewTemplatedErrorResponse(_ *crd.ErrorResponse, _ context.Context, _ int, _ error, _ map[string]interface{}) *filterapi.HTTPResponse {
	return nil
}
