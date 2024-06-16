package types

type Invoice struct {
	ObuID     int     `json:"obu_id"`
	Amount    float64 `json:"amount"`
	Distance  float64 `json:"distance"`
	Timestamp int64   `json:"timestamp"`
}
