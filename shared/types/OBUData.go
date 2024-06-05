package types

type OBUData struct {
	OBUID     int     `json:"obu_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
