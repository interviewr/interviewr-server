package main

import (
	"database/sql"
	"fmt"
	cfg "interviewr-server/config/env"
	migrations "interviewr-server/migrations"
	"interviewr-server/repository"
	httpDeliver "interviewr-server/transport/http"
	"interviewr-server/usecases"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	gorp "gopkg.in/gorp.v1"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()

	if config.GetBool("debug") {
		fmt.Println("Service run on debug mode")
	}
}

func main() {
	dbHost := config.GetString("database.host")
	dbPort := config.GetString("database.port")
	dbUser := config.GetString("database.user")
	dbPass := config.GetString("database.pass")
	dbName := config.GetString("database.name")
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil && config.GetBool("debug") {
		fmt.Println(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	defer dbmap.Db.Close()

	migrations := &migrate.AssetMigrationSource{
		Asset:    migrations.Asset,
		AssetDir: migrations.AssetDir,
		Dir:      "migrations",
	}

	n, err := migrate.Exec(dbmap.Db, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Applied %d migrations!\n", n)

	r := mux.NewRouter()

	orgRepo := repository.NewPostgresOrganizationRepository(dbmap)
	orgUsecase := usecases.NewOrganizationUsecase(orgRepo)
	httpDeliver.NewOrganizationHttpHandler(r, orgUsecase)

	http.ListenAndServe(config.GetString("address"), r)
}
