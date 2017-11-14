package main

import (
	"flag"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"os/signal"
	"syscall"

	"github.com/interviewr/interviewr-server/repository/postgres"
	"github.com/interviewr/interviewr-server/usecases"
	"github.com/interviewr/interviewr-server/transport/http"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8090", "HTTP listen address")
	)

	flag.Parse()

	r := mux.NewRouter()

	orgRepo := NewOrganizationRepository(/* pass DB connection here */)
	orgUsecase := NewOrganizationUsecase(orgRepo)
	NewOrganizationHttpHandler(r) // TODO: not sure about passing r into HttpHandler func

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, r)
	}()

	logger.Log("exit", <-errs)
}