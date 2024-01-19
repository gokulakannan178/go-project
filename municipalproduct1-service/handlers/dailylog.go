package handlers

import (
	"municipalproduct1-service/response"
	"net/http"
)

//TodaysLog : ""
func (h *Handler) TodaysLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	err := h.Service.TodaysLog()
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	response.With200V2(w, "Success", m, platform)
}
