package service

import (
	"scsystem/internal/repo"
	"scsystem/internal/schema"
	"strconv"
)

type Stat struct {
	repo repo.IChart
}

func NewStat() IStat {
	return &Stat{
		repo: repo.NewChart(),
	}
}

func (s *Stat) GetChartData(roomId string) (*schema.ChartMetadata, error) {
	in, out, m, y, err := s.repo.GetData7Day(roomId)
	if err != nil {
		return nil, err
	}
	room, err := repo.NewRoom().GetByID(roomId)
	if err != nil {
		return nil, err
	}
	var data []schema.ChartData
	for i, v := range in {
		data = append(data,
			schema.ChartData{
				Value: v.Value,
				Label: v.Label,
			},
			schema.ChartData{
				Value: out[i].Value,
				Label: out[i].Label,
			},
		)
	}
	return &schema.ChartMetadata{
		Month:    m,
		Year:     strconv.Itoa(y),
		RoomID:   room.Id,
		RoomName: room.Name,
		Data:     data,
	}, nil

}

func (s *Stat) GetRoomData(roomId string) (*schema.RoomStat, error) {
	in, out, _, _, err := s.repo.GetData7Day(roomId)
	if err != nil {
		return nil, err
	}
	room, err := repo.NewRoom().GetByID(roomId)
	if err != nil {
		return nil, err
	}
	var inCount int
	if len(in) != 0 {
		inCount = in[0].Value
	}
	var outCount int
	if len(out) != 0 {
		outCount = out[0].Value
	}
	return &schema.RoomStat{
		RoomName: room.Name,
		RoomID:   roomId,
		In:       inCount,
		Out:      outCount,
		Total:    inCount + outCount,
	}, nil
}
