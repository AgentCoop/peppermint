package pkg

import "gorm.io/gorm"

type Db interface {
	Handle() *gorm.DB
	Backup()
}

type Dbal interface {
	Db() Db
}
