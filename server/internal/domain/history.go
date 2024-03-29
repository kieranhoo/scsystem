package domain

type History struct {
	Id             int    `json:"id" gorm:"column:id"`
	RegistrationId string `json:"registration_id" gorm:"column:registration_id"`
	ActivityType   string `json:"activity_type" gorm:"column:activity_type"`
	Time           string `json:"time" gorm:"column:time"`
	AdminId        string `json:"admin_id" gorm:"column:admin_id"`
}

type HistoryData struct {
	Id               int    `json:"id"`
	ActivityType     string `json:"activity_type"`
	Time             string `json:"time"`
	AdminId          string `json:"admin_id"`
	RegistrationTime string `json:"registration_time"`
	Supervisor       string `json:"supervisor"`
	UserId           string `json:"user_id"`
	StartDay         string `json:"start_day"`
	EndDay           string `json:"end_day"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	RoomName         string `json:"room_name"`
	OfficeName       string `json:"office_name"`
	OrgName          string `json:"org_name"`
}

type HistoryDataResponse struct {
	TotalIn  int           `json:"total_in"`
	TotalOut int           `json:"total_out"`
	RoomName string        `json:"room_name"`
	Data     []HistoryData `json:"data"`
}

type IHistory interface {
	Insert(_his *History) error
	Latest(registrationId string) (*History, error)
	Empty() bool
	GetList(limit string) ([]History, error)
	GetActivityType() string
	GetHistory(date, roomId string) ([]HistoryData, error)
}
