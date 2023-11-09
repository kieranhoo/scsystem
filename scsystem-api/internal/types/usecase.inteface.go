package types

import (
	"qrcheckin/internal/schema"
)

type IAuthService interface {
	SignUp(req schema.SignUpRequest) error
	SignIn(req schema.SignInRequest) (string, error)
}

type ILabService interface {
	RegisterLab(req *schema.RegistrationLabRequest) error
	RegistrationLatest(studentId, roomId string) (*schema.UserLabData, error)
	SaveActivityType(req *schema.CheckInRequest) error
	GetHistories(limit string) ([]History, error)
	GetHistoriesData(limit string) ([]schema.HistoryData, error)
}