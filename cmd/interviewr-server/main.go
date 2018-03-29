package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
	repo "interviewr-server/repository/postgres"
	usecases "interviewr-server/usecases"
	httpDeliver "interviewr-server/transport/http"
	cfg "interviewr-server/config/env"
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
	connectionStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	dbConnection, err := sql.Open("postgres", connectionStr)
	if err != nil && config.GetBool("debug") {
		fmt.Println(err)
	}
	defer dbConnection.Close()

	// err = dbConnection.Ping()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	r := mux.NewRouter()

	orgRepo := repo.NewOrganizationRepository(dbConnection)
	orgUsecase := usecases.NewOrganizationUsecase(orgRepo)
	httpDeliver.NewOrganizationHttpHandler(r, orgUsecase)

	http.ListenAndServe(config.GetString("address"), r)
}
