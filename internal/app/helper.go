package app

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/db"
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/pkg"
	"path"
)

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
