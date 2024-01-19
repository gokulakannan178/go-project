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

// SaveUserChargeDashboard : ""
func (h *Handler) SaveUserChargeDashboard(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	userCharge := new(models.UserChargeDashboard)
	err := json.NewDecoder(r.Body).Decode(&userCharge)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveUserChargeDashboard(ctx, userCharge)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userChargeDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleUserChargeDashboard : ""
func (h *Handler) GetSingleUserChargeDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	userCharge := new(models.RefUserChargeDashboard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	userCharge, err := h.Service.GetSingleUserChargeDashboard(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userChargeDashboard"] = userCharge
	response.With200V2(w, "Success", m, platform)
}

// UpdateUserChargeDashboard : ""
func (h *Handler) UpdateUserChargeDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	userCharge := new(models.UserChargeDashboard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&userCharge)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if userCharge.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateUserChargeDashboard(ctx, userCharge)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userChargeDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableUserChargeDashboard : ""
func (h *Handler) EnableUserChargeDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableUserChargeDashboard(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userChargeDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableUserChargeDashboard : ""
func (h *Handler) DisableUserChargeDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableUserChargeDashboard(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userChargeDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteDashBoardUserCharge : ""
func (h *Handler) DeleteDashBoardUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDashBoardUserCharge(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userChargeDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterUserChargeDashboard : ""
func (h *Handler) FilterUserChargeDashboard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.UserChargeDashboardFilter
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

	var userChargeDashboards []models.RefUserChargeDashboard
	log.Println(pagination)
	userChargeDashboards, err = h.Service.FilterUserChargeDashboard(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(userChargeDashboards) > 0 {
		m["userChargeDashboard"] = userChargeDashboards
	} else {
		res := make([]models.UserChargeDashboard, 0)
		m["userChargeDashboard"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
