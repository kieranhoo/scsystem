package repo

import (
	"scsystem/internal/model"
	"scsystem/internal/schema"
	"scsystem/pkg/database"
	"scsystem/pkg/database/queries"

	"gorm.io/gorm"
)

type Room struct {
	data *model.Room
	conn *gorm.DB
}

func NewRoom() IRoom {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &Room{
		data: &model.Room{},
		conn: conn,
	}
}

func (r *Room) Insert(room *model.Room) error {
	if err := r.conn.Exec(
		"INSERT INTO room (office_id, name, description) VALUES (?, ?, ?);",
		room.OfficeID,
		room.Name,
		room.Description,
	).Error; err != nil {
		return err
	}
	return nil
}

func (r *Room) Get() ([]schema.RoomData, error) {
	var roomData []schema.RoomData
	if err := r.conn.Raw(queries.RoomData).Scan(&roomData).Error; err != nil {
		return nil, err
	}
	return roomData, nil
}
