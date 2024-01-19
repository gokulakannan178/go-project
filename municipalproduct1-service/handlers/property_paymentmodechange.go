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

// SavePropertyPaymentModeChange : ""
func (h *Handler) SavePropertyPaymentModeChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	request := new(models.PropertyPaymentModeChangeRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.SavePropertyPaymentModeChange(ctx, request)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyPaymentModeChange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleReassessmentRequest :""
func (h *Handler) GetSinglePropertyPaymentModeChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "uniqueId is missing", platform)
	}

	rppmcr := new(models.RefPropertyPaymentModeChangeRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	rppmcr, err := h.Service.GetSinglePropertyPaymentModeChange(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyPaymentModeChange"] = rppmcr
	response.With200V2(w, "Success", m, platform)
}

// AcceptPropertyPaymentModeChange : ""
func (h *Handler) AcceptPropertyPaymentModeChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.AcceptPropertyPaymentModeChangeRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.AcceptPropertyPaymentModeChange(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["acceptPropertyPaymentModeChange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// RejectPropertyPaymentModeChange : ""
func (h *Handler) RejectPropertyPaymentModeChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.RejectPropertyPaymentModeChangeRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.RejectPropertyPaymentModeChange(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["rejectPropertyPaymentModeChange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyPaymentModeChange : ""
func (h *Handler) FilterPropertyPaymentModeChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PropertyPaymentModeChangeRequestFilter
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

	var requests []models.RefPropertyPaymentModeChangeRequest
	log.Println(pagination)
	requests, err = h.Service.FilterPropertyPaymentModeChange(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(requests) > 0 {
		m["propertyPaymentModeChange"] = requests
	} else {
		res := make([]models.RefPropertyPaymentModeChangeRequest, 0)
		m["propertyPaymentModeChange"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyPaymentModeChangePropertyID : ""
func (h *Handler) UpdatePropertyPaymentModeChangePropertyID(w http.ResponseWriter, r *http.Request) {

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
	err = h.Service.UpdatePropertyPaymentModeChangePropertyID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["updatePropertyPaymentModeChangePropertyID"] = "success"
	response.With200V2(w, "Success", m, platform)
}
