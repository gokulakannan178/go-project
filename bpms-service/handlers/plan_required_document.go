package handlers

import (
	"bpms-service/app"
	"bpms-service/models"
	"bpms-service/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//SavePlanReqDocument : ""
func (h *Handler) SavePlanReqDocument(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	planReqDocument := new(models.PlanReqDocument)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&planReqDocument)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePlanReqDocument(ctx, planReqDocument)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planReqDocument"] = planReqDocument
	response.With200V2(w, "Success", m, platform)
}

//UpdatePlanReqDocument :""
func (h *Handler) UpdatePlanReqDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	planReqDocument := new(models.PlanReqDocument)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&planReqDocument)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if planReqDocument.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePlanReqDocument(ctx, planReqDocument)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planReqDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePlanReqDocument : ""
func (h *Handler) EnablePlanReqDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnablePlanReqDocument(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planReqDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisablePlanReqDocument : ""
func (h *Handler) DisablePlanReqDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisablePlanReqDocument(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planReqDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePlanReqDocument : ""
func (h *Handler) DeletePlanReqDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeletePlanReqDocument(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planReqDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePlanReqDocument :""
func (h *Handler) GetSinglePlanReqDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	planReqDocument := new(models.RefPlanReqDocument)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	planReqDocument, err := h.Service.GetSinglePlanReqDocument(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planReqDocument"] = planReqDocument
	response.With200V2(w, "Success", m, platform)
}

//FilterPlanReqDocument : ""
func (h *Handler) FilterPlanReqDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var planReqDocument *models.PlanReqDocumentFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
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
	err := json.NewDecoder(r.Body).Decode(&planReqDocument)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var planReqDocuments []models.RefPlanReqDocument
	log.Println(pagination)
	planReqDocuments, err = h.Service.FilterPlanReqDocument(ctx, planReqDocument, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(planReqDocuments) > 0 {
		m["planReqDocument"] = planReqDocuments
	} else {
		res := make([]models.PlanReqDocument, 0)
		m["planReqDocument"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
