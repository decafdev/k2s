package traefik_test

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/techdecaf/k2s/v2/pkg/traefik"
)

func TestTraefikConfig(t *testing.T) {
	// given struct
	type given struct {
		input *traefik.TraefikOptions
		err   bool
	}

	scenario := make(map[string]given)

	scenario["when all options are set"] = given{
		input: &traefik.TraefikOptions{
			ForwardAuthorizers: map[string]*traefik.ForwardAuthorizer{
				"forward-auth": traefik.NewForwardAuthorizerMiddleware(&traefik.ForwardAuth{
					Address:             "https://someplace.com",
					TrustForwardHeader:  true,
					AuthRequestHeaders:  []string{""},
					AuthResponseHeaders: []string{""},
				}),
			},
			PathPrefixStrippers: map[string]*traefik.StripPathPrefixRegex{
				"rewrite-url": traefik.NewStripPathPrefixRegexMiddleware([]string{""}),
			},
			RateLimiters: map[string]*traefik.RateLimiter{
				"rate-limit-100": traefik.NewRateLimiterMiddleware(&traefik.RateLimit{
					SourceCriterion: traefik.SourceCriterion{
						RequestHeaderName: "authorization",
					},
					Average: 100,
					Burst:   100,
				}),
			},
		},
	}

	scenario["when no traefik options are provided"] = given{
		input: &traefik.TraefikOptions{},
	}

	for when, given := range scenario {
		t.Run(when, func(t *testing.T) {
			resource, err := traefik.NewTraefikConfig(given.input)

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
