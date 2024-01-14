package repo

import (
	"scsystem/internal/domain"
	"scsystem/pkg/database"

	"gorm.io/gorm"
)

type Organization struct {
	data *domain.Organization
	conn *gorm.DB
}

func NewOrganization() domain.IOrganization {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &Organization{
		data: &domain.Organization{},
		conn: conn,
	}
}

func (o *Organization) Insert(org *domain.Organization) error {
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

func (o *Organization) Get() ([]domain.Organization, error) {
	var orgData []domain.Organization
	if err := o.conn.Raw("SELECT * FROM organization").Scan(&orgData).Error; err != nil {
		return nil, err
	}
	return orgData, nil
}
