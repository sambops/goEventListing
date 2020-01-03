package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Config() *sql.DB {
	dbConn, _ := sql.Open("postgres", "postgres://postgres:1234@localhost/pscs?sslmode=disable")
	return dbConn
}
