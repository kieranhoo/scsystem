package service

import (
	"errors"
	"scsystem/internal/domain"
	"scsystem/internal/repo"
	"scsystem/pkg/jwt"
)

type Auth struct {
	repo domain.IUsers
}

func NewAuth() domain.IAuthService {
	return &Auth{
		repo: repo.NewUser(),
	}
}

func (auth *Auth) SignUp(req domain.SignUpRequest) error {
	_pass, err := jwt.GenPassword(req.Password)
	if err != nil {
		return err
	}
	_user, err := auth.repo.GetByID(req.Id)
	if err != nil || _user == nil {
		return auth.repo.Insert(&domain.Users{
			Id:          req.Id,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
			Password:    string(_pass),
		})
	}
	if _user.Password == "" {
		return repo.NewUser().PromoteAdmin(req.Id, "admin", string(_pass), req.Email, req.PhoneNumber)
	}
	return errors.New("user already exists")
}

func (auth *Auth) SignIn(req domain.SignInRequest) (string, error) {
	_user, err := auth.repo.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := jwt.ComparePassword(_user.Password, req.Password); err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(_user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (auth *Auth) GetMe(id string) (*domain.UserResponse, error) {
	users, err := auth.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &domain.UserResponse{
		Id:          users.Id,
		FirstName:   users.FirstName,
		LastName:    users.LastName,
		Email:       users.Email,
		PhoneNumber: users.PhoneNumber,
	}, nil
}
