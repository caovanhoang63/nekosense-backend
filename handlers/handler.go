package handlers

import "go.mongodb.org/mongo-driver/mongo"

type GinHandler struct {
	db *mongo.Database
}

func NewGinHandler(db *mongo.Database) *GinHandler {
	return &GinHandler{db: db}
}
