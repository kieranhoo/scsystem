package domain

type Chart struct {
	Id       string `json:"id" gorm:"column:id"`
	Time     string `json:"time" gorm:"column:time"`
	RoomId   string `json:"room_id" gorm:"column:room_id"`
	InCount  int    `json:"in_count" gorm:"column:in_count"`
	OutCount int    `json:"out_count" gorm:"column:out_count"`
}

type ChartData struct {
	Value int    `json:"value"` // counter in/out
	Label string `json:"label"` // day
}

type ChartDataValue struct {
	In  ChartData `json:"in"`
	Out ChartData `json:"out"`
}

type ChartMetadata struct {
	Month    string           `json:"month"`
	Year     string           `json:"year"`
	RoomName string           `json:"room_name"`
	RoomID   string           `json:"room_id"`
	Data     []ChartDataValue `json:"data"`
}

type RoomStatData struct {
	RoomName string `json:"room_name"`
	RoomID   string `json:"room_id"`
	In       int    `json:"in"`
	Out      int    `json:"out"`
	Total    int    `json:"total"`
}
type RoomStat struct {
	Total int            `json:"total"`
	Data  []RoomStatData `json:"data"`
}

type IChart interface {
	UpdateChartData(roomId, activityType string) error
	GetData7Day(roomId string) (_in []ChartData, _out []ChartData, month string, year int, err error)
}
