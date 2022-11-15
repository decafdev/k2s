package util

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestNewConfig(t *testing.T) {
	// given struct
	type given struct {
		params []string
	}

	tests := make(map[string]given)

	tests["when no values are passed"] = given{
		params: []string{""},
	}

	tests["when defined parameters are provided"] = given{
		params: []string{
			"VERSION=99.99.99",
			"LOGGER_LEVEL=DEBUG",
			"LOG_PRETTY_PRINT=TRUE",
			"ENVIRONMENT=PROD",
			"PORT=1337",
		},
	}

	tests["when validation fails"] = given{
		params: []string{
			"VERSION=not-semver",
			"LOGGER_LEVEL=invalid",
			"LOG_PRETTY_PRINT=invalid",
			"PORT=should-be-numeric",
		},
	}

	for when, given := range tests {
		t.Run(when, func(t *testing.T) {
			svc, err := NewConfig(given.params...).Validate()

			if err != nil {
				snaps.MatchSnapshot(t, svc, err.Error())
			} else {
				snaps.MatchSnapshot(t, svc, nil)
			}
		})
	}
}
