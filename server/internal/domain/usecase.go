package domain

type IAuthService interface {
	SignUp(req SignUpRequest) error
	SignIn(req SignInRequest) (string, error)
	GetMe(id string) (*UserResponse, error)
}

type IRoomService interface {
	RegisterRoom(req *RegistrationRoomRequest) error
	RegistrationLatest(studentId, roomId string) (*UserRoomData, error)
	SaveActivityType(req *CheckInRequest) error
	GetHistories(limit string) ([]History, error)
	GetHistoriesData(date, roomId string) (*HistoryDataResponse, error)
	GetRoom() ([]RoomData, error)
}

type IStatService interface {
	GetChartData(roomId string) (*ChartMetadata, error)
	GetRoomData() (*RoomStat, error)
}
