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

//SavePlanDepartmentApproval : ""
func (h *Handler) SavePlanDepartmentApproval(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	planDepartmentApproval := new(models.PlanDepartmentApproval)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&planDepartmentApproval)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePlanDepartmentApproval(ctx, planDepartmentApproval)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDepartmentApproval"] = planDepartmentApproval
	response.With200V2(w, "Success", m, platform)
}

//SaveMultiplePlanDepartmentApproval : ""
func (h *Handler) SaveMultiplePlanDepartmentApproval(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	planDepartmentApproval := []models.PlanDepartmentApproval{}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&planDepartmentApproval)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveMultiplePlanDepartmentApproval(ctx, planDepartmentApproval)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDepartmentApproval"] = planDepartmentApproval
	response.With200V2(w, "Success", m, platform)
}

//UpdatePlanDepartmentApproval :""
func (h *Handler) UpdatePlanDepartmentApproval(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	planDepartmentApproval := new(models.PlanDepartmentApproval)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&planDepartmentApproval)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if planDepartmentApproval.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdatePlanDepartmentApproval(ctx, planDepartmentApproval)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDepartmentApproval"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePlanDepartmentApproval : ""
func (h *Handler) EnablePlanDepartmentApproval(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnablePlanDepartmentApproval(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDepartmentApproval"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisablePlanDepartmentApproval : ""
func (h *Handler) DisablePlanDepartmentApproval(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisablePlanDepartmentApproval(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDepartmentApproval"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePlanDepartmentApproval : ""
func (h *Handler) DeletePlanDepartmentApproval(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeletePlanDepartmentApproval(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDepartmentApproval"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePlanDepartmentApproval :""
func (h *Handler) GetSinglePlanDepartmentApproval(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	planDepartmentApproval := new(models.RefPlanDepartmentApproval)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	planDepartmentApproval, err := h.Service.GetSinglePlanDepartmentApproval(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["planDepartmentApproval"] = planDepartmentApproval
	response.With200V2(w, "Success", m, platform)
}

//FilterPlanDepartmentApproval : ""
func (h *Handler) FilterPlanDepartmentApproval(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var planDepartmentApproval *models.PlanDepartmentApprovalFilter
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
	err := json.NewDecoder(r.Body).Decode(&planDepartmentApproval)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var planDepartmentApprovals []models.RefPlanDepartmentApproval
	log.Println(pagination)
	planDepartmentApprovals, err = h.Service.FilterPlanDepartmentApproval(ctx, planDepartmentApproval, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(planDepartmentApprovals) > 0 {
		m["planDepartmentApproval"] = planDepartmentApprovals
	} else {
		res := make([]models.PlanDepartmentApproval, 0)
		m["planDepartmentApproval"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//GetAPlanDeptsApproval :""
func (h *Handler) GetAPlanDeptsApproval(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	deptID := r.URL.Query().Get("deptId")

	if deptID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	getAPlanDeptsApproval := new(models.GetAPlanDeptsApproval)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	getAPlanDeptsApproval, err := h.Service.GetAPlanDeptsApproval(ctx, deptID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["getAPlanDeptsApproval"] = getAPlanDeptsApproval
	response.With200V2(w, "Success", m, platform)
}

//GetAPlanDeptsApprovalV2 :""
func (h *Handler) GetAPlanDeptsApprovalV2(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	planID := r.URL.Query().Get("planId")

	if planID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	getAPlanDeptsApproval := new(models.GetAPlanDeptsApprovalV2)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	getAPlanDeptsApproval, err := h.Service.GetAPlanDeptsApprovalV2(ctx, planID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["getAPlanDeptsApproval"] = getAPlanDeptsApproval
	response.With200V2(w, "Success", m, platform)
}

//GetAPlanDeptsApprovalV3 :""
func (h *Handler) GetAPlanDeptsApprovalV3(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	planID := r.URL.Query().Get("planId")

	if planID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	getAPlanDeptsApproval := new(models.GetAPlanDeptsApprovalV3)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	getAPlanDeptsApproval, err := h.Service.GetAPlanDeptsApprovalV3(ctx, planID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["getAPlanDeptsApproval"] = getAPlanDeptsApproval
	response.With200V2(w, "Success", m, platform)
}
