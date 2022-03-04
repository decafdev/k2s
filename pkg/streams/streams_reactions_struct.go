package streams

import (
	"github.com/nats-io/nats.go"
	"github.com/reactivex/rxgo/v2"
)

// Rx function description
func Rx() *rx {
	return &rx{}
}

// rx struct
type rx struct{}

// KVEventTypeFilter method
func (t *rx) KVOperationFilter(operation nats.KeyValueOp) rxgo.Predicate {
	return func(item interface{}) bool {
		return item.(nats.KeyValueEntry).Operation() == operation
	}
}
