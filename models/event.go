package models

type Event struct {
	Event     string      `json:"event" bson:"event"`
	Ele       string      `json:"ele" bson:"ele"`
	EleId     string      `json:"eleId" bson:"eleId"`
	Timestamp int64       `json:"timestamp" bson:"timestamp"`
	Url       string      `json:"url" bson:"url"`
	Data      interface{} `json:"data" bson:"data"`
}
