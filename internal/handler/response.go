package handler

import (
	"demi-anchor/pkg/errtrace"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type errResp struct {
	Msg string `json:"msg"`
}

func sendErrResp(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	errResp := &errResp{Msg: msg}
	if err := json.NewEncoder(w).Encode(&errResp); err != nil {
		log.Err(errtrace.AddTrace(err)).Send()
	}
}
