// package postgres

// import (
// 	"os"
// 	"src/pkg/dbs/postgres"
// 	"testing"
// )

// const connURL = "postgresql://postgres:admin@localhost:5432/tests"

// var testDB *postgres.Postgres
// var ids map[string]int64

// func TestMain(m *testing.M) {
// 	testDB = NewTestStorage()

// 	code := m.Run()
// 	DropTestStorage(testDB, ids)
// 	testDB.Close()

// 	os.Exit(code)
// }

// func NewTestStorage() *postgres.Postgres {
// 	conn, err := postgres.New(connURL)
// 	if err != nil {
// 		panic(err)
// 	}

// 	ids = map[string]int64{}
// 	ids["companyID"] = initTestCompanyStorage(NewCompanyStorage(conn))
// 	ids["employeeID"] = initTestEmployeeStorage(NewEmployeeStorage(conn))
// 	ids["infoCardID"] = initTestInfoCardStorage(NewInfoCardStorage(conn))
// 	ids["documentID"] = initTestDocumentStorage(NewDocumentStorage(conn))
// 	ids["checkpointID"] = initTestCheckpointStorage(NewCheckpointStorage(conn))
// 	ids["photoID"] = initTestPhotoMetaStorage(NewPhotoMetaStorage(conn))

// 	return conn
// }