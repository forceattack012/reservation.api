package infrastructure

import (
	"fmt"
	"time"

	"github.com/forceattack012/reservationroom/domain"
	"github.com/forceattack012/reservationroom/enum"
)

type ReservationRepository struct {
	db *GormDB
}

func NewReservationRepository(db *GormDB) domain.ReservationRepository {
	return &ReservationRepository{db: db}
}

// CreateReservation implements domain.ReservationRepository
func (r *ReservationRepository) CreateReservation(reservation *domain.Reservation) error {
	return r.db.Save(reservation).Error
}

// DeleteReservation implements domain.ReservationRepository
func (r *ReservationRepository) DeleteReservation(id int) error {
	return r.db.Delete(domain.Reservation{}, id).Error
}

// GetReservationByRoomId implements domain.ReservationRepository
func (r *ReservationRepository) GetReservationByRoomId(roomId int) ([]domain.Reservation, error) {
	from := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), enum.MIN_START_TIME, 0, 0, 0, time.UTC)
	to := from.Add(time.Duration(time.Hour * time.Duration(enum.MAX_END_TIME)))

	var reservervations []domain.Reservation
	err := r.db.Where("start_date >= ? AND end_date <= ? AND is_reservation = ? AND room_id = ?", from, to, true, roomId).Find(&reservervations).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v", reservervations)
	return reservervations, nil
}

// GetReservationByRoomIdAndTime implements domain.ReservationRepository
func (r *ReservationRepository) GetReservationByRoomIdAndTime(id int, from time.Time, to time.Time) ([]domain.Reservation, error) {
	var reservation []domain.Reservation
	err := r.db.Where(`room_id = ? AND is_reservation = ? AND start_date >= ? AND end_date <= ?`, id, true, from, to).Find(&reservation).Error
	if err != nil {
		return nil, err
	}
	return reservation, nil
}
