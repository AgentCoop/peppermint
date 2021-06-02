package node

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/db"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path"
)

func (a *app) initDb() error {
	if len(a.dbFilename) != 0 {
		pathname := path.Join(*a.appDir, a.dbFilename)
		sqliteDb, err := gorm.Open(sqlite.Open(pathname), &gorm.Config{
			//DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			},
		})
		runtime.GlobalRegistry().SetDb(db.NewDb(sqliteDb, pathname))
		return err
	}
	return nil
}

func (a *app) InitTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		err := a.CliParser().Run()
		task.Assert(err)

		// Default application directory is the current working directory
		if len(*a.appDir) == 0 {
			*a.appDir, err = os.Getwd()
			task.Assert(err)
		}
		err = utils.FS_FileOrDirExists(*a.appDir)
		task.Assert(err)

		err = a.initDb()
		task.Assert(err)

		// Fetch node configuration once DB is initialized
		err = a.NodeConfigurator().Fetch()
		task.Assert(err)

		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}
