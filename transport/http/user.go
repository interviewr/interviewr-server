package http

import (
	t "interviewr-server/transport"
	"interviewr-server/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpUserHandler struct {
	UserUsecase usecases.UserUsecase
}

func NewUserHttpHandler(router *mux.Router, usecase usecases.UserUsecase) {
	handler := &HttpUserHandler{
		UserUsecase: usecase,
	}

	router.HandleFunc("/users", handler.Fetch).Methods("GET")
	router.HandleFunc("/users", handler.Store).Methods("POST")
	router.HandleFunc("/users/{id}", handler.GetById).Methods("GET")
}

func (h *HttpUserHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserUsecase.Fetch()
	if err != nil {
		t.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	t.RespondWithJSON(w, http.StatusOK, users)
}

func (h *HttpUserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		t.RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.UserUsecase.GetById(id)
	if err != nil {
		t.RespondWithError(w, http.StatusNotFound, "User not found")
	}

	t.RespondWithJSON(w, http.StatusOK, user)
}

func (h *HttpUserHandler) Store(w http.ResponseWriter, r *http.Request) {

}
