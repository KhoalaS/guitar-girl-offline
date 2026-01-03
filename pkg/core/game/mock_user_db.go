package game

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewMockDb() *sql.DB {
	db, _ := sql.Open("sqlite3", "./data.db")
	return db
}
