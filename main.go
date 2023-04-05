package main

import (
	"os"

	"github.com/parinyapt/lmwn_assessment_2023/config"
	"github.com/parinyapt/lmwn_assessment_2023/routes"
)

func main() {
	config.FlagSetup()
	if os.Getenv("DEPLOY_MODE") == "development" {
		config.EnvironmentFileSetup()
	}
	if os.Getenv("DEPLOY_MODE") == "production" {
		config.GinReleaseModeSetup()
	}
	config.EnvironmentVariableCheck()
	config.TimezoneSetup()
	
	routes.Setup()
}