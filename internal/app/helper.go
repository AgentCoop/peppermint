package app

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/pkg"
	"os"
)

func ProfileFromEnv() AppProfile {
	var profile AppProfile
	env := os.Getenv(ENV)
	switch {
	case len(env) == 0:
		profile = AppProfiles[PROD]
	default:
		p, ok := AppProfiles[EnvString(env)]
		if !ok {
			panic(fmt.Errorf("invalid PEPPERMINT environment variable value %s", env))
		}
		profile = p
	}
	return profile
}

func AppInit(app pkg.App, t job.Task) {
	rt := runtime.GlobalRegistry().Runtime()
	err := rt.CliParser().Run()
	t.Assert(err)

	err = utils.FS_FileOrDirExists(app.RootDir())
	t.Assert(err)
	t.Done()
}
