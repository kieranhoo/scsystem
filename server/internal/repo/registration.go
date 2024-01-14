package repo

import (
	"scsystem/internal/domain"
	"scsystem/pkg/database"
	"scsystem/pkg/database/queries"

	"gorm.io/gorm"
)

type Registration struct {
	data *domain.Registration
	conn *gorm.DB
}

func NewRegistration() domain.IRegistration {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &Registration{
		data: &domain.Registration{},
		conn: conn,
	}
}

func (res *Registration) Insert(_res *domain.Registration) error {
	if err := res.conn.Exec(
		"INSERT INTO registration (registration_time, supervisor, user_id, room_id, start_day, end_day) VALUES (?, ?, ?, ?, ?, ?);",
		_res.RegistrationTime, _res.Supervisor, _res.UserID, _res.RoomId, _res.StartDay, _res.EndDay).Error; err != nil {
		return err
	}
	return nil
}

func (res *Registration) GetByID(id string) (*domain.Registration, error) {
	if err := res.conn.Raw("SELECT * FROM registration WHERE id = ?", id).Scan(res.data).Error; err != nil {
		return nil, err
	}
	return res.data, nil
}

func (res *Registration) GetByUserIdAndRoom(userId, roomId string) (*domain.Registration, error) {
	if err := res.conn.Raw(
		"SELECT * FROM registration WHERE user_id = ? AND room_id = ?",
		userId, roomId,
	).Scan(res.data).Error; err != nil {
		return nil, err
	}
	return res.data, nil
}

func (res *Registration) UpdateByIDAndRoom(_res *domain.Registration) error {
	if err := res.conn.Exec(
		"UPDATE registration SET registration_time = ?, supervisor = ?, start_day = ?, end_day = ? WHERE user_id = ? AND room_id = ? AND id = ?;",
		_res.RegistrationTime, _res.Supervisor, _res.StartDay, _res.EndDay, _res.UserID, _res.RoomId, _res.Id).Error; err != nil {
		return err
	}
	return nil
}

func (res *Registration) Latest() (*domain.Registration, error) {
	if err := res.conn.Raw("SELECT * FROM registration ORDER BY id desc LIMIT 1;").Scan(res.data).Error; err != nil {
		return nil, err
	}
	return res.data, nil
}

func (res *Registration) Empty() bool {
	return res.data.UserID == ""
}

func (res *Registration) RegistrationLatest(studentId, roomId string) (*domain.UserRoomData, error) {
	RoomData := new(domain.UserRoomData)
	if err := res.conn.Raw(queries.RegistrationLatest, studentId, roomId).Scan(RoomData).Error; err != nil {
		return nil, err
	}
	history, err := NewHistory().Latest(RoomData.RegistrationId)
	if err != nil {
		return nil, err
	}
	switch history.ActivityType {
	case "":
		RoomData.ActivityType = "no access"
	case "out":
		RoomData.ActivityType = "out"
	case "in":
		RoomData.ActivityType = "in"
	default:
		panic("unknown activity type")
	}
	return RoomData, nil
}

func (res *Registration) OneWeekBefore() ([]domain.Registration, error) {
	var registration []domain.Registration
	if err := res.conn.Raw(queries.RegistrationBefore1Week).Scan(&registration).Error; err != nil {
		return nil, err
	}
	return registration, nil
}

func (res *Registration) DeleteById(id string) error {
	return res.conn.Exec("DELETE FROM registration WHERE id = ?", id).Error
}
