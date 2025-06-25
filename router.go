package main

import (
	"github.com/gin-gonic/gin"
	"nekosense-backend/handlers"
)

func MainRouter(router *gin.Engine, g *handlers.GinHandler) {
	router.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
