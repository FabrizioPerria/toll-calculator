package types

type Distance struct {
	ObuId     string  `json:"obu_id"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}
