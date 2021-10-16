package licenselimiter

import (
	"github.com/datawire/apro/lib/licensekeys"
)

type Gauge interface {
	IncrementUsage() error
}

type RedisLimiter interface {
	CreateGauge(licensekeys.Limit) Gauge
}

func NewMockLimiter(_ map[licensekeys.Limit]int, _ bool) RedisLimiter {
	return nil
}
