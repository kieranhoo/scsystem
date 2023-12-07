package schema

type RegistrationRoomRequest struct {
	RegistrationTime string `json:"registration_time" validate:"required"`
	Supervisor       string `json:"supervisor" validate:"required"`
	RoomId           string `json:"room_id" validate:"required"`
	StartDay         string `json:"start_day" validate:"required"`
	EndDay           string `json:"end_day" validate:"required"`
	FirstName        string `json:"first_name" validate:"required"`
	LastName         string `json:"last_name" validate:"required"`
	Email            string `json:"email" validate:"required"`
	PhoneNumber      string `json:"phone_number" validate:"required"`
	StudentId        string `json:"student_id" validate:"required"`
}

type DataResponse struct {
	Success bool        `json:"success" validate:"required"`
	Data    interface{} `json:"data" validate:"required"`
}

type UserRoomData struct {
	Id               string `json:"id" validate:"required"`
	RegistrationTime string `json:"registration_time" validate:"required"`
	Supervisor       string `json:"supervisor" validate:"required"`
	StudentId        string `json:"student_id" validate:"required"`
	StartDay         string `json:"start_day" validate:"required"`
	EndDay           string `json:"end_day" validate:"required"`
	FirstName        string `json:"first_name" validate:"required"`
	LastName         string `json:"last_name" validate:"required"`
	Email            string `json:"email" validate:"required"`
	PhoneNumber      string `json:"phone_number" validate:"required"`
	RoomName         string `json:"room_name" validate:"required"`
	OfficeName       string `json:"office_name" validate:"required"`
	OrgName          string `json:"org_name" validate:"required"`
	ActivityType     string `json:"activity_type" validate:"required"`
}

type OfficeRequest struct {
	Id             string `json:"id"`
	OrganizationId string `json:"organization_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Address        string `json:"address" validate:"required"`
	Manager        string `json:"manager" validate:"required"`
	PhoneNumber    string `json:"phone_number" validate:"required"`
}

type CheckInRequest struct {
	RegistrationId string `json:"registration_id" validate:"required"`
	AdminId        string `json:"admin_id" validate:"required"`
	ActivityType   string `json:"activity_type" validate:"required"`
}

type Response struct {
	Success bool   `json:"success" validate:"required"`
	Msg     string `json:"msg" validate:"required"`
}
