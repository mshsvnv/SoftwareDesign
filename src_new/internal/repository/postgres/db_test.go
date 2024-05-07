package mypostgres

import (
	"os"
	"src_new/pkg/storage/postgres"
	"testing"
)

const connURL = "postgresql://postgres:admin@localhost:5432/Shop"

var testDB *postgres.Postgres

func TestMain(m *testing.M) {

	testDB = NewTestStorage()

	code := m.Run()
	testDB.Close()

	os.Exit(code)
}

func NewTestStorage() *postgres.Postgres {

	conn, err := postgres.New(connURL)

	if err != nil {
		panic(err)
	}

	return conn
}
