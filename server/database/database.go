package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./database/database.sqlite")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db.Exec("CREATE TABLE IF NOT EXISTS exchange(code TEXT, codein TEXT, name TEXT, high TEXT, low TEXT, var_bid TEXT, pct_change TEXT, bid TEXT, ask TEXT, timestamp TEXT, create_date TEXT)")

	return db
}
