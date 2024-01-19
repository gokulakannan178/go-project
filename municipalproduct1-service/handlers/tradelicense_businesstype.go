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

// SaveTradeLicenseBusinessType : ""
func (h *Handler) SaveTradeLicenseBusinessType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	businessType := new(models.TradeLicenseBusinessType)
	err := json.NewDecoder(r.Body).Decode(&businessType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveTradeLicenseBusinessType(ctx, businessType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseBusinessType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleTradeLicenseBusinessType : ""
func (h *Handler) GetSingleTradeLicenseBusinessType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	businessType := new(models.RefTradeLicenseBusinessType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	businessType, err := h.Service.GetSingleTradeLicenseBusinessType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseBusinessType"] = businessType
	response.With200V2(w, "Success", m, platform)
}

// UpdateTradeLicenseBusinessType : ""
func (h *Handler) UpdateTradeLicenseBusinessType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	businessType := new(models.TradeLicenseBusinessType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&businessType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if businessType.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateTradeLicenseBusinessType(ctx, businessType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseBusinessType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableTradeLicenseBusinessType : ""
func (h *Handler) EnableTradeLicenseBusinessType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableTradeLicenseBusinessType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseBusinessType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableTradeLicenseBusinessType : ""
func (h *Handler) DisableTradeLicenseBusinessType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableTradeLicenseBusinessType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseBusinessType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteTradeLicenseBusinessType : ""
func (h *Handler) DeleteTradeLicenseBusinessType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteTradeLicenseBusinessType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseBusinessType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterTradeLicenseBusinessType : ""
func (h *Handler) FilterTradeLicenseBusinessType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.TradeLicenseBusinessTypeFilter
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

	var TradeLicenseBusinessTypes []models.RefTradeLicenseBusinessType
	log.Println(pagination)
	TradeLicenseBusinessTypes, err = h.Service.FilterTradeLicenseBusinessType(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(TradeLicenseBusinessTypes) > 0 {
		m["tradeLicenseBusinessType"] = TradeLicenseBusinessTypes
	} else {
		res := make([]models.TradeLicenseBusinessType, 0)
		m["tradeLicenseBusinessType"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
