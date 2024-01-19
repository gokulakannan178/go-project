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

// SaveWaterTaxConnectionType : ""
func (h *Handler) SaveWaterTaxConnectionType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	watertaxconnectiontype := new(models.WaterTaxConnectionType)
	err := json.NewDecoder(r.Body).Decode(&watertaxconnectiontype)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveWaterTaxConnectionType(ctx, watertaxconnectiontype)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["watertaxconnectiontype"] = watertaxconnectiontype
	response.With200V2(w, "Success", m, platform)
}

// GetSingleWaterTaxConnectionType : ""
func (h *Handler) GetSingleWaterTaxConnectionType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	watertaxconnectiontype := new(models.RefWaterTaxConnectionType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	watertaxconnectiontype, err := h.Service.GetSingleWaterTaxConnectionType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["watertaxconnectiontype"] = watertaxconnectiontype
	response.With200V2(w, "Success", m, platform)
}

// UpdateWaterTaxConnectionType : ""
func (h *Handler) UpdateWaterTaxConnectionType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	watertaxconnectiontype := new(models.WaterTaxConnectionType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&watertaxconnectiontype)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if watertaxconnectiontype.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateWaterTaxConnectionType(ctx, watertaxconnectiontype)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["watertaxconnectiontype"] = watertaxconnectiontype
	response.With200V2(w, "Success", m, platform)
}

//EnableWaterTaxConnectionType : ""
func (h *Handler) EnableWaterTaxConnectionType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableWaterTaxConnectionType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["watertaxconnectiontype"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableWaterTaxConnectionType : ""
func (h *Handler) DisableWaterTaxConnectionType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableWaterTaxConnectionType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["watertaxconnectiontype"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteWaterTaxConnectionType : ""
func (h *Handler) DeleteWaterTaxConnectionType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteWaterTaxConnectionType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["watertaxconnectiontype"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterWaterTaxConnectionType : ""
func (h *Handler) FilterWaterTaxConnectionType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.WaterTaxConnectionTypeFilter
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

	var watertaxconnectiontypes []models.RefWaterTaxConnectionType
	log.Println(pagination)
	watertaxconnectiontypes, err = h.Service.FilterWaterTaxConnectionType(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(watertaxconnectiontypes) > 0 {
		m["watertaxconnectiontype"] = watertaxconnectiontypes
	} else {
		res := make([]models.WaterTaxConnectionType, 0)
		m["watertaxconnectiontype"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
