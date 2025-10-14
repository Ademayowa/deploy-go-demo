package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite", "property.db")
	if err != nil {
		panic("could not connect to database")

	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createPropertiesTable := `
	CREATE TABLE IF NOT EXISTS properties (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		location TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createPropertiesTable)
	if err != nil {
		panic("could not create properties table" + err.Error())
	}
}
