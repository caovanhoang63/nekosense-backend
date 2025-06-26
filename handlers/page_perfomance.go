package handlers

import (
	"context"
	"log"
	"nekosense-backend/models"
)

func (g *GinHandler) onPagePerformance(c context.Context, event *models.Event) error {
	_, err := g.db.Collection("performance").InsertOne(c, event)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
