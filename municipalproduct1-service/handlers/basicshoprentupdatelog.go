package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

// BasicShopRentUpdateGetPaymentsToBeUpdated : ""
func (h *Handler) BasicShopRentUpdateGetPaymentsToBeUpdated(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	rbsrul := new(models.RefBasicShopRentUpdateLog)
	err := json.NewDecoder(r.Body).Decode(&rbsrul)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	payments, err := h.Service.BasicShopRentUpdateGetPaymentsToBeUpdated(ctx, rbsrul)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payments"] = payments
	if len(payments) < 1 {
		m["payments"] = []interface{}{}

	}
	response.With200V2(w, "Success", m, platform)
}
