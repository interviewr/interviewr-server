package http

import (
	t "interviewr-server/transport"
	"interviewr-server/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpVacancyHandler struct {
	VacancyUsecase usecases.VacancyUsecase
}

func NewVacancyHttpHandler(router *mux.Router, usecase usecases.VacancyUsecase) {
	handler := &HttpVacancyHandler{
		VacancyUsecase: usecase,
	}

	router.HandleFunc("/vacancies", handler.Fetch).Methods("GET")
}

func (h *HttpVacancyHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	vacancies, err := h.VacancyUsecase.Fetch()
	if err != nil {
		t.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	t.RespondWithJSON(w, http.StatusOK, vacancies)
}
