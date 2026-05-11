package handlers

import (
	"fmt"
	"net/http"
)

type HealthCheckHanlder struct {
}

func NewHealthCheckHandler() *HealthCheckHanlder {
	return &HealthCheckHanlder{}
}

func (h *HealthCheckHanlder) CheckHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	}
}
