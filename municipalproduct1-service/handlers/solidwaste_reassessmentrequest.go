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

// BasicUpdateSolidWasteReassessmentRequest : ""
func (h *Handler) BasicUpdateSolidWasteReassessmentRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	request := new(models.SolidWasteReassessmentRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.BasicUpdateSolidWasteReassessmentRequest(ctx, request)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basicSolidWasteReassessmentRequestUpdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleSolidWasteReassessmentRequest : ""
func (h *Handler) GetSingleSolidWasteReassessmentRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	res := new(models.RefSolidWasteReassessmentRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	res, err := h.Service.GetSingleSolidWasteReassessmentRequest(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidWasteReassessmentRequest"] = res
	response.With200V2(w, "Success", m, platform)
}

// AcceptSolidWasteReassessmentRequestUpdate : ""
func (h *Handler) AcceptSolidWasteReassessmentRequestUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.AcceptSolidWasteReassessmentRequestUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.AcceptSolidWasteReassessmentRequestUpdate(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["acceptSolidWasteReassessmentRequest"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// RejectSolidWasteReassessmentRequestUpdate : ""
func (h *Handler) RejectSolidWasteReassessmentRequestUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.RejectSolidWasteReassessmentRequestUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.RejectSolidWasteReassessmentRequestUpdate(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["rejectSolidWasteReassessmentRequest"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterSolidWasteReassessmentRequest : ""
func (h *Handler) FilterSolidWasteReassessmentRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.SolidWasteReassessmentRequestFilter
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

	var requests []models.RefSolidWasteReassessmentRequest
	log.Println(pagination)
	requests, err = h.Service.FilterSolidWasteReassessmentRequest(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(requests) > 0 {
		m["solidWasteReassessmentRequest"] = requests
	} else {
		res := make([]models.RefSolidWasteReassessmentRequest, 0)
		m["solidWasteReassessmentRequest"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
