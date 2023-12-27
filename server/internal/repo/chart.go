package repo

import (
	"errors"
	"log"
	"scsystem/internal/model"
	"scsystem/internal/schema"
	"scsystem/pkg/database"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Chart struct {
	data *model.Chart
	conn *gorm.DB
}

func NewChart() IChart {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}
	return &Chart{
		data: &model.Chart{},
		conn: conn,
	}
}
func (chart *Chart) UpdateChartData(roomId, activityType string) error {
	timeNow := time.Now().UTC()
	//TODO: record in Chart table is exits ?
	if err := chart.conn.Raw(
		"SELECT * FROM chart WHERE room_id = ? AND DATE(time) = ?",
		roomId,
		timeNow.Format(time.DateOnly),
	).Scan(chart.data).Error; err != nil {
		return err
	}
	switch activityType {
	case "in":
		chart.data.InCount++
	case "out":
		chart.data.OutCount++
	default:
		return errors.New("unknown activity type")
	}
	//TODO: if chartID is not null, update int_count, out_count in Chart table
	if chart.data.Id != "" || chart.data.RoomId != "" {
		if err := chart.conn.Exec(
			"UPDATE chart SET in_count = ?, out_count = ? WHERE room_id = ? AND time = ? ;",
			chart.data.InCount, chart.data.OutCount, roomId, chart.data.Time,
		).Error; err != nil {
			return err
		}
		return nil
	}
	//TODO: else, create new record in chart table
	if err := chart.conn.Exec(
		"INSERT INTO chart (in_count, out_count, time, room_id) VALUES (?, ?, ?, ?);",
		chart.data.InCount,
		chart.data.OutCount,
		timeNow.Format(time.DateTime),
		roomId,
	).Error; err != nil {
		return err
	}
	return nil
}

func (chart *Chart) GetData7Day(roomId string) (_in []schema.ChartData, _out []schema.ChartData, month string, year int, err error) {
	var chartData []model.Chart
	_in = []schema.ChartData{}
	_out = []schema.ChartData{}
	if err = chart.conn.Raw(
		"SELECT * FROM chart WHERE room_id = ? ORDER BY id DESC LIMIT 7",
		roomId,
	).Scan(&chartData).Error; err != nil {
		return _in, _out, "", -1, err
	}
	for _, v := range chartData {
		date, err := time.Parse(time.DateTime, v.Time)
		month = date.Month().String()
		year = date.Year()
		if err != nil {
			log.Fatal(err)
		}
		_in = append(_in, schema.ChartData{
			Label: strconv.Itoa(date.Day()),
			Value: v.InCount,
		})
		_out = append(_out, schema.ChartData{
			Label: strconv.Itoa(date.Day()),
			Value: v.OutCount,
		})
	}
	return _in, _out, month, year, nil
}
