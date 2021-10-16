package v1

type HeaderFieldTemplate struct {
	Name  string
	Value string
}

func (hf *HeaderFieldTemplate) Execute(_ interface{}) (*string, error) {
	return nil, nil
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
	return ""
}

type RateLimitUnit struct {
	int
}

var RateLimitUnit_MINUTE = RateLimitUnit{0}
