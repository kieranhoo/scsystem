package schema

type ChartData struct {
	Value int    `json:"value"` // counter in/out
	Label string `json:"label"` // day
}

type ChartMetadata struct {
	Month string      `json:"month"`
	Year  string      `json:"year"`
	Data  []ChartData `json:"data"`
}

type RoomStat struct {
	Room   string `json:"room"`
	RoomID string `json:"room_id"`
	In     string `json:"in"`
	Out    string `json:"out"`
}
