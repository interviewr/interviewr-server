package organization

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHundler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/orgs/{id}").Handler(httptransport.NewServer(
		makeGetOrganizationEndpoint(s),
		decodeGetOrganizationRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/orgs").Handler(httptransport.NewServer(
		makePostOrganizationEndpoint(s),
		decodePostOrganizationRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodePostOrganizationRequest(_ context.Content, r *http.Request) (request interface{}, err error) {
	var req postOrganizationRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Profile); e != nil {
		return nil, e
	}

	return req, nil
}

func decodeGetOrganizationRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("bad route")
	}

	return getOrganizationRequest{ID: id}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeError(_ context.Content, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}