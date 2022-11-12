package handler

import (
	"demi-anchor/internal/models"
	"demi-anchor/pkg/errtrace"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/gorilla/mux"
)

type Service interface {
	ValidateUser(u models.User) (bool, string)
	CreateUser(u models.User) error
	ValidateDonation(d Donations) (bool, string)
	CreateDonations(d Donations) ([]byte, error)
}

type handler struct {
	service Service
}

func New(s Service) *handler {
	return &handler{service: s}
}

func (h *handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/readiness", h.healthCheck).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user", h.createUser).Methods(http.MethodPost)
	r.HandleFunc("/donations_info", h.donationsCheck).Methods(http.MethodPost)

	return r
}

func (*handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

func (h *handler) createUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Err(errtrace.AddTrace(err)).Send()
		sendErrResp(w, http.StatusBadRequest, err.Error())
		return
	}

	if ok, msg := h.service.ValidateUser(u); !ok {
		sendErrResp(w, http.StatusBadRequest, msg)
		return
	}

	if err := h.service.CreateUser(u); err != nil {
		log.Err(err).Send()
		sendErrResp(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (h *handler) donationsCheck(w http.ResponseWriter, r *http.Request) {
	var d Donations
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		log.Err(errtrace.AddTrace(err)).Send()
		sendErrResp(w, http.StatusBadRequest, err.Error())
		return
	}

	if ok, msg := h.service.ValidateDonation(d); !ok {
		sendErrResp(w, http.StatusBadRequest, msg)
		return
	}

	data, err := h.service.CreateDonations(d)
	if err != nil {
		log.Err(err).Send()
		sendErrResp(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Write(data)
	return
}
