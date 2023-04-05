package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/parinyapt/lmwn_assessment_2023/handler/logic"
	"github.com/parinyapt/lmwn_assessment_2023/models"
)

func CovidSummaryHandler(c *gin.Context) {
	fetchData, err := handler.CovidCaseApiFetch()
	if err != nil {
		log.Printf("[ Error Fetch Data in CovidSummaryHandler() ] -> %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if len(fetchData.Data) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	provinceSummary, ageGroupSummary := handler.CovidCaseDataSummary(fetchData)
	c.JSON(http.StatusOK, 
		models.CovidSummaryResponse{
			Province: provinceSummary,
			AgeGroup: ageGroupSummary,
		},
	)
}