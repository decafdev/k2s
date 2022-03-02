package streams

import (
	"fmt"
	"strings"
)

type Subject struct {
	Service string
	Version string
	Kind    string
	Action  string
	Stream  string
}

// ToString method
func (t *Subject) ToString() string {
	return fmt.Sprintf("%s.%s.%s.%s.", t.Service, t.Version, t.Kind, t.Action)
}

// NewSubject function description
func NewSubject(subject string) *Subject {
	s := strings.Split(subject, ".")
	kind := strings.Replace(s[2], "*", "dlq", 1)
	action := strings.Replace(s[3], "*", "dlq", 1)
	return &Subject{
		Service: s[0],
		Version: s[1],
		Kind:    kind,
		Action:  action,
		Stream:  fmt.Sprintf("%s_%s_%s_stream", s[0], s[1], kind),
	}
}
