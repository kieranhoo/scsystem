package model

type Users struct {
	Id          string `json:"id" gorm:"column:id"`
	FirstName   string `json:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" gorm:"column:last_name"`
	Email       string `json:"email" gorm:"column:email"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Role        string `json:"role" gorm:"column:role"`
	Title       string `json:"title" gorm:"column:title"`
	Password    string `json:"password" gorm:"column:password"`
}

type Registration struct {
	Id               string `json:"id" gorm:"column:id"`
	RegistrationTime string `json:"registration_time" gorm:"column:registration_time"`
	Supervisor       string `json:"supervisor" gorm:"column:supervisor"`
	UserID           string `json:"user_id" gorm:"column:user_id"`
	RoomId           string `json:"room_id" gorm:"column:room_id"`
	StartDay         string `json:"start_day" gorm:"column:start_day"`
	EndDay           string `json:"end_day" gorm:"column:end_day"`
}

type History struct {
	Id             int    `json:"id" gorm:"column:id"`
	RegistrationId string `json:"registration_id" gorm:"column:registration_id"`
	ActivityType   string `json:"activity_type" gorm:"column:activity_type"`
	Time           string `json:"time" gorm:"column:time"`
	AdminId        string `json:"admin_id" gorm:"column:admin_id"`
}
type Office struct {
	Id             string `json:"id" gorm:"column:id"`
	OrganizationId string `json:"organization_id" gorm:"column:organization_id"`
	Name           string `json:"name" gorm:"column:name"`
	Address        string `json:"address" gorm:"column:address"`
	Manager        string `json:"manager" gorm:"column:manager"`
	PhoneNumber    string `json:"phone_number" gorm:"column:phone_number"`
}

type Organization struct {
	Id          string `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Email       string `json:"email" gorm:"column:email"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Website     string `json:"website" gorm:"column:website"`
	Head        string `json:"head" gorm:"column:head"`
}

type Room struct {
	Id          string `json:"id" gorm:"column:id"`
	OfficeID    string `json:"office_id" gorm:"column:office_id"`
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

type Chart struct {
	Id       string `json:"id" gorm:"column:id"`
	Time     string `json:"time" gorm:"column:time"`
	RoomId   string `json:"room_id" gorm:"column:room_id"`
	InCount  int    `json:"in_count" gorm:"column:in_count"`
	OutCount int    `json:"out_count" gorm:"column:out_count"`
}
