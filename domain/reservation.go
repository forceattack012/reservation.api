package domain

import (
	"time"

	"github.com/forceattack012/reservationroom/dto"
)

type Reservation struct {
	Id            uint      `json:"id" gorm:"primaryKey"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	IsReservation bool      `json:"is_reservation"`
	PersonID      uint      `json:"person_id"`
	RoomID        uint      `json:"room_id"`
}

type ReservationService interface {
	GetReservationByRoomId(roomId int) (dto.ReserversionDto, error)
	CreateReservation(*Reservation) error
	DeleteReservation(int) error
}

type ReservationRepository interface {
	GetReservationByRoomId(roomId int) ([]Reservation, error)
	CreateReservation(*Reservation) error
	DeleteReservation(int) error
	GetReservationByRoomIdAndTime(id int, from time.Time, to time.Time) ([]Reservation, error)
}
