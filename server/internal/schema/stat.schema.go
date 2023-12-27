package schema

type ChartData struct {
	Value int    `json:"value"` // counter in/out
	Label string `json:"label"` // day
}

type ChartMetadata struct {
	Month    string      `json:"month"`
	Year     string      `json:"year"`
	RoomName string      `json:"room_name"`
	RoomID   string      `json:"room_id"`
	Data     []ChartData `json:"data"`
}

type RoomStat struct {
	RoomName string `json:"room_name"`
	RoomID   string `json:"room_id"`
	In       int    `json:"in"`
	Out      int    `json:"out"`
	Total    int    `json:"total"`
}
