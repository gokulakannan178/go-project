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

//SavePlanRegistrationType : ""
func (h *Handler) SavePlanRegistrationType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	planRegistrationType := new(models.PlanRegistrationType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&planRegistrationType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePlanRegistrationType(ctx, planRegistrationType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planRegistrationType"] = planRegistrationType
	response.With200V2(w, "Success", m, platform)
}

//UpdatePlanRegistrationType :""
func (h *Handler) UpdatePlanRegistrationType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	planRegistrationType := new(models.PlanRegistrationType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&planRegistrationType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if planRegistrationType.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePlanRegistrationType(ctx, planRegistrationType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planRegistrationType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePlanRegistrationType : ""
func (h *Handler) EnablePlanRegistrationType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnablePlanRegistrationType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planRegistrationType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisablePlanRegistrationType : ""
func (h *Handler) DisablePlanRegistrationType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisablePlanRegistrationType(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planRegistrationType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePlanRegistrationType : ""
func (h *Handler) DeletePlanRegistrationType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeletePlanRegistrationType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planRegistrationType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePlanRegistrationType :""
func (h *Handler) GetSinglePlanRegistrationType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	planRegistrationType := new(models.RefPlanRegistrationType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	planRegistrationType, err := h.Service.GetSinglePlanRegistrationType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planRegistrationType"] = planRegistrationType
	response.With200V2(w, "Success", m, platform)
}

//FilterPlanRegistrationType : ""
func (h *Handler) FilterPlanRegistrationType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var planRegistrationType *models.PlanRegistrationTypeFilter
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
	err := json.NewDecoder(r.Body).Decode(&planRegistrationType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var planRegistrationTypes []models.RefPlanRegistrationType
	log.Println(pagination)
	planRegistrationTypes, err = h.Service.FilterPlanRegistrationType(ctx, planRegistrationType, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(planRegistrationTypes) > 0 {
		m["planRegistrationType"] = planRegistrationTypes
	} else {
		res := make([]models.PlanRegistrationType, 0)
		m["planRegistrationType"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
