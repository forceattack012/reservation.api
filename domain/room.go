package domain

type Room struct {
	Id          uint          `json:"id" gorm:"primaryKey"`
	RoomName    string        `json:"room_name"`
	Kind        string        `json:"kind"`
	Price       float64       `json:"price"`
	IsAvaliable bool          `json:"is_available"`
	Reservation []Reservation `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type RoomService interface {
	GetAllRoom() ([]Room, error)
	CreateRoom(*Room) error
	UpdateRoom(id int, room *Room) error
}

type RoomRepository interface {
	GetAll() ([]Room, error)
	GetRoomById(id int) (*Room, error)
	Save(room *Room) error
	Update(id int, room *Room) error
}
