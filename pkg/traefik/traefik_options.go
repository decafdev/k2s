package traefik

import (
	"context"

	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

type TraefikResourceOptions struct {
	Name              string `validate:"required" mod:"lcase"`
	Namespace         string `validate:"required" mod:"lcase"`
	Version           string `validate:"semver"`
	Replicas          int32  `mod:"default=1"`
	HostHTTPPort      int32  `validate:"min=32000" mod:"default=32080"`
	HostHTTPSPort     int32  `validate:"min=32000" mod:"default=32443"`
	HostDashboardPort int32  `validate:"min=32000" mod:"default=32088"`
}

func (t *TraefikResourceOptions) Validate() (*TraefikResourceOptions, error) {
	err := modifiers.New().Struct(context.Background(), t)
	if err != nil {
		return t, err
	}
	return t, validator.New().Struct(t)
}
