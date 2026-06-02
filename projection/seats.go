package projection

import (
	"encoding/json"
	"ticket-booking/core"
)

type SeatsProjection struct {
	seats map[string]map[string]string // concertID -> seatID -> status
}

type seatReservedData struct {
	SeatID string
	UserID string
}

type ConcertCreatedData struct {
	Name  string
	Seats []string
}

func NewSeatsProjection() *SeatsProjection {
	return &SeatsProjection{
		seats: make(map[string]map[string]string),
	}
}

func (p *SeatsProjection) Handle(event core.Event) error {
	switch event.EventType {
	case "ConcertCreated":
		var data ConcertCreatedData
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		p.seats[event.AggregateID] = make(map[string]string)
		for _, seatID := range data.Seats {
			p.seats[event.AggregateID][seatID] = "available"
		}
	case "SeatReserved":
		var data seatReservedData
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		if _, exists := p.seats[event.AggregateID]; !exists {
			p.seats[event.AggregateID] = make(map[string]string)
		}
		p.seats[event.AggregateID][data.SeatID] = "booked"
	}
	return nil
}

func (p *SeatsProjection) GetAvailableSeats(concertID string) []string {
	var availableSeats []string
	for seatID, status := range p.seats[concertID] {
		if status == "available" {
			availableSeats = append(availableSeats, seatID)
		}
	}
	return availableSeats
}
