package routes

import (
	"github.com/gin-gonic/gin"
)

func configApiRoutes(router *gin.Engine) {
	covid := router.Group("/covid")
	{
		covid.GET("/summary", GetSummary)
	}
}