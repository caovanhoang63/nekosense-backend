package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"nekosense-backend/models"
	"net/http"
)

func (g *GinHandler) onClick(c context.Context, event *models.Event) error {
	_, err := g.db.Collection("click").InsertOne(c, event)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (g *GinHandler) GetClicks() gin.HandlerFunc {
	return func(c *gin.Context) {
		cursor, err := g.db.Collection("click").Find(c, bson.M{})
		if err != nil {
			log.Println("Failed to execute find command:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve clicks"})
			return
		}
		defer cursor.Close(c)

		var clicks []models.Event

		if err = cursor.All(c, &clicks); err != nil {
			log.Println("Failed to decode documents:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process click data"})
			return
		}

		if clicks == nil {
			clicks = make([]models.Event, 0)
		}

		c.JSON(http.StatusOK, clicks)
	}
}
