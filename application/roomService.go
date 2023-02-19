package application

import (
	"errors"

	"github.com/forceattack012/reservationroom/domain"
)

type RoomService struct {
	repo domain.RoomRepository
}

func NewRoomService(repo domain.RoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

func (r *RoomService) GetAllRoom() ([]domain.Room, error) {
	var rooms []domain.Room
	var err error
	if rooms, err = r.repo.GetAll(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *RoomService) CreateRoom(room *domain.Room) error {
	return r.repo.Save(room)
}

func (r *RoomService) UpdateRoom(id int, room *domain.Room) error {
	findRoom, err := r.repo.GetRoomById(id)
	if err != nil {
		return err
	}

	if findRoom.Id == 0 {
		return errors.New("id not found")
	}

	findRoom = &domain.Room{
		RoomName:    room.RoomName,
		Kind:        room.Kind,
		Price:       room.Price,
		IsAvaliable: room.IsAvaliable,
	}
	err = r.repo.Update(id, findRoom)

	return err
}
