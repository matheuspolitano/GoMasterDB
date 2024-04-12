package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver  = "postgres"
	conString = "postgresql://root:MP@TEST123@localhost:5432/fintech_db?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, conString)
	if err != nil {
		log.Fatal(err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
