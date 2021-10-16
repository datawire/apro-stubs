package filterapi

import (
	"net/http"
)

type HTTPResponse struct {
	Header http.Header
	Body   string
}
