package service

import (
	"scsystem/internal/model"
	"scsystem/internal/schema"
)

type IAuthService interface {
	SignUp(req schema.SignUpRequest) error
	SignIn(req schema.SignInRequest) (string, error)
	GetMe(id string) (*schema.UserResponse, error)
}

type IRoomService interface {
	RegisterRoom(req *schema.RegistrationRoomRequest) error
	RegistrationLatest(studentId, roomId string) (*schema.UserRoomData, error)
	SaveActivityType(req *schema.CheckInRequest) error
	GetHistories(limit string) ([]model.History, error)
	GetHistoriesData(limit string) ([]schema.HistoryData, error)
	GetRoom() ([]schema.RoomData, error)
}

type IStat interface {
	GetChartData(roomId string) (*schema.ChartMetadata, error)
	GetRoomData(roomId string) (*schema.RoomStat, error)
}
