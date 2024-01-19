package handlers

import (
	"encoding/json"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"

	"net/http"
	"strconv"
)

// SaveWaterBillDayWiseDashboard : ""
func (h *Handler) SaveWaterBillDayWiseDashboard(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	waterBill := new(models.WaterBillDashboardDayWise)
	err := json.NewDecoder(r.Body).Decode(&waterBill)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveWaterBillDayWiseDashboard(ctx, waterBill)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["waterBillDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleWaterBillDayWiseDashboard : ""
func (h *Handler) GetSingleWaterBillDayWiseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	waterBill := new(models.RefWaterBillDashboardDayWise)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	waterBill, err := h.Service.GetSingleWaterBillDayWiseDashboard(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["waterBillDashboardDayWise"] = waterBill
	response.With200V2(w, "Success", m, platform)
}

// UpdateWaterBillDayWiseDashboard : ""
func (h *Handler) UpdateWaterBillDayWiseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	waterBill := new(models.WaterBillDashboardDayWise)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&waterBill)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if waterBill.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateWaterBillDayWiseDashboard(ctx, waterBill)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["waterBillDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableWaterBillDayWiseDashboard : ""
func (h *Handler) EnableWaterBillDayWiseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableWaterBillDayWiseDashboard(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["waterBillDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableWaterBillDayWiseDashboard: ""
func (h *Handler) DisableWaterBillDayWiseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableWaterBillDayWiseDashboard(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["waterBillDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteWaterBillDayWiseDashboard : ""
func (h *Handler) DeleteWaterBillDayWiseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteWaterBillDayWiseDashboard(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["waterBillDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterWaterBillDayWiseDashboard : ""
func (h *Handler) FilterWaterBillDayWiseDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.WaterBillDashboardDayWiseFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	var pagination *models.Pagination
	if pageNo != "no" {
		pagination = new(models.Pagination)
		if pagination.PageNum = 1; pageNo != "" {
			page, err := strconv.Atoi(pageNo)
			if pagination.PageNum = 1; err == nil {
				pagination.PageNum = page
			}
		}
		if pagination.Limit = 10; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var waterBillDayWiseDashboardDayWises []models.RefWaterBillDashboardDayWise
	log.Println(pagination)
	waterBillDayWiseDashboardDayWises, err = h.Service.FilterWaterBillDayWiseDashboard(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(waterBillDayWiseDashboardDayWises) > 0 {
		m["waterBillDashboardDayWise"] = waterBillDayWiseDashboardDayWises
	} else {
		res := make([]models.WaterBillDashboardDayWise, 0)
		m["waterBillDashboardDayWise"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
