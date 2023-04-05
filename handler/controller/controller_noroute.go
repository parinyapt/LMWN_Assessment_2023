package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "API is running but no route found",
		"timestamp": time.Now().Unix(),
	})
}