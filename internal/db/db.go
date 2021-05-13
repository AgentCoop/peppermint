package db

type SqlTransaction interface {
	Start()
	Rollback()
	Commit()
	AddQuery()
}

type SqlQuery interface {
	
}