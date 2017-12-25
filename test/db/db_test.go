package db_test

import (
	"testing"

	"database/sql"
)

func TestDbConnection(t *testing.T) {
	db, err := sql.Open("txdb", "db connection check")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
}
