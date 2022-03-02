package config

import (
	"context"
	"reflect"
	"strings"

	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

// ConfigService struct
type ConfigService struct {
	SERVICE_NAME             string `json:"SERVICE_NAME" mod:"lcase"`
	VERSION                  string `json:"VERSION" validate:"semver" mod:"default=0.0.0"`
	BASE_PATH                string `json:"BASE_PATH" mod:"default=/"`
	PORT                     string `json:"PORT" validate:"numeric" mod:"default=3000"`
	ENVIRONMENT              string `json:"ENVIRONMENT" mod:"default=local,lcase"`
	LOGGER_PRETTY_PRINT      string `json:"LOGGER_PRETTY_PRINT" validate:"oneof=true false" mod:"default=false,lcase"`
	LOGGER_LEVEL             string `json:"LOGGER_LEVEL" validate:"oneof=debug info warn error fatal" mod:"default=info,lcase"`
	TRAEFIK_VERSION          string `json:"TRAEFIK_VERSION" validate:"semver" mod:"default=2.5.4"`
	TRAEFIK_REPLICAS         string `json:"TRAEFIK_REPLICAS" validate:"numeric" mod:"default=1"`
	PRIVATE_REGISTRY_ENABLED string `json:"PRIVATE_REGISTRY_ENABLED"  validate:"oneof=true false" mod:"default=false,lcase"`
	PRIVATE_REGISTRY_URL     string `json:"PRIVATE_REGISTRY_URL" validate:"url" mod:"default=https://index.docker.io/v1/"`
	PRIVATE_REGISTRY_USER    string `json:"PRIVATE_REGISTRY_USER" validate:"required_if=PRIVATE_REGISTRY_ENABLED true"`
	PRIVATE_REGISTRY_PASS    string `json:"PRIVATE_REGISTRY_PASS" validate:"required_if=PRIVATE_REGISTRY_ENABLED true"`
}

func (t *ConfigService) Validate() (*ConfigService, error) {
	if err := modifiers.New().Struct(context.Background(), t); err != nil {
		return t, err
	}
	return t, validator.New().Struct(t)
}

// NewConfigService function description
func NewConfigService(params ...string) *ConfigService {
	env := env2map(params...)
	config := &ConfigService{}

	ref := reflect.ValueOf(config).Elem()
	for i := 0; i < ref.NumField(); i++ {
		key := ref.Type().Field(i).Name
		ref.Field(i).SetString(env[key])
	}

	return config
}

// env2map key=val pairs to map
func env2map(env ...string) map[string]string {
	out := make(map[string]string)
	for _, e := range env {
		if i := strings.Index(e, "="); i >= 0 {
			out[e[:i]] = e[i+1:]
		}
	}
	return out
}
