package repo

import (
	"scsystem/internal/model"
	"scsystem/internal/schema"
	"scsystem/pkg/database"
	"scsystem/pkg/database/queries"

	"gorm.io/gorm"
)

type Office struct {
	data *model.Office
	conn *gorm.DB
}

func NewOffice() IOffice {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &Office{
		data: &model.Office{},
		conn: conn,
	}
}

func (o *Office) Insert(office *model.Office) error {
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

func (o *Office) Get() ([]schema.OfficeData, error) {
	var data []schema.OfficeData
	if err := o.conn.Raw(queries.OfficeData).Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
