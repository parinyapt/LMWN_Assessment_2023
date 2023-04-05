package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/parinyapt/lmwn_assessment_2023/handler/controller"
)

func configApiRoutes(router *gin.Engine) {
	// No Route 404 Notfound
	router.NoRoute(handler.NoRoute)

	// Health Check
	router.GET("/health", handler.CheckApiStatus)

	// Covid Group
	covid := router.Group("/covid")
	{
		// Summary Endpoint
		covid.GET("/summary", handler.CovidSummaryHandler)
	}
}
