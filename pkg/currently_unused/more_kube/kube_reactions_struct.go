package morekube

import (
	"context"

	"github.com/reactivex/rxgo/v2"
	coreV1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/watch"
)

// Rx struct
type Rx struct{}

// EventType method
func (t *Rx) EventTypeFilter(eventType watch.EventType) rxgo.Predicate {
	return func(item interface{}) bool {
		event := item.(watch.Event)
		return event.Type == eventType
	}
}

// NamespaceMap function description
func (t *Rx) NamespaceMap() rxgo.Func {
	return func(_ context.Context, item interface{}) (interface{}, error) {
		return item.(watch.Event).Object.(*coreV1.Namespace), nil
	}
}
