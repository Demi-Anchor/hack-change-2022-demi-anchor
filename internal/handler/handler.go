package handler

import (
	"demi-anchor/internal/models"
	"demi-anchor/pkg/errtrace"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Service interface {
	ValidateDonation(d *models.Donation) (bool, string)
	ValidatePeriod(p *models.Period) (bool, string)

	AddDonation(d *models.Donation) error

	GetDailyDonations(p *models.Period) ([]models.DailyDonation, error)
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
	r.HandleFunc("/api/v1/donations", h.addDonation).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/donations/daily", h.getDailyDonations).Methods(http.MethodPost)

	return r
}

func (*handler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

func (h *handler) addDonation(w http.ResponseWriter, r *http.Request) {
	var d *models.Donation
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		log.Err(errtrace.AddTrace(err)).Send()
		sendMsgResp(w, http.StatusBadRequest, err.Error())
		return
	}

	if ok, msg := h.service.ValidateDonation(d); !ok {
		sendMsgResp(w, http.StatusBadRequest, msg)
		return
	}

	if err := h.service.AddDonation(d); err != nil {
		sendMsgResp(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (h *handler) getDailyDonations(w http.ResponseWriter, r *http.Request) {
	var p *models.Period
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Err(errtrace.AddTrace(err)).Send()
		sendMsgResp(w, http.StatusBadRequest, err.Error())
		return
	}

	if ok, msg := h.service.ValidatePeriod(p); !ok {
		sendMsgResp(w, http.StatusBadRequest, msg)
		return
	}

	d, err := h.service.GetDailyDonations(p)
	if err != nil {
		sendMsgResp(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResp(w, http.StatusOK, d)
	return
}

func sendResp(w http.ResponseWriter, statusCode int, dataStruct any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(&dataStruct); err != nil {
		log.Err(errtrace.AddTrace(err)).Send()
	}
}

func sendMsgResp(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	msgResp := &struct {
		Msg string
	}{msg}

	if err := json.NewEncoder(w).Encode(&msgResp); err != nil {
		log.Err(errtrace.AddTrace(err)).Send()
	}
}
