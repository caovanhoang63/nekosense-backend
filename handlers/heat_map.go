package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"nekosense-backend/models"
	"net/http"
)

func (g *GinHandler) onHeatMap(c context.Context, event *models.Event) error {
	_, err := g.db.Collection("heatMap").InsertOne(c, event)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

type SimplifiedHeatMapPoint struct {
	X         int64 `json:"x" bson:"x"`
	Y         int64 `json:"y" bson:"y"`
	Timestamp int64 `json:"timestamp" bson:"timestamp"`
}

func (g *GinHandler) GetHeatMaps() gin.HandlerFunc {
	return func(c *gin.Context) {
		pipeline := mongo.Pipeline{
			bson.D{{Key: "$match", Value: bson.D{{Key: "event", Value: "heatMap"}}}},
			bson.D{{Key: "$sort", Value: bson.D{{Key: "_id", Value: -1}}}},
			bson.D{{Key: "$limit", Value: 10}},
			bson.D{{Key: "$unwind", Value: "$data.mousePositions"}},
			bson.D{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: "$data.mousePositions"}}}},
		}

		cursor, err := g.db.Collection("heatMap").Aggregate(c, pipeline)
		if err != nil {
			log.Println("Failed to execute aggregation for heatmap:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve heatmap data"})
			return
		}
		defer cursor.Close(c)

		var results []SimplifiedHeatMapPoint
		if err = cursor.All(c, &results); err != nil {
			log.Println("Failed to decode heatmap documents:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process heatmap data"})
			return
		}

		if results == nil {
			results = make([]SimplifiedHeatMapPoint, 0)
		}

		c.JSON(http.StatusOK, results)
	}
}
