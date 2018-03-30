package main

import (
	"fmt"
	cfg "interviewr-server/config/env"
	"interviewr-server/repository"
	httpDeliver "interviewr-server/transport/http"
	"interviewr-server/usecases"
	"net/http"

	"github.com/gorilla/mux"
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

	dbmap := repository.InitStorage("postgres", dsn)
	defer repository.DropAndClose(dbmap)

	// if err != nil && config.GetBool("debug") {
	// 	fmt.Println(err)
	// }

	r := mux.NewRouter()

	orgRepo := repository.NewOrganizationRepository(dbmap)
	orgUsecase := usecases.NewOrganizationUsecase(orgRepo)
	httpDeliver.NewOrganizationHttpHandler(r, orgUsecase)

	http.ListenAndServe(config.GetString("address"), r)
}
