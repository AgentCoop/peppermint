package runtime

import (
	"github.com/AgentCoop/peppermint/internal/db"
	"gorm.io/gorm/schema"
	"strings"

	//	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	job "github.com/AgentCoop/go-work"
	//	"github.com/AgentCoop/peppermint/internal/runtime/cliparser"
)

func (r *runtime) InitTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		if len(r.dbFilename) != 0 {
			sqliteDb, err := gorm.Open(sqlite.Open(r.dbFilename), &gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
				NamingStrategy: schema.NamingStrategy{
					//TablePrefix: "t_",   // table name prefix, table for `User` would be `t_users`
					SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
					//NameReplacer: strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
				},
			})
			GlobalRegistry().SetDb(db.NewDb(sqliteDb, r.dbFilename))
			task.Assert(err)
		}
		err := r.CliParser.Run()
		task.Assert(err)
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

