package infrastructure

import (
	"github.com/forceattack012/reservationroom/domain"
)

type RoomRepository struct {
	db *GormDB
}

func NewRoomRepository(gormdb *GormDB) domain.RoomRepository {
	return &RoomRepository{db: gormdb}
}

func (r *RoomRepository) GetAll() ([]domain.Room, error) {
	var rooms []domain.Room
	if err := r.db.Where(&domain.Room{IsAvaliable: true}).Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomRepository) Save(room *domain.Room) error {
	return r.db.Save(room).Error
}

func (r *RoomRepository) Update(id int, room *domain.Room) error {
	return r.db.Model(&domain.Room{}).Where("id = ?", id).Updates(domain.Room{
		RoomName:    room.RoomName,
		Kind:        room.Kind,
		Price:       room.Price,
		IsAvaliable: false,
	}).Error
}

func (r *RoomRepository) GetRoomById(id int) (*domain.Room, error) {
	var room *domain.Room
	if err := r.db.Find(&room, id).Error; err != nil {
		return nil, err
	}
	return room, nil
}
