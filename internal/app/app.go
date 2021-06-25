package app

import (
	job "github.com/AgentCoop/go-work"
)

type EnvString string

const (
	DEV  EnvString = "dev"
	PROD EnvString = "prod"
	TEST EnvString = "test"
)

const (
	ENV_ROOT = "PEPPERMINT_ROOT"
	ENV      = "PEPPERMINT"
)

func (e EnvString) String() string {
	return string(e)
}


type app struct {
	appDir     string
	profile    AppProfile
	appJob     job.Job
}

type AppProfile struct {
	DbFilename string
}

func (app *app) WithDb() {
	app.appJob.AddTask(app.InitDbTask)
}

func (app *app) Job() job.Job {
	return app.appJob
}

func (app *app) RootDir() string {
	return app.appDir
}
