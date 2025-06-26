package handlers

import (
	"context"
	"log"
	"nekosense-backend/models"
)

func (g *GinHandler) onHeatMap(c context.Context, event *models.Event) error {
	_, err := g.db.Collection("heatMap").InsertOne(c, event)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
