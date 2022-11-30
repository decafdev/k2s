package kube

import (
	"encoding/base64"
	"fmt"
	"strings"
)

const (
	CreatedBySelector = "k2s.techdecaf.io/created-by=k2s-operator"
	CreatedBy         = "k2s-operator"
	IdKey             = "k2s.techdecaf.io/id"
	NameKey           = "k2s.techdecaf.io/name"
	VersionKey        = "k2s.techdecaf.io/version"
	CreateByKey       = "k2s.techdecaf.io/created-by"
)

// Labels struct
type Labels struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewLabels() *Labels {
	return &Labels{}
}

// ResourceLabels method
func (t *Labels) ResourceLabels() map[string]string {
	return t.ToMap([]string{"k2s.techdecaf.io/name", "k2s.techdecaf.io/created-by"})
}

// FromID method
func (t *Labels) FromID(id string) *Labels {
	data, _ := base64.StdEncoding.DecodeString(id)
	parts := strings.Split(string(data), "|")

	t.Id = id
	t.Name = parts[0]
	t.Version = parts[1]

	return t
}

// identity method
func (t *Labels) identity(name, version string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s|%s", name, version)))
}

// FromKeys method
func (t *Labels) From(name, version string) *Labels {
	t.Id = t.identity(name, version)
	t.Name = name
	t.Version = version
	return t
}

// FromMap method
func (t *Labels) FromMap(labels map[string]string) *Labels {
	t.Id = labels[IdKey]
	t.Name = labels[NameKey]
	t.Version = labels[VersionKey]
	return t
}

// ToMap method
func (t *Labels) ToMap(filter []string) map[string]string {
	labels := map[string]string{
		IdKey:       t.identity(t.Name, t.Version),
		NameKey:     t.Name,
		VersionKey:  t.Version,
		CreateByKey: CreatedBy,
	}

	if len(filter) == 0 {
		return labels
	}

	filtered := map[string]string{}
	for _, key := range filter {
		filtered[key] = labels[key]
	}
	return filtered
}

func (t *Labels) Selector(key string) string {
	return fmt.Sprintf("%s=%s", key, t.ToMap([]string{})[key])
}
