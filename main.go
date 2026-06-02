package main

import (
	"encoding/json"
	"fmt"
	"ticket-booking/cmd"
	"ticket-booking/core"
	"ticket-booking/projection"
	"ticket-booking/store"
)

func main() {
	s := store.NewInMemoryStore()
	proj := projection.NewSeatsProjection()
	s.Subscribe(proj.Handle)

	data, _ := json.Marshal(projection.ConcertCreatedData{
		Name:  "Rock Concert",
		Seats: []string{"A1", "A2", "A3"},
	})
	s.Append(core.Event{
		AggregateID: "concert1",
		EventType:   "ConcertCreated",
		EventID:     "evt-1",
		Version:     1,
		Data:        data,
	})

	cmd.BookSeat(s, "concert1", "A1", "user123")

	fmt.Println("available:", proj.GetAvailableSeats("concert1"))
}
