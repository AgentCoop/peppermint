package db

import (
	"github.com/AgentCoop/peppermint/internal/utils/fs"
	"github.com/AgentCoop/peppermint/pkg"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"path"
)

func NewDb(handle *gorm.DB, dbFilename string) sqlitedb {
	return sqlitedb{
		handle,
		dbFilename,
	}
}

func Init(app pkg.App) error {
	dbRoot := path.Join(app.RootDir(), "db")
	if !fs.IfExists(dbRoot) {
		err := os.Mkdir(dbRoot, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func Open(pathname string) (*sqlitedb, error) {
	sqliteDb, err := gorm.Open(sqlite.Open(pathname), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		return nil, err
	}
	db := &sqlitedb{
		handle:     sqliteDb,
		dbFilename: pathname,
	}
	return db, nil
}
