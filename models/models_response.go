package models

type CovidSummaryResponse struct {
	Province ProvinceData `json:"Province"`
	AgeGroup AgeGroupData `json:"AgeGroup"`
}

type ProvinceData map[string]int64

type AgeGroupData struct {
	Age0to30  int64 `json:"0-30"`
	Age31to60 int64 `json:"31-60"`
	Age61up   int64 `json:"61+"`
	Unknown   int64 `json:"N/A"`
}
