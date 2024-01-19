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

// SaveEmployeeDeduction : ""
func (h *Handler) SaveEmployeeDeduction(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeDeduction := new(models.EmployeeDeduction)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeDeduction)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeDeduction(ctx, employeeDeduction)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeduction"] = employeeDeduction
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeDeduction : ""
func (h *Handler) GetSingleEmployeeDeduction(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeeDeduction := new(models.RefEmployeeDeduction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeDeduction, err := h.Service.GetSingleEmployeeDeduction(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeduction"] = employeeDeduction
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeDeduction : ""
func (h *Handler) UpdateEmployeeDeduction(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeDeduction := new(models.EmployeeDeduction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeDeduction)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeDeduction.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeDeduction(ctx, employeeDeduction)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeduction"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeDeduction : ""
func (h *Handler) EnableEmployeeDeduction(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeDeduction(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeduction"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeDeduction : ""
func (h *Handler) DisableEmployeeDeduction(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeDeduction(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeduction"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeeDeduction : ""
func (h *Handler) DeleteEmployeeDeduction(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeDeduction(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeduction"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeDeduction : ""
func (h *Handler) FilterEmployeeDeduction(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeeDeduction *models.FilterEmployeeDeduction
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeeDeduction)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeeDeductions []models.RefEmployeeDeduction
	log.Println(pagination)
	filterEmployeeDeductions, err = h.Service.FilterEmployeeDeduction(ctx, filterEmployeeDeduction, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeeDeductions) > 0 {
		m["EmployeeDeduction"] = filterEmployeeDeductions
	} else {
		res := make([]models.EmployeeDeduction, 0)
		m["EmployeeDeduction"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
