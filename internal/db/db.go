package db

import (
	"gorm.io/gorm"
)

type PrimaryKey uint64

type Db interface {
	Handle() *gorm.DB
	Backup()
}

type sqlitedb struct {
	handle *gorm.DB
	dbFilename string
}

func NewDb(handle *gorm.DB, dbFilename string) sqlitedb {
	return sqlitedb{
		handle,
		dbFilename,
	}
}

func (s sqlitedb) Handle() *gorm.DB {
	return s.handle
}

func (s sqlitedb) Backup() {
	s.handle.Raw("BEGIN EXCLUSIVE")
	s.handle.Raw("COMMIT")
}
