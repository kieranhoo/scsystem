package repo

import (
	"qrcheckin/internal/model"
)

type IUsers interface {
	GetByEmail(email string) (*model.Users, error)
	GetByID(Id string) (*model.Users, error)
	Insert(_user *model.Users) error
	Empty() bool
	PromoteAdmin(id, role, password, email, phoneNumber string) error
	GetPassword() string
	GetEmail() string
}

type IRegistration interface {
	Insert(_res *model.Registration) error
	GetByID(id string) (*model.Registration, error)
	GetByUserIdAndRoom(userId, roomId string) (*model.Registration, error)
	UpdateByIDAndRoom(_res *model.Registration) error
	Latest() (*model.Registration, error)
	Empty() bool
}

type IHistory interface {
	Insert(_his *model.History) error
	Latest(registrationId string) (*model.History, error)
	Empty() bool
	GetList(limit string) ([]model.History, error)
	GetActivityType() string
}
