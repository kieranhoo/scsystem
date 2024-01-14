package repo

import (
	"scsystem/internal/domain"
	"scsystem/pkg/database"
	"scsystem/pkg/database/queries"

	"gorm.io/gorm"
)

type Office struct {
	data *domain.Office
	conn *gorm.DB
}

func NewOffice() domain.IOffice {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &Office{
		data: &domain.Office{},
		conn: conn,
	}
}

func (o *Office) Insert(office *domain.Office) error {
	if err := o.conn.Exec(
		"INSERT INTO office (address, manager, name,organization_id, phone_number) VALUES (?, ?, ?, ?, ?);",
		office.Address,
		office.Manager,
		office.Name,
		office.OrganizationId,
		office.PhoneNumber,
	).Error; err != nil {
		return err
	}
	return nil
}

func (o *Office) Get() ([]domain.OfficeData, error) {
	var data []domain.OfficeData
	if err := o.conn.Raw(queries.OfficeData).Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
