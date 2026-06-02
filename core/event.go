package core

type Event struct {
	AggregateID string
	EventType   string
	// Data        interface{}
	Data      []byte
	Timestamp int64
	EventID   string
	Version   int
}

type EventStore interface {
	Append(event Event) error
	Load(aggregateID string) ([]Event, error)
}

type EventHandler func(event Event) error
