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

// SaveEmployeeEducation : ""
func (h *Handler) SaveEmployeeEducation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeEducation := new(models.EmployeeEducation)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeEducation)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeEducation(ctx, employeeEducation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEducation"] = employeeEducation
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeEducation : ""
func (h *Handler) GetSingleEmployeeEducation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeeEducation := new(models.RefEmployeeEducation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeEducation, err := h.Service.GetSingleEmployeeEducation(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEducation"] = employeeEducation
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeEducation : ""
func (h *Handler) UpdateEmployeeEducation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeEducation := new(models.EmployeeEducation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeEducation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeEducation.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeEducation(ctx, employeeEducation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEducation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeEducation : ""
func (h *Handler) EnableEmployeeEducation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeEducation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEducation"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeEducation : ""
func (h *Handler) DisableEmployeeEducation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeEducation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEducation"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeeEducation : ""
func (h *Handler) DeleteEmployeeEducation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeEducation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEducation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeEducation : ""
func (h *Handler) FilterEmployeeEducation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeeEducation *models.FilterEmployeeEducation
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeeEducation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeeEducations []models.RefEmployeeEducation
	log.Println(pagination)
	filterEmployeeEducations, err = h.Service.FilterEmployeeEducation(ctx, filterEmployeeEducation, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeeEducations) > 0 {
		m["EmployeeEducation"] = filterEmployeeEducations
	} else {
		res := make([]models.EmployeeEducation, 0)
		m["EmployeeEducation"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
