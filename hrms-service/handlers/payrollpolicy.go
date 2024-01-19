package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/constants"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

// SavePayrollPolicy : ""
func (h *Handler) SavePayrollPolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	payrollPolicy := new(models.PayrollPolicy)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&payrollPolicy)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SavePayrollPolicy(ctx, payrollPolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicy"] = payrollPolicy
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePayrollPolicy : ""
func (h *Handler) GetSinglePayrollPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	payrollPolicy := new(models.RefPayrollPolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	payrollPolicy, err := h.Service.GetSinglePayrollPolicy(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicy"] = payrollPolicy
	response.With200V2(w, "Success", m, platform)
}

//UpdatePayrollPolicy : ""
func (h *Handler) UpdatePayrollPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	payrollPolicy := new(models.PayrollPolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&payrollPolicy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if payrollPolicy.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePayrollPolicy(ctx, payrollPolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnablePayrollPolicy : ""
func (h *Handler) EnablePayrollPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnablePayrollPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicy"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisablePayrollPolicy : ""
func (h *Handler) DisablePayrollPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisablePayrollPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicy"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeletePayrollPolicy : ""
func (h *Handler) DeletePayrollPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeletePayrollPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterPayrollPolicy : ""
func (h *Handler) FilterPayrollPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterPayrollPolicy *models.FilterPayrollPolicy
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
	err := json.NewDecoder(r.Body).Decode(&filterPayrollPolicy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterPayrollPolicys []models.RefPayrollPolicy
	log.Println(pagination)
	filterPayrollPolicys, err = h.Service.FilterPayrollPolicy(ctx, filterPayrollPolicy, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterPayrollPolicys) > 0 {
		m["PayrollPolicy"] = filterPayrollPolicys
	} else {
		res := make([]models.PayrollPolicy, 0)
		m["PayrollPolicy"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSalaryCalc(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("amount")

	if UniqueID == "" {
		response.With400V2(w, "amount is missing", platform)
		return
	}

	payrollPolicy := new(models.SalaryCalc)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	payrollPolicy, err := h.Service.GetSalaryCalc(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicy"] = payrollPolicy
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSalaryCalcWithEmployeeType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("amount")
	employeeType := r.URL.Query().Get("employeeType")

	if UniqueID == "" {
		response.With400V2(w, "amount is missing", platform)
		return
	}
	if employeeType == "" {
		response.With400V2(w, "employeeType is missing", platform)
		return
	}
	if employeeType == constants.EMPLOYEESTATUSONBORADING {
		employeeType = constants.EMPLOYEESTATUSPROBATIONARY
	}
	payrollPolicy := new(models.SalaryCalc)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	payrollPolicy, err := h.Service.GetSalaryCalcV2(ctx, UniqueID, employeeType)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollPolicy"] = payrollPolicy
	response.With200V2(w, "Success", m, platform)
}
