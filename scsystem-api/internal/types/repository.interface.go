package types

type IUsers interface {
	GetByEmail(email string) (*Users, error)
	GetByID(Id string) (*Users, error)
	Insert(_user *Users) error
	Empty() bool
	PromoteAdmin(id, role, password, email, phoneNumber string) error
	GetPassword() string
	GetEmail() string
}

type IRegistration interface {
	Insert(_res *Registration) error
	GetByID(id string) (*Registration, error)
	GetByUserIdAndRoom(userId, roomId string) (*Registration, error)
	UpdateByIDAndRoom(_res *Registration) error
	Latest() (*Registration, error)
	Empty() bool
}

type IHistory interface {
	Insert(_his *History) error
	Latest(registrationId string) (*History, error)
	Empty() bool
	GetList(limit string) ([]History, error)
	GetActivityType() string
}
