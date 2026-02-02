package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func Open() *sql.DB {
	db, err := sql.Open("sqlite", "R:\\projects\\rust\\fipe_rs\\target\\debug\\fipe_rs.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
