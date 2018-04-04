package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	gorp "gopkg.in/gorp.v1"
)

func InitStorage(storage string, dsn string) (*gorp.DbMap, error) {
	dbmap := createStorage(storage, dsn)

	migrations := &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "migrations",
	}

	// TODO: do not hardcore driver name!
	n, err := migrate.Exec(dbmap.Db, "postgres", migrations, migrate.Up)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Applied %d migrations!\n", n)

	return dbmap, nil
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
	dbmap.Db.Close()
}

var (
	INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
	NOT_FOUND_ERROR       = errors.New("Requested Item not found")
	CONFLICT_ERROR        = errors.New("Item already exist")
)
