package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"strconv"

	"net/http"
)

// PmDashboard : ""
func (h *Handler) PmDashboard(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	appVersion := r.URL.Query().Get("appVersion")
	if appVersion == "" {
		response.With400V2(w, "appVersion is missing", platform)
		return
	}
	key, err1 := strconv.ParseFloat(appVersion, 32)
	if err1 != nil {
		return
	}
	filter := new(models.PmDashboardFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	productConfig, err := h.Service.GetSingleDefaultProductConfiguration(ctx)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	pmDashboardData, err := h.Service.PmDashboard(ctx, filter)

	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["pmDashboardData"] = pmDashboardData
	if productConfig.AppVersion > key {
		m["pmDashboardData"] = nil
	}
	response.With200V2(w, "Success", m, platform)
}
