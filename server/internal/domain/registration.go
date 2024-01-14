package domain

type Registration struct {
	Id               string `json:"id" gorm:"column:id"`
	RegistrationTime string `json:"registration_time" gorm:"column:registration_time"`
	Supervisor       string `json:"supervisor" gorm:"column:supervisor"`
	UserID           string `json:"user_id" gorm:"column:user_id"`
	RoomId           string `json:"room_id" gorm:"column:room_id"`
	StartDay         string `json:"start_day" gorm:"column:start_day"`
	EndDay           string `json:"end_day" gorm:"column:end_day"`
}

type IRegistration interface {
	Insert(_res *Registration) error
	GetByID(id string) (*Registration, error)
	GetByUserIdAndRoom(userId, roomId string) (*Registration, error)
	UpdateByIDAndRoom(_res *Registration) error
	Latest() (*Registration, error)
	Empty() bool
	RegistrationLatest(studentId, roomId string) (*UserRoomData, error)
	OneWeekBefore() ([]Registration, error)
	DeleteById(id string) error
}
