package domain

type Office struct {
	Id             string `json:"id" gorm:"column:id"`
	OrganizationId string `json:"organization_id" gorm:"column:organization_id"`
	Name           string `json:"name" gorm:"column:name"`
	Address        string `json:"address" gorm:"column:address"`
	Manager        string `json:"manager" gorm:"column:manager"`
	PhoneNumber    string `json:"phone_number" gorm:"column:phone_number"`
}
type OfficeData struct {
	Id             string `json:"id" gorm:"column:id" validate:"required"`
	Name           string `json:"name" gorm:"column:name" validate:"required"`
	PhoneNumber    string `json:"phone_number" gorm:"column:phone_number"`
	Manager        string `json:"manager" gorm:"column:manager" validate:"required"`
	OrgID          string `json:"org_id" gorm:"column:org_id" validate:"required"`
	OrgName        string `json:"org_name" gorm:"column:org_name" validate:"required"`
	OrgPhoneNumber string `json:"org_phone_number" gorm:"column:org_phone_number" validate:"required"`
	OrgEmail       string `json:"org_email" gorm:"column:org_email" validate:"required"`
	OrgHead        string `json:"org_head" gorm:"column:org_head" validate:"required"`
	OrgWebsite     string `json:"org_website" gorm:"column:org_website" validate:"required"`
}
type NewOfficeRequest struct {
	OrganizationId string `json:"organization_id" gorm:"column:organization_id" validate:"required"`
	Name           string `json:"name" gorm:"column:name" validate:"required"`
	Address        string `json:"address" gorm:"column:address" validate:"required"`
	Manager        string `json:"manager" gorm:"column:manager" validate:"required"`
	PhoneNumber    string `json:"phone_number" gorm:"column:phone_number" validate:"required"`
}
type OfficeRequest struct {
	Id             string `json:"id"`
	OrganizationId string `json:"organization_id" validate:"required"`
	Name           string `json:"name" validate:"required"`
	Address        string `json:"address" validate:"required"`
	Manager        string `json:"manager" validate:"required"`
	PhoneNumber    string `json:"phone_number" validate:"required"`
}

type IOffice interface {
	Insert(office *Office) error
	Get() ([]OfficeData, error)
}
