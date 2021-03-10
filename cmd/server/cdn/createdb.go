package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)


func createDb() {
	stdin := bufio.NewReader(os.Stdin)
	dbf := InstallDir + string(os.PathSeparator) + DbFilename
	if _, err := os.Stat(dbf); err == nil {
		fmt.Printf("Found an old database %s, are you sure you want to remove it? (y|n):\n", dbf)
		answer, _ := stdin.ReadString('\n')
		switch answer {
		case "y\n":
			os.Remove(dbf)
		case "n\n":
			return
		default:
			return
		}
	}

	dbfile, err := os.Create(dbf)
	if err != nil { panic(err) }
	dbfile.Close()

	db, erropen := sql.Open("sqlite3", dbf)
	if erropen != nil { panic(erropen) }

	createContentTable(db)
}

func execSql(db *sql.DB, sql string) {
	log.Printf("\tExecuting:\n%s", sql)
	statement, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}

func createContentTable(db *sql.DB) {
	sql := `CREATE TABLE content (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	);`
	execSql(db, sql)
}
