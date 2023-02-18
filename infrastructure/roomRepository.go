package infrastructure

import "github.com/forceattack012/reservationroom/domain"

type RoomRepository struct {
	db *GormDB
}

func NewRoomRepository(gormdb *GormDB) domain.RoomRepository {
	return &RoomRepository{db: gormdb}
}

func (r *RoomRepository) GetAll() ([]domain.Room, error) {
	var rooms []domain.Room
	if err := r.db.Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomRepository) Save(room *domain.Room) error {
	return r.db.Save(room).Error
}

func (r *RoomRepository) Update(room *domain.Room) error {
	return r.db.Updates(room).Error
}

func (r *RoomRepository) GetRoomById(id int) (*domain.Room, error) {
	var room *domain.Room
	if err := r.db.Find(room, id).Error; err != nil {
		return nil, err
	}
	return room, nil
}
