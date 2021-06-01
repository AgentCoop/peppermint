package runtime

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func (r *runtime) initDb() error {
	if len(r.dbFilename) != 0 {
		sqliteDb, err := gorm.Open(sqlite.Open(r.dbFilename), &gorm.Config{
			//DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			},
		})
		GlobalRegistry().SetDb(db.NewDb(sqliteDb, r.dbFilename))
		return err
	}
	return nil
}

func (r *runtime) InitTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		err := r.initDb()
		task.Assert(err)

		err = r.parser.Run()
		task.Assert(err)

		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

