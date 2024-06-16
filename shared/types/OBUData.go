package types

type OBUData struct {
	ObuId     string  `json:"obu_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}
