package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/pkg/errors"

	crd "github.com/datawire/apro/v2/apis/getambassador.io/v3alpha1"
	"github.com/datawire/apro/v2/lib/filterapi"
)

type requestIDContextKey struct{}

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDContextKey{}, requestID)
}

var fallbackTmpl = &crd.ErrorResponse{
	Headers: []crd.HeaderFieldTemplate{
		{Name: "Content-Type", Value: "application/json"},
	},
	RawBodyTemplate: `{{ . | json "" }}`,
}

type errorData map[string]interface{}

func (ed errorData) MarshalJSON() ([]byte, error) {
	copy := make(map[string]interface{})
	for k, v := range ed {
		switch k {
		case "request_id":
			httpStatus, httpStatusOK := ed["status_code"].(int)
			if httpStatusOK && httpStatus/100 == 5 {
				copy[k] = v
			}
		default:
			copy[k] = v
		}
	}
	return json.Marshal(copy)
}

func NewTemplatedErrorResponse(
	tmpl *crd.ErrorResponse,
	ctx context.Context,
	httpStatus int,
	err error,
	extra map[string]interface{},
) *filterapi.HTTPResponse {
	bodyData := errorData{
		"status_code": httpStatus,
		"message":     err.Error(),
		"request_id":  ctx.Value(requestIDContextKey{}).(string),
	}
	for k, v := range extra {
		if _, set := bodyData[k]; !set {
			bodyData[k] = v
		}
	}
	header := make(http.Header)
	for _, field := range tmpl.Headers {
		value, err := field.Execute(bodyData)
		if err != nil {
			return NewTemplatedErrorResponse(fallbackTmpl, ctx, 500,
				errors.Wrapf(err, "errorResponse: generating header %q", field.Name), nil)
		}
		if value != nil {
			header.Set(field.Name, *value)
		}
	}

	bodyTmpl, err := template.
		New("bodyTemplate").
		Funcs(template.FuncMap{
			"json": func(prefix string, data interface{}) (string, error) {
				bs, err := json.MarshalIndent(data, prefix, "\t")
				return string(bs), err
			},
		}).
		Parse(tmpl.RawBodyTemplate)
	if err != nil {
		panic(fmt.Errorf("the stubs do not support this: error parsing bodyTemplate: %w", err))
	}
	body := new(strings.Builder)
	if err := bodyTmpl.Execute(body, bodyData); err != nil {
		return NewTemplatedErrorResponse(fallbackTmpl, ctx, 500,
			errors.Wrap(err, "errorResponse: generating body"), nil)
	}

	return &filterapi.HTTPResponse{
		Header: header,
		Body:   body.String(),
	}
}
