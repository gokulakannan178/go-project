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

// SaveEmployeeSalary : ""
func (h *Handler) SaveEmployeeSalary(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeSalary := new(models.EmployeeSalary)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeSalary)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeSalary(ctx, employeeSalary)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeSalary"] = employeeSalary
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeSalary : ""
func (h *Handler) GetSingleEmployeeSalary(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeeSalary := new(models.RefEmployeeSalary)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeSalary, err := h.Service.GetSingleEmployeeSalary(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeSalary"] = employeeSalary
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeSalary : ""
func (h *Handler) UpdateEmployeeSalary(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeSalary := new(models.EmployeeSalary)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeSalary)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeSalary.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeSalary(ctx, employeeSalary)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeSalary"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeSalary : ""
func (h *Handler) EnableEmployeeSalary(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeSalary(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeSalary"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeSalary : ""
func (h *Handler) DisableEmployeeSalary(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeSalary(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeSalary"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeeSalary : ""
func (h *Handler) DeleteEmployeeSalary(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeSalary(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeSalary"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeSalary : ""
func (h *Handler) FilterEmployeeSalary(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeeSalary *models.FilterEmployeeSalary
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeeSalary)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeeSalarys []models.RefEmployeeSalary
	log.Println(pagination)
	filterEmployeeSalarys, err = h.Service.FilterEmployeeSalary(ctx, filterEmployeeSalary, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeeSalarys) > 0 {
		m["EmployeeSalary"] = filterEmployeeSalarys
	} else {
		res := make([]models.EmployeeSalary, 0)
		m["EmployeeSalary"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveEmployeeSalaryWithEmployee(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeSalary := new(models.FilterEmployeeSalary)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeSalary)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	Salaryerrors, err := h.Service.SaveEmployeeSalaryWithEmployee(ctx, employeeSalary)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeSalary"] = Salaryerrors
	response.With200V2(w, "Success", m, platform)
}
