package db

import (
	"database/sql"
	"os"

	_"github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	connstr := os.Getenv("CONN_STR")
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}
	return db, err
}