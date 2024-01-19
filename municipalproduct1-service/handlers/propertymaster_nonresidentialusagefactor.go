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

//SaveNonResidentialUsageFactor : ""
func (h *Handler) SaveNonResidentialUsageFactor(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	nonResidentialUsageFactor := new(models.NonResidentialUsageFactor)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&nonResidentialUsageFactor)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveNonResidentialUsageFactor(ctx, nonResidentialUsageFactor)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["nonResidentialUsageFactor"] = nonResidentialUsageFactor
	response.With200V2(w, "Success", m, platform)
}

//UpdateNonResidentialUsageFactor :""
func (h *Handler) UpdateNonResidentialUsageFactor(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	nonResidentialUsageFactor := new(models.NonResidentialUsageFactor)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&nonResidentialUsageFactor)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if nonResidentialUsageFactor.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateNonResidentialUsageFactor(ctx, nonResidentialUsageFactor)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["nonResidentialUsageFactor"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableNonResidentialUsageFactor : ""
func (h *Handler) EnableNonResidentialUsageFactor(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableNonResidentialUsageFactor(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["nonResidentialUsageFactor"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableNonResidentialUsageFactor : ""
func (h *Handler) DisableNonResidentialUsageFactor(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableNonResidentialUsageFactor(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["nonResidentialUsageFactor"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteNonResidentialUsageFactor : ""
func (h *Handler) DeleteNonResidentialUsageFactor(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteNonResidentialUsageFactor(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["nonResidentialUsageFactor"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleNonResidentialUsageFactor :""
func (h *Handler) GetSingleNonResidentialUsageFactor(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	nonResidentialUsageFactor := new(models.RefNonResidentialUsageFactor)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	nonResidentialUsageFactor, err := h.Service.GetSingleNonResidentialUsageFactor(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["nonResidentialUsageFactor"] = nonResidentialUsageFactor
	response.With200V2(w, "Success", m, platform)
}

//FilterNonResidentialUsageFactor : ""
func (h *Handler) FilterNonResidentialUsageFactor(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var nonResidentialUsageFactor *models.NonResidentialUsageFactorFilter
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
	err := json.NewDecoder(r.Body).Decode(&nonResidentialUsageFactor)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var nonResidentialUsageFactors []models.RefNonResidentialUsageFactor
	log.Println(pagination)
	nonResidentialUsageFactors, err = h.Service.FilterNonResidentialUsageFactor(ctx, nonResidentialUsageFactor, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(nonResidentialUsageFactors) > 0 {
		m["nonResidentialUsageFactor"] = nonResidentialUsageFactors
	} else {
		res := make([]models.NonResidentialUsageFactor, 0)
		m["nonResidentialUsageFactor"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
