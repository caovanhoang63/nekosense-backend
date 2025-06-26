package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"nekosense-backend/models"
	"net/http"
)

func (g *GinHandler) onPagePerformance(c context.Context, event *models.Event) error {
	_, err := g.db.Collection("performance").InsertOne(c, event)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (g *GinHandler) GetPagePerformances() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use an empty BSON document to find all documents.
		cursor, err := g.db.Collection("performance").Find(c, bson.M{})
		if err != nil {
			log.Println("Failed to execute find command for performance:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve performance data"})
			return
		}
		// Ensure the cursor is closed when the function finishes.
		defer cursor.Close(c)

		// Create a slice to hold the decoded documents.
		var performances []models.Event

		// Decode all documents from the cursor into the slice.
		if err = cursor.All(c, &performances); err != nil {
			log.Println("Failed to decode performance documents:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process performance data"})
			return
		}

		// It's good practice to return an empty array instead of null
		// if no documents are found.
		if performances == nil {
			performances = make([]models.Event, 0)
		}

		// Return the slice of performance data with a 200 OK status.
		c.JSON(http.StatusOK, performances)
	}
}
