package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/pkg"
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
	db      pkg.Db
	appDir  string
	profile AppProfile
	appJob  job.Job
	cliOpts interface{}
}

type AppProfile struct {
	DbFilename string
}

func (app *app) Db() pkg.Db {
	return app.db
}

func (app *app) WithDb(db pkg.Db) {
	app.db = db
}

func (app *app) Job() job.Job {
	return app.appJob
}

func (app *app) RootDir() string {
	return app.appDir
}
