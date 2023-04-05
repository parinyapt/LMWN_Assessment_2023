package models

type CovidCaseApiFetchResponse struct {
	Data []CovidCaseApiFetchResponseData `json:"Data"`
}

type CovidCaseApiFetchResponseData struct {
	ConfirmDate    *string `json:"ConfirmDate"`
	No             *int64  `json:"No"`
	Age            *int64  `json:"Age"`
	Gender         *string `json:"Gender"`
	GenderEn       *string `json:"GenderEn"`
	Nation         *string `json:"Nation"`
	NationEn       *string `json:"NationEn"`
	Province       *string `json:"Province"`
	ProvinceID     *int64  `json:"ProvinceId"`
	District       *string `json:"District"`
	ProvinceEn     *string `json:"ProvinceEn"`
	StatQuarantine *int64  `json:"StatQuarantine"`
}
