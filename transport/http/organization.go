package http

import (
	"encoding/json"
	"interviewr-server/domain"
	t "interviewr-server/transport"
	"interviewr-server/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpOrganizationHandler struct {
	OrganizationUsecase usecases.OrganizationUsecase
}

func NewOrganizationHttpHandler(router *mux.Router, usecase usecases.OrganizationUsecase) {
	handler := &HttpOrganizationHandler{
		OrganizationUsecase: usecase,
	}

	router.HandleFunc("/orgs", handler.Fetch).Methods("GET")
	router.HandleFunc("/orgs", handler.Store).Methods("POST")
	router.HandleFunc("/orgs/{id}", handler.GetById).Methods("GET")
	router.HandleFunc("/orgs/{id}/members", handler.GetMembers).Methods("GET")
	router.HandleFunc("/orgs/{id}/vacancies", handler.CreateVacancy).Methods("POST")
	router.HandleFunc("/orgs/{id}", handler.Delete).Methods("DELETE")
}

func (h *HttpOrganizationHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	orgs, err := h.OrganizationUsecase.Fetch()
	if err != nil {
		t.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	t.RespondWithJSON(w, http.StatusOK, orgs)
}

func (h *HttpOrganizationHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// id := vars["id"]
	// if !ok {
	// 	t.RespondWithError(w, http.StatusBadRequest, "Invalid organization ID")
	// 	return
	// }

	idP, err := strconv.Atoi(vars["id"])
	id := int64(idP)

	org, err := h.OrganizationUsecase.GetById(id)
	if err != nil {
		t.RespondWithError(w, http.StatusNotFound, "Organization not found")
	}

	t.RespondWithJSON(w, http.StatusOK, org)
}

func (h *HttpOrganizationHandler) Store(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var org *domain.Organization
	err := decoder.Decode(&org)
	if err != nil {
		t.RespondWithError(w, http.StatusBadRequest, "Couldn't decode organization data")
	}

	id, err := h.OrganizationUsecase.Store(org)
	if err != nil {
		t.RespondWithError(w, http.StatusInternalServerError, "Can't create organization")
	}

	t.RespondWithJSON(w, http.StatusCreated, id)
}

func (h *HttpOrganizationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idP, err := strconv.Atoi(vars["id"])
	id := int64(idP)

	// if !ok {
	// 	t.RespondWithError(w, http.StatusBadRequest, "Invalid organization ID")
	// 	return
	// }

	status, err := h.OrganizationUsecase.Delete(id)
	if err != nil {
		t.RespondWithError(w, http.StatusNotFound, "Can't delete organization")
	}

	t.RespondWithJSON(w, http.StatusNoContent, status)
}

func (h *HttpOrganizationHandler) GetMembers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idP, err := strconv.Atoi(vars["id"])
	id := int64(idP)
	// if !ok {
	// 	t.RespondWithError(w, http.StatusBadRequest, "Invalid organization ID")
	// 	return
	// }

	members, err := h.OrganizationUsecase.GetMembers(id)
	if err != nil {
		t.RespondWithError(w, http.StatusNotFound, "Can't get members of organization")
	}

	t.RespondWithJSON(w, http.StatusOK, members)
}

// TODO: fix me
func (h *HttpOrganizationHandler) CreateVacancy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idP, err := strconv.Atoi(vars["id"])
	id := int64(idP)
	// if !ok {
	// 	t.RespondWithError(w, http.StatusBadRequest, "Invalid organization ID")
	// 	return
	// }

	status, err := h.OrganizationUsecase.CreateVacancy(id)
	if err != nil {
		t.RespondWithError(w, http.StatusBadRequest, "Can't create vacancy for that organization")
	}

	t.RespondWithJSON(w, http.StatusCreated, status)
}
