package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"net/http"
)

//SaveUserLocation : ""
func (h *Handler) SaveUserLocation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	userLocation := new(models.UserLocation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&userLocation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveUserLocation(ctx, userLocation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userLocation"] = userLocation
	response.With200V2(w, "Success", m, platform)
}
