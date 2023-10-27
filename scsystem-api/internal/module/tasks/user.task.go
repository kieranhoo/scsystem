package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"qrcheckin/internal/module/model"
	"qrcheckin/internal/module/types"

	"github.com/hibiken/asynq"
)

func SignUp(id, firstName, lastName, phoneNumber, email, password string) error {
	_user, err := model.NewUser().GetByID(id)
	if err != nil || _user.Email == "" {
		return model.NewUser().Insert(&types.Users{
			Id:          id,
			FirstName:   firstName,
			LastName:    lastName,
			PhoneNumber: phoneNumber,
			Email:       email,
			Password:    password,
		})
	}
	if _user.Password == "" {
		return model.NewUser().PromoteAdmin(id, "admin", password, email, phoneNumber)
	}
	return errors.New("user already exists")
}

func SaveUser(id, firstName, lastName, phoneNumber, email string) error {
	return new(model.Users).Insert(&types.Users{
		Id:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	})
}

func HandleSaveUser(_ context.Context, task *asynq.Task) error {
	var user types.Users
	if err := json.Unmarshal(task.Payload(), &user); err != nil {
		return err
	}
	return model.NewUser().Insert(&user)
}
