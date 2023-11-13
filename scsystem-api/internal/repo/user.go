package repo

import (
	"qrcheckin/internal/model"
	"qrcheckin/pkg/database"

	"gorm.io/gorm"
)

type Users struct {
	data *model.Users
	conn *gorm.DB
}

func NewUser() IUsers {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &Users{
		data: &model.Users{},
		conn: conn,
	}
}

func (user *Users) GetByEmail(email string) (*model.Users, error) {
	if err := user.conn.Raw("SELECT * FROM users WHERE email = ?", email).Scan(user.data).Error; err != nil {
		return nil, err
	}
	return user.data, nil
}

func (user *Users) GetByID(Id string) (*model.Users, error) {
	// if err := conn.Raw("SELECT * FROM users WHERE id = ?", Id).Scan(_user).Error; err != nil {
	// 	return nil, err
	// }
	if err := user.conn.Table("users").Where("id=?", Id).First(user.data).Error; err != nil {
		return nil, err
	}
	return user.data, nil
}

func (user *Users) Insert(_user *model.Users) error {
	if err := user.conn.Exec(
		"INSERT INTO users (id, first_name, last_name, email, phone_number, role, title, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		_user.Id, _user.FirstName, _user.LastName, _user.Email, _user.PhoneNumber, _user.Role, _user.Title, _user.Password).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) Empty() bool {
	return user.data.Email == ""
}

func (user *Users) PromoteAdmin(id, role, password, email, phoneNumber string) error {
	if err := user.conn.Exec(
		"UPDATE users SET password = ?, role=?, email = ?, phone_number = ? WHERE id = ?;",
		password, role, email, phoneNumber, id).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) GetPassword() string {
	return user.data.Password
}

func (user *Users) GetEmail() string {
	return user.data.Email
}
