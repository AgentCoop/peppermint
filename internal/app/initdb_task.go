package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/db"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"path"
)

func (a *app) InitDbTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		pathname := path.Join(a.appDir, a.profile.DbFilename)
		sqliteDb, err := gorm.Open(sqlite.Open(pathname), &gorm.Config{
			//DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			},
		})
		task.Assert(err)
		runtime.GlobalRegistry().SetDb(db.NewDb(sqliteDb, pathname))
	}
	return init, nil, nil
}

