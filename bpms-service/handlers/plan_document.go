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

//SavePlanDocument : ""
func (h *Handler) SavePlanDocument(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	planDocument := new(models.PlanDocument)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&planDocument)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePlanDocument(ctx, planDocument)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDocument"] = planDocument
	response.With200V2(w, "Success", m, platform)
}

//UpdatePlanDocument :""
func (h *Handler) UpdatePlanDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	planDocument := new(models.PlanDocument)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&planDocument)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if planDocument.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePlanDocument(ctx, planDocument)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePlanDocument : ""
func (h *Handler) EnablePlanDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnablePlanDocument(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisablePlanDocument : ""
func (h *Handler) DisablePlanDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisablePlanDocument(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePlanDocument : ""
func (h *Handler) DeletePlanDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeletePlanDocument(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDocument"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePlanDocument :""
func (h *Handler) GetSinglePlanDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	planDocument := new(models.RefPlanDocument)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	planDocument, err := h.Service.GetSinglePlanDocument(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDocument"] = planDocument
	response.With200V2(w, "Success", m, platform)
}

//FilterPlanDocument : ""
func (h *Handler) FilterPlanDocument(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var planDocument *models.PlanDocumentFilter
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
	err := json.NewDecoder(r.Body).Decode(&planDocument)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var planDocuments []models.RefPlanDocument
	log.Println(pagination)
	planDocuments, err = h.Service.FilterPlanDocument(ctx, planDocument, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(planDocuments) > 0 {
		m["planDocument"] = planDocuments
	} else {
		res := make([]models.PlanDocument, 0)
		m["planDocument"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//GetPendingDocuments : ""
func (h *Handler) GetPendingDocuments(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var planDocument *models.GetPendingPlanDocumentFilter
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
	err := json.NewDecoder(r.Body).Decode(&planDocument)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var planDocuments []models.RefPlanDocument
	log.Println(pagination)
	planDocuments, err = h.Service.GetPendingDocuments(ctx, planDocument, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(planDocuments) > 0 {
		m["planDocument"] = planDocuments
	} else {
		res := make([]models.PlanDocument, 0)
		m["planDocument"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
