package service

import (
	"scsystem/internal/domain"
	"scsystem/internal/repo"
	"strconv"
)

type Stat struct {
	repo domain.IChart
}

func NewStat() domain.IStatService {
	return &Stat{
		repo: repo.NewChart(),
	}
}

func (s *Stat) GetChartData(roomId string) (*domain.ChartMetadata, error) {
	in, out, m, y, err := s.repo.GetData7Day(roomId)
	if err != nil {
		return nil, err
	}
	room, err := repo.NewRoom().GetByID(roomId)
	if err != nil {
		return nil, err
	}
	var data []domain.ChartDataValue
	for i, v := range in {
		data = append(data,
			domain.ChartDataValue{
				In: domain.ChartData{
					Value: v.Value,
					Label: v.Label,
				},
				Out: domain.ChartData{
					Value: out[i].Value,
					Label: out[i].Label,
				},
			},
		)
	}
	return &domain.ChartMetadata{
		Month:    m,
		Year:     strconv.Itoa(y),
		RoomID:   room.Id,
		RoomName: room.Name,
		Data:     data,
	}, nil

}

func (s *Stat) GetRoomData() (*domain.RoomStat, error) {
	var data []domain.RoomStatData
	room, err := repo.NewRoom().Get()
	if err != nil {
		return nil, err
	}
	for _, r := range room {
		in, out, _, _, err := s.repo.GetData7Day(r.RoomID)
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
		data = append(data, domain.RoomStatData{
			RoomName: r.RoomName,
			RoomID:   r.RoomID,
			In:       inCount,
			Out:      outCount,
			Total:    inCount + outCount,
		})
	}
	return &domain.RoomStat{
		Total: len(room),
		Data:  data,
	}, nil
}
