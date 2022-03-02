package kube

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ContainerRegistry struct
type ContainerRegistry struct {
	Secret *coreV1.Secret
}

// ToYAML method
func (t *ContainerRegistry) ToYAML() ([]byte, error) {
	return tx.ResourcesToYAML([]runtime.Object{t.Secret})
}

// ContainerRegistryOptions struct
type ContainerRegistryOptions struct {
	Name      string `json:"-" validate:"required"`
	Namespace string `json:"-" validate:"required"`
	Registry  string `json:"-" validate:"required,url"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Auth      string `json:"auth" validate:"required,base64"`
	Email     string `json:"email" mod:"default=unused"`
}

func (t *ContainerRegistryOptions) Validate() (*ContainerRegistryOptions, error) {
	if err := modifiers.New().Struct(context.Background(), t); err != nil {
		return t, err
	}
	// set auth by base64 encoding username and password
	t.Auth = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", t.Username, t.Password)))
	return t, validator.New().Struct(t)
}

// NewContainerRegistry function description
func NewContainerRegistry(o *ContainerRegistryOptions) (reg *ContainerRegistry, err error) {
	if o, err = o.Validate(); err != nil {
		return &ContainerRegistry{}, err
	}

	credentials := map[string]map[string]*ContainerRegistryOptions{"auths": {o.Registry: o}}
	config, err := json.Marshal(credentials)

	return &ContainerRegistry{
		Secret: &coreV1.Secret{
			TypeMeta: metaV1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metaV1.ObjectMeta{
				Name:      o.Name,
				Namespace: o.Namespace,
			},
			StringData: map[string]string{
				".dockerconfigjson": string(config),
			},
		},
	}, err
}
