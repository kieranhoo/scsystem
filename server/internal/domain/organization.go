package domain

type Organization struct {
	Id          string `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Email       string `json:"email" gorm:"column:email"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Website     string `json:"website" gorm:"column:website"`
	Head        string `json:"head" gorm:"column:head"`
}

type NewOrganizationRequest struct {
	Name        string `json:"name" gorm:"column:name" validate:"required"`
	Email       string `json:"email" gorm:"column:email" validate:"required"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number" validate:"required"`
	Website     string `json:"website" gorm:"column:website" validate:"required"`
	Head        string `json:"head" gorm:"column:head" validate:"required"`
}

type IOrganization interface {
	Insert(org *Organization) error
	Get() ([]Organization, error)
}
