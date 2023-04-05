package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"

	"github.com/parinyapt/lmwn_assessment_2023/models"
)

func CovidCaseApiFetch() (models.CovidCaseApiFetchResponse, error) {
	errorMessage := errors.New("[ CovidCaseApiFetch() ]")

	response, err := http.Get(os.Getenv("API_COVID_CASE_URL"))
	if err != nil {
		errorMessage = errors.Wrap(errorMessage, fmt.Sprintf("[ API Fetch Error ] -> %s", err.Error()))
		return models.CovidCaseApiFetchResponse{}, errorMessage
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorMessage = errors.Wrap(errorMessage, fmt.Sprintf("[ ioutil.ReadAll Response Body Error ] -> %s", err.Error()))
		return models.CovidCaseApiFetchResponse{}, errorMessage
	}
	// fmt.Println(string(responseData))

	var responseData models.CovidCaseApiFetchResponse
  if err = json.Unmarshal(responseBody, &responseData); err == nil {
		return responseData, nil
	}else{
		errorMessage = errors.Wrap(errorMessage, fmt.Sprintf("[ JSON Unmarshal Error ] -> %s", err.Error()))
		return models.CovidCaseApiFetchResponse{}, errorMessage
	}
}