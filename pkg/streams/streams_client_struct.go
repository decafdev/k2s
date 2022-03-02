package streams

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"github.com/reactivex/rxgo/v2"
)

// subject  orders_service.v1.noun.verb
// stream   orders_service_noun_stream
// consumer consumer_service_noun_queue

// NewClient constructor
func NewClient(ServiceName string, nc *nats.Conn) (*Client, error) {
	js, err := nc.JetStream()
	if err != nil {
		return &Client{}, err
	}

	client := &Client{
		ServiceName: strings.ToLower(ServiceName),
		NATS:        nc,
		Streams:     js,
	}

	go client.OnOSExit()
	// nats.ErrorHandler(NatsErrHandler)

	return client, nil
}

// Client struct
type Client struct {
	ServiceName string
	NATS        *nats.Conn
	Streams     nats.JetStreamContext
}

// OnOSExit method
func (t *Client) OnOSExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	sig := <-c
	fmt.Println(fmt.Sprintf("streams_client: os %v signal received draining nats connection", sig))
	if err := t.NATS.Drain(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("streams_client: goodbye")
}

// emit method
func (t *Client) emit(sub *nats.Subscription, ch chan rxgo.Item) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		messages, err := sub.Fetch(1, nats.Context(ctx))
		if err != nil {
			ch <- rxgo.Error(err)
		}

		for _, msg := range messages {
			ch <- rxgo.Of(msg)
		}
		cancel()
	}
}

// OnRequestEvent method
func (t *Client) OnRequestEvent(subject string, payload []byte, timeout time.Duration) rxgo.Observable {
	return rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		msg, err := t.NATS.Request(subject, payload, timeout)
		if err != nil {
			ch <- rxgo.Error(err)
		}

		if err := t.NATS.Flush(); err != nil {
			ch <- rxgo.Error(err)
			close(ch)
		}
		ch <- rxgo.Of(msg)

		close(ch)
	}})
}

// SendResponse method
func (t *Client) SendResponse(subject string) rxgo.Observable {
	ch := make(chan rxgo.Item)
	stream := rxgo.FromChannel(ch)

	_, err := t.NATS.QueueSubscribe(subject, t.ServiceName, func(msg *nats.Msg) {
		ch <- rxgo.Of(msg)
	})

	if err != nil {
		ch <- rxgo.Error(err)
	}

	return stream
}

// OnTableEvent method
func (t *Client) OnTableEvent(table nats.KeyValue) rxgo.Observable {
	ch := make(chan rxgo.Item)
	watcher, err := table.WatchAll()
	if err != nil {
		ch <- rxgo.Error(err)
	}

	go func() {
		for update := range watcher.Updates() {
			if update != nil {
				ch <- rxgo.Of(update)
			}
		}
	}()

	return rxgo.FromChannel(ch)
}

// OnStreamEvent method
func (t *Client) OnStreamEvent(subject string) rxgo.Observable {
	ch := make(chan rxgo.Item)
	observable := rxgo.FromChannel(ch)

	sub, err := t.Subscription(subject)
	if err != nil {
		ch <- rxgo.Error(err)
	}

	// emit to channel
	go t.emit(sub, ch)

	// Create an Observable
	return observable
}

// CreateStream method
func (t *Client) CreateStream(subjects ...string) (*nats.StreamInfo, error) {
	name := NewSubject(subjects[0])

	for _, sub := range subjects {
		if NewSubject(sub).Stream != name.Stream {
			return &nats.StreamInfo{}, errors.New("invalid stream, you can not create a stream with mixed schemas")
		}
	}
	// subject  orders_service.v1.noun.verb
	// subject  orders_service.v1.noun.verb
	// subject := fmt.Sprintf("%s.%s.%s.*", t.ServiceName, version, noun)

	return t.Streams.AddStream(&nats.StreamConfig{
		Name: name.Stream,
		// Description:       fmt.Sprintf("persistence layer for %s", name.Stream),
		Subjects:          subjects,
		Retention:         0,
		MaxConsumers:      0,
		MaxMsgs:           0,
		MaxBytes:          0,
		Discard:           0,
		MaxAge:            0,
		MaxMsgsPerSubject: 0,
		MaxMsgSize:        0,
		Storage:           0,
		Replicas:          0,
		NoAck:             false,
		Duplicates:        0,
		Sealed:            false,
		DenyDelete:        false,
		DenyPurge:         false,
		AllowRollup:       false,
		// Template:          "",
		// Placement:         &nats.Placement{},
		// Mirror:            &nats.StreamSource{},
		// Sources:           []*nats.StreamSource{},
	})
}

// CreateConsumer - method
func (t *Client) CreateConsumer(subjects ...string) (*nats.ConsumerInfo, error) {
	stream, err := t.CreateStream(subjects...)
	if err != nil {
		return &nats.ConsumerInfo{}, err
	}

	// consumer consumer_service_noun_queue
	options := []nats.JSOpt{}
	fmt.Println(stream.Config.Name)

	return t.Streams.AddConsumer(stream.Config.Name, &nats.ConsumerConfig{
		// DeliverGroup:    "",
		// OptStartSeq:     0,
		// OptStartTime:    &time.Time{},
		MaxDeliver: 10,
		// Durable:         fmt.Sprintf("%s_q", subject),
		// DeliverSubject:  fmt.Sprintf("%s_q", subject),
		// FilterSubject:   subject,
		Description:     "nats jetstream durable pull consumer",
		DeliverPolicy:   nats.DeliverAllPolicy,
		AckPolicy:       nats.AckExplicitPolicy,
		AckWait:         time.Second * 30,
		ReplayPolicy:    nats.ReplayInstantPolicy,
		RateLimit:       0,
		SampleFrequency: "",
		MaxWaiting:      0,
		MaxAckPending:   20000,
		FlowControl:     false,
		Heartbeat:       time.Second * 5,
		HeadersOnly:     false,
	}, options...)
}

// Subscription method
func (t *Client) Subscription(subject string) (*nats.Subscription, error) {
	return t.Streams.PullSubscribe(subject, t.ServiceName, nats.PullMaxWaiting(128))
}

// Table method
func (t *Client) Table(name string) (nats.KeyValue, error) {
	bucket := strings.ToLower(fmt.Sprintf("%s_%s_table", t.ServiceName, name))
	return t.Streams.CreateKeyValue(&nats.KeyValueConfig{
		Bucket:  bucket,
		History: 3,
		Storage: nats.MemoryStorage,
		TTL:     time.Millisecond * 100,
		// Description: "",
		//MaxValueSize: 0,
		//MaxBytes:     0,
		//Replicas:     0,
	})

}
