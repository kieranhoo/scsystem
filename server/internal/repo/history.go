package repo

import (
	"scsystem/internal/model"
	"scsystem/pkg/database"
	"time"

	"gorm.io/gorm"
)

type History struct {
	data *model.History
	conn *gorm.DB
}

func NewHistory() IHistory {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &History{
		data: &model.History{},
		conn: conn,
	}
}

func (his *History) Insert(_his *model.History) error {
	_time := time.Now().UTC().Format(time.DateTime)
	if err := his.conn.Exec(
		"INSERT INTO history (registration_id, activity_type, time, admin_id) VALUES (?, ?, ?, ?);",
		_his.RegistrationId, _his.ActivityType, _time, _his.AdminId).Error; err != nil {
		return err
	}
	return nil
}

func (his *History) Latest(registrationId string) (*model.History, error) {
	if err := his.conn.Raw(
		"SELECT * FROM history WHERE registration_id = ? ORDER BY id DESC LIMIT 1;",
		registrationId,
	).Scan(his.data).Error; err != nil {
		return nil, err
	}
	return his.data, nil
}

func (his *History) Empty() bool {
	return his.data.ActivityType == ""
}

func (his *History) GetList(limit string) ([]model.History, error) {
	var histories []model.History
	if err := his.conn.Raw(
		"SELECT * FROM history ORDER BY id DESC LIMIT ?;",
		limit,
	).Scan(&histories).Error; err != nil {
		return nil, err
	}
	return histories, nil
}

func (his *History) GetActivityType() string {
	return his.data.ActivityType
}
