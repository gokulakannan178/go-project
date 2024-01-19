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

// SaveMobileTowerDayWise : ""
func (h *Handler) SaveMobileTowerDayWise(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	mobileTower := new(models.MobiletowerDashboardDayWise)
	err := json.NewDecoder(r.Body).Decode(&mobileTower)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveMobileTowerDayWise(ctx, mobileTower)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDayWiseDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleMobileTower : ""
func (h *Handler) GetSingleMobileTowerDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	mobileTower := new(models.RefMobileTowerDayWiseDashboard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	mobileTower, err := h.Service.GetSingleMobileTowerDayWise(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDayWiseDashboard"] = mobileTower
	response.With200V2(w, "Success", m, platform)
}

// UpdateMobileTowerDayWise : ""
func (h *Handler) UpdateMobileTowerDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	mobileTower := new(models.MobiletowerDashboardDayWise)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&mobileTower)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if mobileTower.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateMobileTowerDayWise(ctx, mobileTower)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDayWiseDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableMobileTowerDayWise : ""
func (h *Handler) EnableMobileTowerDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableMobileTowerDayWise(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDayWiseDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableMobileTowerDayWise : ""
func (h *Handler) DisableMobileTowerDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableMobileTowerDayWise(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDayWiseDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteMobileTowerDayWise : ""
func (h *Handler) DeleteMobileTowerDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteMobileTowerDayWise(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerDayWiseDashboard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterMobileTowerDayWise : ""
func (h *Handler) FilterMobileTowerDayWise(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.MobileTowerDashboardDayWiseFilter
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

	var leaseDashboards []models.RefMobileTowerDayWiseDashboard
	log.Println(pagination)
	leaseDashboards, err = h.Service.FilterMobileTowerDayWise(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(leaseDashboards) > 0 {
		m["mobileTowerDayWiseDashboard"] = leaseDashboards
	} else {
		res := make([]models.MobiletowerDashboardDayWise, 0)
		m["mobileTowerDayWiseDashboard"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
