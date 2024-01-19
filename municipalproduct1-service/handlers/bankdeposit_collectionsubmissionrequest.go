package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

//CollectionSubmissionRequest : ""
func (h *Handler) CollectionSubmissionRequest(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	collectionSubmissionRequest := new(models.CollectionSubmissionRequest)
	err := json.NewDecoder(r.Body).Decode(&collectionSubmissionRequest)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.CollectionSubmissionRequest(ctx, collectionSubmissionRequest)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["collectionSubmissionRequest"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//CollectionSubmissionRequestFilter
func (h *Handler) CollectionSubmissionRequestFilter(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	csr := new(models.CollectionSubmissionRequestFilter)
	err := json.NewDecoder(r.Body).Decode(&csr)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	var pagination *models.Pagination
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
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

	data, err := h.Service.CollectionSubmissionRequestFilter(ctx, csr, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["collectionsubmissionrequest"] = data
	m["pagination"] = pagination
	response.With200V2(w, "Success", m, platform)
}
