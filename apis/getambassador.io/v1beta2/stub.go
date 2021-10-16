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

type RateLimitUnit struct {
	int
}

var RateLimitUnit_MINUTE = RateLimitUnit{0}
