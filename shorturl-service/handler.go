package handler

import "github.com/gin-gonic/gin"

type ShortURLHandler struct{}

func (h *ShortURLHandler) CreateShortURL(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "CreateShortURL handler",
	})
}

func (h *ShortURLHandler) GetShortURL(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetShortURL handler",
	})
}

func (h *ShortURLHandler) ListShortURLs(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ListShortURLs handler",
	})
}

func (h *ShortURLHandler) DeleteShortURL(c *gin.Context) {
	c.JSON(204, gin.H{
		"message": "DeleteShortURL handler",
	})
}

func (h *ShortURLHandler) UpdateShortURL(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateShortURL handler",
	})
}
