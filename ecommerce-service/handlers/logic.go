package handlers

import (
	"ecommerce-service/app"
	"ecommerce-service/models"
	"ecommerce-service/response"
	"encoding/json"
	"net/http"
)

//Logic : ""
func (h *Handler) Logic(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var feature []models.VarientInputLogic

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&feature)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	out, err := h.Service.Logic(ctx, feature)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["feature"] = out
	response.With200V2(w, "Success", m, platform)
}
