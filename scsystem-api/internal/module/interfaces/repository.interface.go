package interfaces

import "qrcheckin/internal/module/entity"

type IUsers interface {
	GetByEmail(email string) (*entity.Users, error)
	GetByID(Id string) (*entity.Users, error)
	Insert(_user *entity.Users) error
	Empty() bool
	PromoteAdmin(id, role, password, email, phoneNumber string) error
	GetPassword() string
	GetEmail() string
}

type IRegistration interface {
	Insert(_res *entity.Registration) error
	GetByID(id string) (*entity.Registration, error)
	GetByUserIdAndRoom(userId, roomId string) (*entity.Registration, error)
	UpdateByIDAndRoom(_res *entity.Registration) error
	Latest() (*entity.Registration, error)
	Empty() bool
}

type IHistory interface {
	Insert(_his *entity.History) error
	Latest(registrationId string) (*entity.History, error)
	Empty() bool
	GetList(limit string) ([]entity.History, error)
	GetActivityType() string
}
