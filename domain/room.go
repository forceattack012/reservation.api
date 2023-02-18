package domain

type Room struct {
	Id          int
	RoomName    string  `json:"room_name"`
	Kind        string  `json:"kind"`
	Price       float64 `json:"price"`
	IsAvaliable bool    `json:"is_available"`
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
	Update(room *Room) error
}
