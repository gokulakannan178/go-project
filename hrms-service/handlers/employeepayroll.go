package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

// SaveEmployeePayroll : ""
func (h *Handler) SaveEmployeePayroll(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeePayroll := new(models.EmployeePayroll)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeePayroll)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeePayroll(ctx, employeePayroll)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeePayroll"] = employeePayroll
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeePayroll : ""
func (h *Handler) GetSingleEmployeePayroll(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeePayroll := new(models.RefEmployeePayroll)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeePayroll, err := h.Service.GetSingleEmployeePayroll(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeePayroll"] = employeePayroll
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeePayroll : ""
func (h *Handler) UpdateEmployeePayroll(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeePayroll := new(models.EmployeePayroll)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeePayroll)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeePayroll.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeePayroll(ctx, employeePayroll)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeePayroll"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeePayroll : ""
func (h *Handler) EnableEmployeePayroll(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeePayroll(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeePayroll"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeePayroll : ""
func (h *Handler) DisableEmployeePayroll(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeePayroll(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeePayroll"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeePayroll : ""
func (h *Handler) DeleteEmployeePayroll(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeePayroll(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeePayroll"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeePayroll : ""
func (h *Handler) FilterEmployeePayroll(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeePayroll *models.FilterEmployeePayroll
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeePayroll)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeePayrolls []models.RefEmployeePayroll
	log.Println(pagination)
	filterEmployeePayrolls, err = h.Service.FilterEmployeePayroll(ctx, filterEmployeePayroll, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeePayrolls) > 0 {
		m["EmployeePayroll"] = filterEmployeePayrolls
	} else {
		res := make([]models.EmployeePayroll, 0)
		m["EmployeePayroll"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) EmployeeUpdatePayrollWithNewPayroll(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeePayroll := new(models.EmployeePayrollWithEarningDeduction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeePayroll)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeePayroll.EmployeeId == "" {
		response.With400V2(w, "Employeeid is missing", platform)
	}
	err = h.Service.EmployeeUpdatePayrollWithNewPayroll(ctx, employeePayroll)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeePayroll"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveEmployeePayrollWithEaringDeduction(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeePayroll := new(models.EmployeePayrollWithEarningDeduction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeePayroll)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeePayroll.EmployeeId == "" {
		response.With400V2(w, "Employeeid is missing", platform)
	}
	err = h.Service.SaveEmployeePayrollWithEaringDeduction(ctx, employeePayroll)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeePayroll"] = employeePayroll
	response.With200V2(w, "Success", m, platform)
}
