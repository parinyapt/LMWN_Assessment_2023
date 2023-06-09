package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckApiStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API is running",
		"timestamp": time.Now().Unix(),
	})
}