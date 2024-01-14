package repo

import (
	"scsystem/internal/domain"
	"scsystem/pkg/database"
	"scsystem/pkg/database/queries"

	"gorm.io/gorm"
)

type Room struct {
	data *domain.Room
	conn *gorm.DB
}

func NewRoom() domain.IRoom {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &Room{
		data: &domain.Room{},
		conn: conn,
	}
}

func (r *Room) Insert(room *domain.Room) error {
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

func (r *Room) Get() ([]domain.RoomData, error) {
	var roomData []domain.RoomData
	if err := r.conn.Raw(queries.RoomData).Scan(&roomData).Error; err != nil {
		return nil, err
	}
	return roomData, nil
}

func (r *Room) GetByID(roomId string) (*domain.Room, error) {
	if err := r.conn.Raw(
		"SELECT * FROM room WHERE id = ?", roomId,
	).Scan(r.data).Error; err != nil {
		return nil, err
	}
	return r.data, nil
}
