package test

import (
	"github.com/AgentCoop/peppermint/internal/app"
)


func NewTestApp() *appTest {
	prof := app.ProfileFromEnv()
	appTest := new(appTest)
	appTest.App = app.NewApp(prof, &options)
	appTest.Job().AddOneshotTask(appTest.InitTask)
	appTest.Job().AddTask(appTest.ParserTask)
	return appTest
}
