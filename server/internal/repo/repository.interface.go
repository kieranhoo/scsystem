package repo

import (
	"scsystem/internal/model"
	"scsystem/internal/schema"
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
	RegistrationLatest(studentId, roomId string) (*schema.UserRoomData, error)
}

type IHistory interface {
	Insert(_his *model.History) error
	Latest(registrationId string) (*model.History, error)
	Empty() bool
	GetList(limit string) ([]model.History, error)
	GetActivityType() string
	GetHistory(limit string) ([]schema.HistoryData, error)
}

type IOffice interface {
	Insert(office *model.Office) error
	Get() ([]schema.OfficeData, error)
}

type IRoom interface {
	Insert(room *model.Room) error
	Get() ([]schema.RoomData, error)
}

type IOrganization interface {
	Insert(org *model.Organization) error
	Get() ([]model.Organization, error)
}
