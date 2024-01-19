package handlers

import (
	"encoding/json"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

func (h *Handler) DashboardUserCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var user *models.DashboardUserCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var Users []models.DashboardUserCountReport
	Users, err = h.Service.DashboardUserCount(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Users) > 0 {
		m["user"] = Users
	} else {
		res := make([]models.User, 0)
		m["user"] = res
	}
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DayWiseUserDemandChart(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var filter *models.DashboardUserCountFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// DayWiseUserDemandChart
	data, err := h.Service.DayWiseUserDemandChart(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = data
	response.With200V2(w, "Success", m, platform)

}
