package cmd

import (
	"encoding/json"
	"ticket-booking/core"
	"ticket-booking/projection"

	"github.com/google/uuid"
)

func CreateConcert(store core.EventStore, concertID, name string, seats []string) error {
	data, err := json.Marshal(projection.ConcertCreatedData{
		Name:  name,
		Seats: seats,
	})
	if err != nil {
		return err
	}
	return store.Append(core.Event{
		AggregateID: concertID,
		EventType:   "ConcertCreated",
		EventID:     uuid.New().String(),
		Version:     1,
		Data:        data,
	})
}
