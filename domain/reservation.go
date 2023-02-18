package domain

import "time"

type Reservation struct {
	Id   int
	From time.Time
	To   time.Time
}

type ReservationService interface {
	GetReservationByRoomId(roomId int) (*Reservation, error)
	CreateReservation(*Reservation) error
	DeleteReservation(int) error
}

type ReservationRepository interface {
}
