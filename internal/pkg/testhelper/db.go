package testhelper

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/DATA-DOG/go-sqlmock"

	// database
	_ "github.com/lib/pq"
)

const (
	testUser = "test_pp"
	testPass = "test_pass"
	testDB   = "test_db"
	testPort = 5433

	e2eEnv = "PP_E2E_TEST"
)

// IsEnd2EndTest returns whether the tests are being run in e2e mode
func IsEnd2EndTest() bool {
	return os.Getenv(e2eEnv) != ""
}

// OpenDB open and return a database connection
func OpenDB() (*sql.DB, sqlmock.Sqlmock, error) {

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	if !IsEnd2EndTest() {
		return mockDB, mock, err
	}

	psqlconn := fmt.Sprintf("host=0.0.0.0 port=%d user=%s password=%s dbname=%s sslmode=disable",
		testPort, testUser, testPass, testDB,
	)

	realDB, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, nil, err
	}

	if err := realDB.Ping(); err != nil {
		return nil, nil, err
	}

	return realDB, mock, nil
}

// ExpectationsWereMet branches on whether e2e tests are set
func ExpectationsWereMet(mock sqlmock.Sqlmock) error {
	if !IsEnd2EndTest() {
		return mock.ExpectationsWereMet()
	}
	return nil
}
