package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"nekosense-backend/models"
)

type GinHandler struct {
	db *mongo.Database
}

func NewGinHandler(db *mongo.Database) *GinHandler {
	return &GinHandler{db: db}
}

func (g *GinHandler) OnNekoSenseEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var event models.Event
		err := c.ShouldBindBodyWithJSON(&event)
		if err != nil {
			log.Println(err)
			return
		}
		err = nil
		switch event.Event {
		case "performance":
			err = g.onPagePerformance(c.Request.Context(), &event)
			break
		case "heatMap":
			err = g.onHeatMap(c.Request.Context(), &event)
			break
		case "click":
			err = g.onClick(c.Request.Context(), &event)
			break
		case "hover-to-click":
			err = g.onHoverToClick(c.Request.Context(), &event)
			break
		case "pageView":
			err = g.onPageView(c.Request.Context(), &event)
			break
		case "timeOnPage":
			err = g.onTimeOnPage(c.Request.Context(), &event)
			break
		default:
			c.JSON(400, gin.H{
				"status":  "error",
				"message": "event not supported",
			})
			return
		}

		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"status":  "error",
				"message": "event not processed",
			})
			return
		}
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "event processed",
		})
	}
}
