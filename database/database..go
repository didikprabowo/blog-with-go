package database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func MySQL() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "DIDIKprabowo_1995"
	dbName := "blog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
