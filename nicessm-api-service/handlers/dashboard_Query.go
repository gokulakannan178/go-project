package handlers

import (
	"encoding/json"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

func (h *Handler) DashboardQueryCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var query *models.DashboardQueryCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&query)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var Querys []models.DashboardQueryCountReport
	Querys, err = h.Service.DashboardQueryCount(ctx, query)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Querys) > 0 {
		m["query"] = Querys
	} else {
		res := make([]models.Query, 0)
		m["query"] = res
	}
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DayWiseQueryDemandChart(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var filter *models.DashboardQueryCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// DayWiseQueryDemandChart
	data, err := h.Service.DayWiseQueryDemandChart(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["query"] = data
	response.With200V2(w, "Success", m, platform)

}
