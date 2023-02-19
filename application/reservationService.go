package application

import (
	"errors"

	"github.com/forceattack012/reservationroom/domain"
	"github.com/forceattack012/reservationroom/dto"
)

type ReservationService struct {
	repo domain.ReservationRepository
}

func NewReservationService(repo domain.ReservationRepository) domain.ReservationService {
	return &ReservationService{
		repo: repo,
	}
}

// CreateReservation implements domain.ReservationService
func (r *ReservationService) CreateReservation(reservation *domain.Reservation) error {
	reservertions, err := r.repo.GetReservationByRoomIdAndTime(int(reservation.RoomID), reservation.StartDate, reservation.EndDate)
	if err != nil {
		return err
	}
	if len(reservertions) > 0 {
		return errors.New("reservation already")
	}

	return r.repo.CreateReservation(reservation)
}

// DeleteReservation implements domain.ReservationService
func (r *ReservationService) DeleteReservation(id int) error {
	return r.repo.DeleteReservation(id)
}

// GetReservationByRoomId implements domain.ReservationService
func (r *ReservationService) GetReservationByRoomId(roomId int) (dto.ReserversionDto, error) {
	reservations, err := r.repo.GetReservationByRoomId(roomId)
	if err != nil {
		return dto.ReserversionDto{}, err
	}

	groupByDay := make(map[string][]interface{})
	for _, r := range reservations {
		key := r.StartDate.Weekday().String()
		groupByDay[key] = append(groupByDay[key], r)
	}

	response := dto.ReserversionDto{
		RoomId:         uint(roomId),
		ReservationDay: []dto.ReservationDay{},
	}
	ReservationDay := []dto.ReservationDay{}

	for k, group := range groupByDay {
		ReservationDay = append(ReservationDay, dto.ReservationDay{
			Day:          k,
			Reservations: group,
		})
	}

	response.ReservationDay = append(response.ReservationDay, ReservationDay...)

	return response, nil
}
