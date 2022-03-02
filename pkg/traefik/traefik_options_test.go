package traefik_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/techdecaf/k2s/v2/pkg/traefik"
)

func TestTraefikOptions(t *testing.T) {
	// given struct
	type given struct {
		input  *traefik.ResourceOptions
		expect traefik.ResourceOptions
		err    string
	}

	// traefik default options
	defaults := traefik.ResourceOptions{
		Replicas:          1,
		HostHTTPPort:      32080,
		HostHTTPSPort:     32443,
		HostDashboardPort: 32088,
	}

	scenario := make(map[string]given)

	scenario["when I pass all required params"] = given{
		input: &traefik.ResourceOptions{
			Name:         "Test",
			Namespace:    "Testing",
			Version:      "99.99.99",
			Replicas:     3,
			HostHTTPPort: 32099,
		},
		expect: traefik.ResourceOptions{
			Name:              "test",
			Namespace:         "testing",
			Version:           "99.99.99",
			Replicas:          3,
			HostHTTPPort:      32099,
			HostHTTPSPort:     32443,
			HostDashboardPort: 32088,
		},
	}

	scenario["when I specify a host port smaller than 32000"] = given{
		input: &traefik.ResourceOptions{
			Name:         "Test",
			Namespace:    "Testing",
			Version:      "99.99.99",
			Replicas:     1,
			HostHTTPPort: 80,
		},
		expect: traefik.ResourceOptions{
			Name:              "test",
			Namespace:         "testing",
			Version:           "99.99.99",
			Replicas:          1,
			HostHTTPPort:      80,
			HostHTTPSPort:     32443,
			HostDashboardPort: 32088,
		},
		err: "Key: 'ResourceOptions.HostHTTPPort' Error:Field validation for 'HostHTTPPort' failed on the 'min' tag",
	}

	scenario["when I validate an empty deployment spec"] = given{
		input:  &traefik.ResourceOptions{},
		expect: defaults,
		err: strings.Join([]string{
			"Key: 'ResourceOptions.Name' Error:Field validation for 'Name' failed on the 'required' tag",
			"Key: 'ResourceOptions.Namespace' Error:Field validation for 'Namespace' failed on the 'required' tag",
			"Key: 'ResourceOptions.Version' Error:Field validation for 'Version' failed on the 'semver' tag",
		}, "\n"),
	}

	for when, given := range scenario {
		t.Run(when, func(t *testing.T) {
			resource, err := given.input.Validate()

			if (err != nil) && (given.err != err.Error()) {
				t.Errorf("unexpected error with [%s]", err.Error())
				return
			}

			if !reflect.DeepEqual(*resource, given.expect) {
				t.Errorf("%s got %v, but expected %v", when, *resource, given.expect)
			}
		})
	}
}
