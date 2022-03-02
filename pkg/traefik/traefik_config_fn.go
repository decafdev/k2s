package traefik

import (
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"
)

// TraefikOptions struct
type TraefikOptions struct {
	ForwardAuthorizers  map[string]*ForwardAuthorizer
	PathPrefixStrippers map[string]*StripPathPrefixRegex
	RateLimiters        map[string]*RateLimiter
}

// HTTP struct
type HTTP struct {
	Middlewares map[string]interface{} `yaml:"middlewares,omitempty"`
}

// TraefikConfig struct
type TraefikConfig struct {
	HTTP *HTTP `yaml:"http,omitempty"`
}

// ToYAML method
func (t *TraefikConfig) ToYAML() ([]byte, error) {
	return yaml.Marshal(t)
}

// ToJSON method
func (t *TraefikConfig) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

// FromJSON method
func (t *TraefikConfig) FromJSON(data []byte) error {
	return json.Unmarshal(data, t)
}

// NewTraefikConfig function description
func NewTraefikConfig(o *TraefikOptions) (*TraefikConfig, error) {
	t := &TraefikConfig{
		HTTP: &HTTP{
			Middlewares: make(map[string]interface{}),
		},
	}

	if len(o.PathPrefixStrippers) != 0 {
		for k, v := range o.PathPrefixStrippers {
			if err := t.NameConflict(k); err != nil {
				return t, err
			}
			t.HTTP.Middlewares[k] = v
		}
	}

	if len(o.ForwardAuthorizers) != 0 {
		for k, v := range o.ForwardAuthorizers {
			if err := t.NameConflict(k); err != nil {
				return t, err
			}
			t.HTTP.Middlewares[k] = v
		}
	}

	if len(o.RateLimiters) != 0 {
		for k, v := range o.RateLimiters {
			if err := t.NameConflict(k); err != nil {
				return t, err
			}
			t.HTTP.Middlewares[k] = v
		}
	}

	return t, nil
}

// NameConflict method
func (t *TraefikConfig) NameConflict(name string) (err error) {
	if t.HTTP.Middlewares[name] != nil {
		return errors.New(fmt.Sprintf("name conflict, a middleware with the name %s already exists", name))
	}
	return nil
}
