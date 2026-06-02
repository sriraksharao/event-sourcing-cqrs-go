package store

import (
	"sync"
	"ticket-booking/core"
)

type InMemoryStore struct {
	events   map[string][]core.Event
	mu       sync.RWMutex
	handlers []core.EventHandler
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		events: make(map[string][]core.Event),
	}
}

func (s *InMemoryStore) Append(event core.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events[event.AggregateID] = append(s.events[event.AggregateID], event)
	for _, h := range s.handlers {
		h(event)
	}
	return nil
}

func (s *InMemoryStore) Load(aggregateID string) ([]core.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.events[aggregateID], nil
}

func (s *InMemoryStore) Subscribe(handler core.EventHandler) {
	s.handlers = append(s.handlers, handler)
	// return nil
}
