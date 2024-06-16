package types

type Invoice struct {
	ObuId     string  `json:"obu_id"`
	Amount    float64 `json:"amount"`
	Distance  float64 `json:"distance"`
	Timestamp int64   `json:"timestamp"`
}
