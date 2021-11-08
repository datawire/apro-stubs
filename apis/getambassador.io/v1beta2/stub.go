package v1

import (
	"fmt"
	"strings"
	"text/template"
)

type HeaderFieldTemplate struct {
	Name  string
	Value string
}

func (hf *HeaderFieldTemplate) Execute(data interface{}) (*string, error) {
	tmpl, err := template.
		New(hf.Name).
		Parse(hf.Value)
	if err != nil {
		return nil, err
	}
	w := new(strings.Builder)
	if err := tmpl.Execute(w, data); err != nil {
		return nil, err
	}
	value := w.String()
	return &value, nil
}

type ErrorResponse struct {
	Headers []HeaderFieldTemplate

	RawBodyTemplate string
}

type RateLimitSpec struct {
	AmbassadorID []string
	Domain       string
	Limits       []Limit
}

func (s *RateLimitSpec) Validate(_ string) error {
	return nil
}

type Limit struct {
	Unit          RateLimitUnit
	ErrorResponse ErrorResponse
}

type RateLimitAction struct {
	int
}

var RateLimitAction_LOG_ONLY = RateLimitAction{1}

func (a *RateLimitAction) ToString() string {
	if a == nil {
		return "Enforce"
	}
	switch a.int {
	case 0:
		return "Enforce"
	case 1:
		return "LogOnly"
	default:
		panic(fmt.Errorf("should not happen: invalid RateLimitAction{%d}", *a))
	}
}

type RateLimitUnit struct {
	int
}

var RateLimitUnit_MINUTE = RateLimitUnit{0}
