package repo

import (
	"scsystem/internal/model"
	"scsystem/pkg/database"

	"gorm.io/gorm"
)

type Organization struct {
	data *model.Organization
	conn *gorm.DB
}

func NewOrganization() IOrganization {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &Organization{
		data: &model.Organization{},
		conn: conn,
	}
}

func (o *Organization) Insert(org *model.Organization) error {
	if err := o.conn.Exec(
		"INSERT INTO organization (name, email, head, phone_number, website) VALUES (?, ?, ?, ?, ?);",
		org.Name,
		org.Email,
		org.Head,
		org.PhoneNumber,
		org.Website,
	).Error; err != nil {
		return err
	}
	return nil
}

func (o *Organization) Get() ([]model.Organization, error) {
	var orgData []model.Organization
	if err := o.conn.Raw("SELECT * FROM organization").Scan(&orgData).Error; err != nil {
		return nil, err
	}
	return orgData, nil
}
