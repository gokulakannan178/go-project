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

// SaveOverallPropertyDemand : ""
func (h *Handler) SaveOverallPropertyDemand(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	mtrt := new(models.OverallPropertyDemand)
	err := json.NewDecoder(r.Body).Decode(&mtrt)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveOverallPropertyDemand(ctx, mtrt)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["overallpropertydemands"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleOverallPropertyDemand : ""
func (h *Handler) GetSingleOverallPropertyDemand(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	mtrt := new(models.RefOverallPropertyDemand)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	mtrt, err := h.Service.GetSingleOverallPropertyDemand(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["overallpropertydemands"] = mtrt
	response.With200V2(w, "Success", m, platform)
}

// UpdateOverallPropertyDemand : ""
func (h *Handler) UpdateOverallPropertyDemand(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	mtrt := new(models.OverallPropertyDemand)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&mtrt)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if mtrt.PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateOverallPropertyDemand(ctx, mtrt)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["overallpropertydemands"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableOverallPropertyDemand : ""
func (h *Handler) EnableOverallPropertyDemand(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	PropertyID := r.URL.Query().Get("id")

	if PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableOverallPropertyDemand(ctx, PropertyID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["overallpropertydemands"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableOverallPropertyDemand : ""
func (h *Handler) DisableOverallPropertyDemand(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	PropertyID := r.URL.Query().Get("id")

	if PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableOverallPropertyDemand(ctx, PropertyID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["overallpropertydemands"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteOverallPropertyDemand : ""
func (h *Handler) DeleteOverallPropertyDemand(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	PropertyID := r.URL.Query().Get("id")

	if PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteOverallPropertyDemand(ctx, PropertyID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["overallpropertydemands"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterOverallPropertyDemand : ""
func (h *Handler) FilterOverallPropertyDemand(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.OverallPropertyDemandFilter
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

	var overallPropertyDemands []models.RefOverallPropertyDemand
	log.Println(pagination)
	overallPropertyDemands, err = h.Service.FilterOverallPropertyDemand(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(overallPropertyDemands) > 0 {
		m["overallpropertydemands"] = overallPropertyDemands
	} else {
		res := make([]models.OverallPropertyDemand, 0)
		m["overallpropertydemands"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// UpdateOverAllPropertyDemandPropertyID : ""
func (h *Handler) UpdateOverAllPropertyDemandPropertyID(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.UpdatePropertyUniqueID)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if len(property.UniqueIDs) == 0 {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateOverAllPropertyDemandPropertyID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["updateOverAllPropertyDemandPropertyID"] = "success"
	response.With200V2(w, "Success", m, platform)
}
