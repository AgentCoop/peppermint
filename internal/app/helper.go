package app

import (
	"fmt"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/db"
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/pkg"
	"os"
	"path"
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

func InitDb(app pkg.App, dbFilename string) (pkg.Db, error) {
	err := db.Init(app)
	if err != nil {
		return nil, err
	}
	pathname := path.Join(app.RootDir(), "db", dbFilename)
	return db.Open(pathname)
}

func AppInit(app pkg.App) error {
	rt := runtime.GlobalRegistry().Runtime()
	if err := rt.CliParser().Run(); err != nil {
		return err
	}
	if err := utils.FS_FileOrDirExists(app.RootDir()); err != nil {
		return err
	}
	return nil
}
