package domain

import (
	"encoding/json"
	"errors"
	"ticket-booking/core"
	"time"

	"github.com/google/uuid"
)

var ErrSeatAlreadyBooked = errors.New("seat already booked")

type Concert struct {
	ID          string
	Name        string
	bookedSeats map[string]bool
	version     int
}

func NewConcert(id, name string) *Concert {
	return &Concert{
		ID:          id,
		Name:        name,
		bookedSeats: make(map[string]bool),
		version:     0,
	}
}

type SeatReservedData struct {
	SeatID string
	UserID string
}

func (c *Concert) Apply(event core.Event) error {
	switch event.EventType {
	case "SeatReserved":
		var data SeatReservedData
		// Assume we have a function to unmarshal event data into SeatReservedData
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		c.bookedSeats[data.SeatID] = true

	}
	c.version++
	return nil

}

func (c *Concert) BookSeat(seatID, userID string) (core.Event, error) {
	if c.bookedSeats[seatID] {
		return core.Event{}, ErrSeatAlreadyBooked
	}

	eventData := SeatReservedData{
		SeatID: seatID,
		UserID: userID,
	}
	dataBytes, err := json.Marshal(eventData)
	if err != nil {
		return core.Event{}, err
	}

	event := core.Event{
		AggregateID: c.ID,
		EventType:   "SeatReserved",
		Data:        dataBytes,
		Timestamp:   time.Now().Unix(),
		EventID:     uuid.New().String(),
		Version:     c.version + 1,
	}

	return event, nil
}
