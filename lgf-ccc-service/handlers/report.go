package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"net/http"
)

func (h *Handler) DayWiseDumphistoryCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")

	filter := new(models.FilterDumpHistory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var report []models.MonthWiseDumphistoryCount
	//report := new(models.MonthWiseDumphistoryCount)
	report, err = h.Service.DayWiseDumphistoryCount(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dateRangeWise"] = report
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) MonthWiseDumphistoryCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")

	filter := new(models.FilterDumpHistory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var report []models.MonthWiseDumphistoryCount
	//report := new(models.MonthWiseDumphistoryCount)
	report, err = h.Service.MonthWiseDumphistoryCount(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dateRangeWise"] = report
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) CircleWiseHouseVisitedCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")

	filter := new(models.FilterHouseVisited)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var report []models.CircleWiseHouseVisitedv2
	//report := new(models.MonthWiseDumphistoryCount)
	report, err = h.Service.CircleWiseHouseVisitedCount(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["circlewisehousevisitd"] = report
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DayWiseWardHouseVisitedCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")

	filter := new(models.FilterHouseVisited)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var report []models.CircleWiseHouseVisitedv2
	//report := new(models.MonthWiseDumphistoryCount)
	report, err = h.Service.DayWiseWardHouseVisitedCount(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["daywisehousevisted"] = report
	response.With200V2(w, "Success", m, platform)
}
