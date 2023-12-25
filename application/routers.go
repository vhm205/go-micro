package application

import (
	"github.com/gin-gonic/gin"
	handler "github.com/vhm205/go-micro/shorturl-service"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Use(LoggerMiddleware)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world from /",
		})
	})

	LoadShortURLRouter(router)

	return router
}

func LoadShortURLRouter(router *gin.Engine) {
	ShortURLHandler := handler.ShortURLHandler{}

	g := router.Group("/api/v1/shorturl")

	g.GET("/", ShortURLHandler.ListShortURLs)
	g.GET("/{id}", ShortURLHandler.GetShortURL)
	g.POST("/", ShortURLHandler.CreateShortURL)
	g.PUT("/{id}", ShortURLHandler.UpdateShortURL)
	g.DELETE("/{id}", ShortURLHandler.DeleteShortURL)
}
