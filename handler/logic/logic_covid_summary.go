package handler

import (
	"github.com/parinyapt/lmwn_assessment_2023/models"
)

func CovidCaseDataSummary(responseData models.CovidCaseApiFetchResponse) (models.ProvinceData, models.AgeGroupData) {
	var provinceCount = make(models.ProvinceData)
	var ageGroupCount models.AgeGroupData

	for _, data := range responseData.Data {
		if data.Province != nil {
			provinceCount[*data.Province] = provinceCount[*data.Province] + 1
		}

		if data.Age != nil {
			if *data.Age >= 0 && *data.Age <= 30 {
				ageGroupCount.Age0to30 = ageGroupCount.Age0to30 + 1
			}else if *data.Age >= 31 && *data.Age <= 60 {
				ageGroupCount.Age31to60 = ageGroupCount.Age31to60 + 1
			}else if *data.Age >= 61 {
				ageGroupCount.Age61up = ageGroupCount.Age61up + 1
			}else{
				ageGroupCount.Unknown = ageGroupCount.Unknown + 1
			}
		}else{
			ageGroupCount.Unknown = ageGroupCount.Unknown + 1
		}
	}
	
	return provinceCount, ageGroupCount
}