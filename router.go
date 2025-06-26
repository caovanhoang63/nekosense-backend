package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"nekosense-backend/handlers"
)

func MainRouter(router *gin.Engine, g *handlers.GinHandler) {
	router.Use(cors.Default())
	router.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/event", g.OnNekoSenseEvent())
	router.GET("/clicks", g.GetClicks())
	router.GET("/performance", g.GetPagePerformances())
	router.GET("/heat-map", g.GetHeatMaps())

}
