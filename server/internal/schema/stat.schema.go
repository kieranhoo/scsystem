package schema

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
