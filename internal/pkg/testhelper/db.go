package testhelper

import (
	"database/sql"
	"fmt"

	// database
	_ "github.com/lib/pq"
)

const (
	testUser = "test_pp"
	testPass = "test_pass"
	testDB   = "test_db"
	testPort = 5433
)

// OpenDB open and return a database connection
func OpenDB() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=0.0.0.0 port=%d user=%s password=%s dbname=%s sslmode=disable",
		testPort, testUser, testPass, testDB,
	)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
