package handlers

import (
	"ecommerce-service/app"
	"ecommerce-service/models"
	"ecommerce-service/response"
	"encoding/json"
	"net/http"
)

// SaveRegisterProduct : ""
func (h *Handler) SaveRegisterProduct(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	block := new(models.RegisterProduct)
	err := json.NewDecoder(r.Body).Decode(&block)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveRegisterProduct(ctx, block)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["product"] = block
	response.With200V2(w, "Success", m, platform)
}
