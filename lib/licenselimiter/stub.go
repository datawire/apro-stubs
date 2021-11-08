package licenselimiter

import (
	"github.com/datawire/apro/lib/licensekeys"
)

type mockGauge struct{}

func (_ mockGauge) IncrementUsage() error {
	return nil
}

type Gauge interface {
	IncrementUsage() error
}

type RedisLimiter interface {
	CreateGauge(licensekeys.Limit) Gauge
}

type mockLimiter struct{}

func (_ mockLimiter) CreateGauge(_ licensekeys.Limit) Gauge {
	return mockGauge{}
}

func NewMockLimiter(_ map[licensekeys.Limit]int, _ bool) RedisLimiter {
	return mockLimiter{}
}
