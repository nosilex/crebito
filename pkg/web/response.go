package web

import (
	"encoding/json"
	"net/http"
)

var Response = response{}

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (r response) OK(w http.ResponseWriter, v any) {
	r.output(w, v, http.StatusOK)
}

func (r response) Error(w http.ResponseWriter, e error, c int) {
	r.Code = c
	r.Message = e.Error()
	r.output(w, r, c)
}

func (r response) output(w http.ResponseWriter, v any, c int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(c)
	_ = json.NewEncoder(w).Encode(v)
}
