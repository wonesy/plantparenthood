package db

import (
	"database/sql"
	"fmt"

	// database
	_ "github.com/lib/pq"
)

// Open open and return a database connection
func Open(user, password, dbname string) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=0.0.0.0 port=5432 user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
