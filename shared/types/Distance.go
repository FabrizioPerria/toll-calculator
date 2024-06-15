package types

type Distance struct {
	OBUID     int     `json:"obu_id"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}
