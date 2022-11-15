package traefik

type SourceCriterion struct {
	RequestHeaderName string `yaml:"requestHeaderName,omitempty"`
}

type RateLimit struct {
	SourceCriterion SourceCriterion `yaml:"sourceCriterion,omitempty"`
	Average         int             `yaml:"average,omitempty"`
	Burst           int             `yaml:"burst,omitempty"`
}

// RateLimiter struct
type RateLimiter struct {
	RateLimit *RateLimit `yaml:"rateLimit,omitempty"`
}

// NewRateLimiterMiddleware function description
func NewRateLimiterMiddleware(o *RateLimit) *RateLimiter {
	return &RateLimiter{RateLimit: o}
}
