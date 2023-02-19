package dto

type ReserversionDto struct {
	RoomId         uint
	ReservationDay []ReservationDay
}

type ReservationDay struct {
	Day          string
	Reservations []interface{}
}
