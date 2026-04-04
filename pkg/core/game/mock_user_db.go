package game

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewMockDb() *sql.DB {
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		panic("could not open game database")
	}
	return db
}
