package schema

type Error struct {
	Success bool        `json:"status"`
	Msg     string      `json:"msg"`
}
