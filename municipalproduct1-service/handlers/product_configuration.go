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

// SaveProductConfiguration
func (h *Handler) SaveProductConfiguration(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	pc := new(models.ProductConfiguration)
	err := json.NewDecoder(r.Body).Decode(&pc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveProductConfiguration(ctx, pc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["productConfiguration"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleProductConfiguration : ""
func (h *Handler) GetSingleProductConfiguration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	// UniqueID := r.URL.Query().Get("id")

	// if UniqueID == "" {
	// 	response.With400V2(w, "id is missing", platform)
	// 	return
	// }

	property := new(models.RefProductConfiguration)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	property, err := h.Service.GetSingleProductConfiguration(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["productConfiguration"] = property
	response.With200V2(w, "Success", m, platform)
}

// GetProductLogo : ""
func (h *Handler) GetProductLogo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	// UniqueID := r.URL.Query().Get("id")

	// if UniqueID == "" {
	// 	response.With400V2(w, "id is missing", platform)
	// 	return
	// }

	property := new(models.Logo)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	property, err := h.Service.GetProductLogo(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["logo"] = property
	response.With200V2(w, "Success", m, platform)
}

// GetWatermarkLogo : ""
func (h *Handler) GetWatermarkLogo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	// UniqueID := r.URL.Query().Get("id")

	// if UniqueID == "" {
	// 	response.With400V2(w, "id is missing", platform)
	// 	return
	// }

	property := new(models.WatermarkLogo)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	property, err := h.Service.GetWatermarkLogo(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["watermarkLogo"] = property
	response.With200V2(w, "Success", m, platform)
}

// FilterProductConfiguration : ""
func (h *Handler) FilterProductConfiguration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ProductConfigurationFilter
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
		if pagination.Limit = 1; Limit != "" {
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

	var pds []models.RefProductConfiguration
	log.Println(pagination)
	pds, err = h.Service.FilterProductConfiguration(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(pds) > 0 {
		m["productConfiguration"] = pds
	} else {
		res := make([]models.RefProductConfiguration, 0)
		m["productConfiguration"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
