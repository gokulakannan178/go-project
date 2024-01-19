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

//SavePlan : ""
func (h *Handler) SavePlan(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	plan := new(models.Plan)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&plan)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePlan(ctx, plan)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["plan"] = plan
	response.With200V2(w, "Success", m, platform)
}

//UpdatePlan :""
func (h *Handler) UpdatePlan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	plan := new(models.Plan)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&plan)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if plan.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePlan(ctx, plan)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["plan"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePlan : ""
func (h *Handler) EnablePlan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnablePlan(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["plan"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisablePlan : ""
func (h *Handler) DisablePlan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisablePlan(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["plan"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePlan : ""
func (h *Handler) DeletePlan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeletePlan(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["plan"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePlan :""
func (h *Handler) GetSinglePlan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	plan := new(models.RefPlan)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	plan, err := h.Service.GetSinglePlan(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["plan"] = plan
	response.With200V2(w, "Success", m, platform)
}

//FilterPlan : ""
func (h *Handler) FilterPlan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var plan *models.PlanFilter
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
	err := json.NewDecoder(r.Body).Decode(&plan)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var plans []models.RefPlan
	log.Println(pagination)
	plans, err = h.Service.FilterPlan(ctx, plan, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(plans) > 0 {
		m["plan"] = plans
	} else {
		res := make([]models.Plan, 0)
		m["plan"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
