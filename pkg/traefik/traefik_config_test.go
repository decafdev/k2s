package traefik

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestTraefikConfig(t *testing.T) {
	// given struct
	type given struct {
		input *TraefikOptions
		err   bool
	}

	scenario := make(map[string]given)

	scenario["when all options are set"] = given{
		input: &TraefikOptions{
			ForwardAuthorizers: map[string]*ForwardAuthorizer{
				"forward-auth": NewForwardAuthorizerMiddleware(&ForwardAuth{
					Address:             "https://someplace.com",
					TrustForwardHeader:  true,
					AuthRequestHeaders:  []string{""},
					AuthResponseHeaders: []string{""},
				}),
			},
			PathPrefixStrippers: map[string]*StripPathPrefixRegex{
				"rewrite-url": NewStripPathPrefixRegexMiddleware([]string{""}),
			},
			RateLimiters: map[string]*RateLimiter{
				"rate-limit-100": NewRateLimiterMiddleware(&RateLimit{
					SourceCriterion: SourceCriterion{
						RequestHeaderName: "authorization",
					},
					Average: 100,
					Burst:   100,
				}),
			},
		},
	}

	scenario["when no traefik options are provided"] = given{
		input: &TraefikOptions{},
	}

	for when, given := range scenario {
		t.Run(when, func(t *testing.T) {
			resource, err := NewTraefikConfig(given.input)

			if (err != nil) != given.err {
				t.Errorf("unexpected error with [%v]", err)
				return
			}

			b, err := resource.ToJSON()
			if err != nil {
				t.Error(err)
			}
			snaps.MatchSnapshot(t, string(b))

		})
	}
}
