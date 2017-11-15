package http

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/interviewr/interviewr-server/domain"
	"github.com/interviewr/interviewr-server/usecases"
)

type HttpOrganizationHandler struct {
	OrganizationUsecase OrganizationUsecase
}

func NewOrganizationHttpHandler(router *mux.Router, usecase OrganizationUsecase) {
	handler := &HttpOrganizationHandler{
		OrganizationUsecase: usecase,
	}

	router.HandleFunc("/orgs", handler.Fetch).Methods("GET")
	router.HandleFunc("/orgs", handler.Store).Methods("POST")
	router.HandleFunc("/orgs/{id}", handler.GetById).Methods("GET")
}

func (h *HttpOrganizationHandler) Fetch(w http.ResponseWriter, r *http.Request) error {
	orgs, error := h.OrganizationUsecase.Fetch()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, orgs)
}

func (h *HttpOrganizationHandler) GetById(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id, err := vars["id"]
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid organization ID")
		return
	}

	org, err := h.OrganizationUsecase.GetById(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Organization not found")
	}
	respondWithJSON(w, http.StatusOK, org)
}

func (h *HttpOrganizationHandler) Store(w http.ResponseWriter, r *http.Request) error {

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}