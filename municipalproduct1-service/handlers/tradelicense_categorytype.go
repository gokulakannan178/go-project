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

// SaveTradeLicenseCategoryType : ""
func (h *Handler) SaveTradeLicenseCategoryType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	categoryType := new(models.TradeLicenseCategoryType)
	err := json.NewDecoder(r.Body).Decode(&categoryType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveTradeLicenseCategoryType(ctx, categoryType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseCategoryType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleTradeLicenseCategoryType : ""
func (h *Handler) GetSingleTradeLicenseCategoryType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	categoryType := new(models.RefTradeLicenseCategoryType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	categoryType, err := h.Service.GetSingleTradeLicenseCategoryType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseCategoryType"] = categoryType
	response.With200V2(w, "Success", m, platform)
}

// UpdateTradeLicenseCategoryType : ""
func (h *Handler) UpdateTradeLicenseCategoryType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	categoryType := new(models.TradeLicenseCategoryType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&categoryType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if categoryType.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateTradeLicenseCategoryType(ctx, categoryType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseCategoryType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableTradeLicenseCategoryType : ""
func (h *Handler) EnableTradeLicenseCategoryType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableTradeLicenseCategoryType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseCategoryType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableTradeLicenseCategoryType : ""
func (h *Handler) DisableTradeLicenseCategoryType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableTradeLicenseCategoryType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseCategoryType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteTradeLicenseCategoryType : ""
func (h *Handler) DeleteTradeLicenseCategoryType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteTradeLicenseCategoryType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tradeLicenseCategoryType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterTradeLicenseCategoryType : ""
func (h *Handler) FilterTradeLicenseCategoryType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.TradeLicenseCategoryTypeFilter
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

	var TradeLicenseCategoryTypes []models.RefTradeLicenseCategoryType
	log.Println(pagination)
	TradeLicenseCategoryTypes, err = h.Service.FilterTradeLicenseCategoryType(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(TradeLicenseCategoryTypes) > 0 {
		m["tradeLicenseCategoryType"] = TradeLicenseCategoryTypes
	} else {
		res := make([]models.TradeLicenseCategoryType, 0)
		m["tradeLicenseCategoryType"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
