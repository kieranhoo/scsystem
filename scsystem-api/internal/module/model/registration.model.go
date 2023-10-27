package model

import (
	"gorm.io/gorm"
	"qrcheckin/internal/module/types"
	"qrcheckin/pkg/database"
)

type Registration struct {
	data *types.Registration
	conn *gorm.DB
}

func NewRegistration() types.IRegistration {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &Registration{
		data: &types.Registration{},
		conn: conn,
	}
}

func (res *Registration) Insert(_res *types.Registration) error {
	if err := res.conn.Exec(
		"INSERT INTO registration (registration_time, supervisor, user_id, room_id, start_day, end_day) VALUES (?, ?, ?, ?, ?, ?);",
		_res.RegistrationTime, _res.Supervisor, _res.UserID, _res.RoomId, _res.StartDay, _res.EndDay).Error; err != nil {
		return err
	}
	return nil
}

func (res *Registration) GetByID(id string) (*types.Registration, error) {
	if err := res.conn.Raw("SELECT * FROM registration WHERE id = ?", id).Scan(res.data).Error; err != nil {
		return nil, err
	}
	return res.data, nil
}

func (res *Registration) GetByUserIdAndRoom(userId, roomId string) (*types.Registration, error) {
	if err := res.conn.Raw(
		"SELECT * FROM registration WHERE user_id = ? AND room_id = ?",
		userId, roomId,
	).Scan(res.data).Error; err != nil {
		return nil, err
	}
	return res.data, nil
}

func (res *Registration) UpdateByIDAndRoom(_res *types.Registration) error {
	if err := res.conn.Exec(
		"UPDATE registration SET registration_time = ?, supervisor = ?, start_day = ?, end_day = ? WHERE user_id = ? AND room_id = ? AND id = ?;",
		_res.RegistrationTime, _res.Supervisor, _res.StartDay, _res.EndDay, _res.UserID, _res.RoomId, _res.Id).Error; err != nil {
		return err
	}
	return nil
}

func (res *Registration) Latest() (*types.Registration, error) {
	if err := res.conn.Raw("SELECT * FROM registration ORDER BY id desc LIMIT 1;").Scan(res.data).Error; err != nil {
		return nil, err
	}
	return res.data, nil
}

func (res *Registration) Empty() bool {
	return res.data.UserID == ""
}
