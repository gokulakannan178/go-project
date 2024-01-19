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

// SaveTradeLicenseDashboardDayWise : ""
func (h *Handler) SaveTradeLicenseDashboardDayWise(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	tradeLicense := new(models.TradeLicenseDashboardDayWise)
	err := json.NewDecoder(r.Body).Decode(&tradeLicense)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveTradeLicenseDashboardDayWise(ctx, tradeLicense)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleTradeLicenseDashboardDayWise : ""
func (h *Handler) GetSingleTradeLicenseDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	tradeLicense := new(models.RefTradeLicenseDashboardDayWise)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	tradeLicense, err := h.Service.GetSingleTradeLicenseDashboardDayWise(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboardDayWise"] = tradeLicense
	response.With200V2(w, "Success", m, platform)
}

// UpdateTradeLicenseDashboardDayWise : ""
func (h *Handler) UpdateTradeLicenseDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	tradeLicense := new(models.TradeLicenseDashboardDayWise)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&tradeLicense)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if tradeLicense.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateTradeLicenseDashboardDayWise(ctx, tradeLicense)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableTradeLicenseDashboardDayWise : ""
func (h *Handler) EnableTradeLicenseDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableTradeLicenseDashboardDayWise(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableTradeLicenseDashboardDayWise : ""
func (h *Handler) DisableTradeLicenseDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableTradeLicenseDashboardDayWise(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteDashBoardTradeLicense : ""
func (h *Handler) DeleteDashBoardTradeLicenseDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDashBoardTradeLicenseDayWise(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseDashboardDayWise"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterTradeLicenseDashboardDayWise : ""
func (h *Handler) FilterTradeLicenseDashboardDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.TradeLicenseDashboardDayWiseFilter
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

	var tradeLicenseDashboardDayWises []models.RefTradeLicenseDashboardDayWise
	log.Println(pagination)
	tradeLicenseDashboardDayWises, err = h.Service.FilterTradeLicenseDashboardDayWise(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(tradeLicenseDashboardDayWises) > 0 {
		m["tradeLicenseDashboardDayWise"] = tradeLicenseDashboardDayWises
	} else {
		res := make([]models.TradeLicenseDashboardDayWise, 0)
		m["tradeLicenseDashboardDayWise"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
