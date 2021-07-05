package db

import (
	"gorm.io/gorm"
)

type PrimaryKey uint64

type sqlitedb struct {
	handle     *gorm.DB
	dbFilename string
}

//type Dbal struct {
//	db *gorm.DB
//}
//
//func (dbal Dbal) Db() *gorm.DB {
//	return dbal.db
//}

func (s sqlitedb) Handle() *gorm.DB {
	return s.handle
}

func (s sqlitedb) Backup() {
	s.handle.Raw("BEGIN EXCLUSIVE")
	s.handle.Raw("COMMIT")
}
