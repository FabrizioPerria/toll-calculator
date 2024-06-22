package types

type Distance struct {
	ObuId     string  `json:"obu_id" bson:"obu_id"`
	Value     float64 `json:"value" bson:"value"`
	Timestamp int64   `json:"timestamp" bson:"timestamp"`
}
