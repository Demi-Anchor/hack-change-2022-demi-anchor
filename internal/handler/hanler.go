package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
}

func New() *handler {
	return &handler{}
}

func (h *handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/readiness", h.healthCheck).Methods(http.MethodGet)

	return r
}

func (*handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
