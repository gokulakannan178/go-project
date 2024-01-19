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

// SaveTradeLicenseRateMaster : ""
func (h *Handler) SaveTradeLicenseRateMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	rateLicense := new(models.TradeLicenseRateMaster)
	err := json.NewDecoder(r.Body).Decode(&rateLicense)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveTradeLicenseRateMaster(ctx, rateLicense)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseRateMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleTradeLicenseRateMaster : ""
func (h *Handler) GetSingleTradeLicenseRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	rateLicense := new(models.RefTradeLicenseRateMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	rateLicense, err := h.Service.GetSingleTradeLicenseRateMaster(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseRateMaster"] = rateLicense
	response.With200V2(w, "Success", m, platform)
}

// UpdateTradeLicenseRateMaster : ""
func (h *Handler) UpdateTradeLicenseRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	rateLicense := new(models.TradeLicenseRateMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&rateLicense)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if rateLicense.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateTradeLicenseRateMaster(ctx, rateLicense)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseRateMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableTradeLicenseRateMaster : ""
func (h *Handler) EnableTradeLicenseRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableTradeLicenseRateMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseRateMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableTradeLicenseRateMaster : ""
func (h *Handler) DisableTradeLicenseRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableTradeLicenseRateMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseRateMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteTradeLicenseRateMaster : ""
func (h *Handler) DeleteTradeLicenseRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteTradeLicenseRateMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseRateMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterTradeLicenseRateMaster : ""
func (h *Handler) FilterTradeLicenseRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.TradeLicenseRateMasterFilter
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

	var TradeLicenseRateMasters []models.RefTradeLicenseRateMaster
	log.Println(pagination)
	TradeLicenseRateMasters, err = h.Service.FilterTradeLicenseRateMaster(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(TradeLicenseRateMasters) > 0 {
		m["tradeLicenseRateMaster"] = TradeLicenseRateMasters
	} else {
		res := make([]models.TradeLicenseRateMaster, 0)
		m["tradeLicenseRateMaster"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
