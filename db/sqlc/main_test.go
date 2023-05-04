package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5434/simple_bank?sslmode=disable"
)

var testQueries *Queries
var TestDb *sql.DB
var err error

func TestMain(m *testing.M) {
	TestDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("error occured whicle connecting to db", err)
	}
	testQueries = New(TestDb)
	os.Exit(m.Run())
}
