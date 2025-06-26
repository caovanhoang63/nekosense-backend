package models

type Event struct {
	Event string `json:"event" bson:"event"`
	Ele   string `json:"ele" bson:"ele"`
	//Timestamp *time.Time `json:"timestamp" bson:"timestamp"`
	Url  string      `json:"url" bson:"url"`
	Data interface{} `json:"data" bson:"data"`
}
