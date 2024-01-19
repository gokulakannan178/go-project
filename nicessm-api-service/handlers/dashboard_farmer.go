package handlers

import (
	"encoding/json"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

func (h *Handler) DashboardFarmerCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var farmer *models.DashboardFarmerCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var Farmers []models.DashboardFarmerCountReport
	Farmers, err = h.Service.DashboardFarmerCount(ctx, farmer)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Farmers) > 0 {
		m["farmer"] = Farmers
	} else {
		res := make([]models.Farmer, 0)
		m["farmer"] = res
	}
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DayWiseFarmerDemandChart(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var filter *models.DashboardFarmerCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// DayWiseFarmerDemandChart
	data, err := h.Service.DayWiseFarmerDemandChart(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmer"] = data
	response.With200V2(w, "Success", m, platform)

}
