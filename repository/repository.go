package repository

import (
	"database/sql"
	"errors"
	"interviewr-server/domain"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

func InitStorage(storage string, dsn string) *gorp.DbMap {
	dbmap := createStorage(storage, dsn)
	table := dbmap.AddTableWithName(domain.Organization{}, "organization").SetKeys(true, "ID")
	table.ColMap("ID").Rename("id")
	table.ColMap("Name").Rename("name")
	table.ColMap("Email").Rename("email")
	table.ColMap("Description").Rename("description")
	table.ColMap("Location").Rename("location")

	err := dbmap.DropTablesIfExists()
	if err != nil {
		panic(err)
	}

	err = dbmap.CreateTables()
	if err != nil {
		panic(err)
	}

	err = dbmap.CreateIndex()
	if err != nil {
		panic(err)
	}

	return dbmap
}

func createStorage(storage string, dsn string) *gorp.DbMap {
	dialect, driver := dialectAndDriver(storage)
	dbmap := &gorp.DbMap{Db: connect(driver, dsn), Dialect: dialect}
	return dbmap
}

func connect(driver string, dsn string) *sql.DB {
	if dsn == "" {
		panic("DSN is no set.")
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic("Error connecting to db: " + err.Error())
	}
	return db
}

func dialectAndDriver(storage string) (gorp.Dialect, string) {
	switch storage {
	case "postgres":
		return gorp.PostgresDialect{}, "postgres"
	case "sqlite":
		return gorp.SqliteDialect{}, "sqlite3"
	default:
		return gorp.PostgresDialect{}, "postgres"
	}
}

func DropAndClose(dbmap *gorp.DbMap) {
	dbmap.DropTablesIfExists()
	dbmap.Db.Close()
}

var (
	INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
	NOT_FOUND_ERROR       = errors.New("Requested Item not found")
	CONFLICT_ERROR        = errors.New("Item already exist")
)
