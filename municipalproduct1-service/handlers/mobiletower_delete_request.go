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

// BasicUpdateMobileTowerDeleteRequest : ""
func (h *Handler) BasicUpdateMobileTowerDeleteRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	request := new(models.MobileTowerDeleteRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.BasicUpdateMobileTowerDeleteRequest(ctx, request)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basicMobileTowerDeleteRequestUpdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleMobileTowerDeleteRequest : ""
func (h *Handler) GetSingleMobileTowerDeleteRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	pmr := new(models.RefMobileTowerDeleteRequest)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	pmr, err := h.Service.GetSingleMobileTowerDeleteRequest(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tMobileTowerDeleteRequest"] = pmr
	response.With200V2(w, "Success", m, platform)
}

// AcceptMobileTowerDeleteRequestUpdate : ""
func (h *Handler) AcceptMobileTowerDeleteRequestUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.AcceptMobileTowerDeleteRequestUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.AcceptMobileTowerDeleteRequestUpdate(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["acceptMobileTowerDeleteRequest"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// RejectMobileTowerDeleteRequestUpdate : ""
func (h *Handler) RejectMobileTowerDeleteRequestUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.RejectMobileTowerDeleteRequestUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.RejectMobileTowerDeleteRequestUpdate(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["rejectMobileTowerDeleteRequest"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterMobileTowerDeleteRequest : ""
func (h *Handler) FilterMobileTowerDeleteRequest(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.MobileTowerDeleteRequestFilter
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

	var requests []models.RefMobileTowerDeleteRequest
	log.Println(pagination)
	requests, err = h.Service.FilterMobileTowerDeleteRequest(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(requests) > 0 {
		m["mobileTowerDeleteRequest"] = requests
	} else {
		res := make([]models.RefMobileTowerDeleteRequest, 0)
		m["mobileTowerDeleteRequest"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
