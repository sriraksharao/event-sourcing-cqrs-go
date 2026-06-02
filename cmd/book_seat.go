package cmd

import (
	"ticket-booking/core"
	"ticket-booking/domain"
)

func BookSeat(store core.EventStore, concertID, seatID, userID string) error {
	// step 1: load all past events for this concert
	events, err := store.Load(concertID)
	if err != nil {
		return err
	}

	// step 2: replay events to rebuild concert state
	concert := domain.NewConcert(concertID, "")
	for _, e := range events {
		err := concert.Apply(e)
		if err != nil {
			return err
		}
	}

	// step 3: try to book the seat
	event, err := concert.BookSeat(seatID, userID)
	if err != nil {
		return err
	}

	// step 4: save the new event
	return store.Append(event)
}
