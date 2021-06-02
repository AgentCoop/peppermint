package runtime

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/db"
	"github.com/AgentCoop/peppermint/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path"
)

func (r *runtime) initDb() error {
	if len(r.dbFilename) != 0 {
		pathname := path.Join(r.AppDir(), r.dbFilename)
		sqliteDb, err := gorm.Open(sqlite.Open(pathname), &gorm.Config{
			//DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			},
		})
		GlobalRegistry().SetDb(db.NewDb(sqliteDb, pathname))
		return err
	}
	return nil
}

func (r *runtime) InitTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		err := r.parser.Run()
		task.Assert(err)

		// Default application directory is the current working directory
		if len(r.AppDir()) == 0 {
			*r.appDir, err = os.Getwd()
			task.Assert(err)
		}
		err = utils.FS_FileOrDirExists(r.AppDir())
		task.Assert(err)

		err = r.initDb()
		task.Assert(err)

		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

