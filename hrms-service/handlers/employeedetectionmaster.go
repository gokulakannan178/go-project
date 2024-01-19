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

// SaveEmployeeDeductionMaster : ""
func (h *Handler) SaveEmployeeDeductionMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeDeductionMaster := new(models.EmployeeDeductionMaster)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeDeductionMaster)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeDeductionMaster(ctx, employeeDeductionMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeductionMaster"] = employeeDeductionMaster
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeDeductionMaster : ""
func (h *Handler) GetSingleEmployeeDeductionMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeeDeductionMaster := new(models.RefEmployeeDeductionMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeDeductionMaster, err := h.Service.GetSingleEmployeeDeductionMaster(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeductionMaster"] = employeeDeductionMaster
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeDeductionMaster : ""
func (h *Handler) UpdateEmployeeDeductionMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeDeductionMaster := new(models.EmployeeDeductionMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeDeductionMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeDeductionMaster.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeDeductionMaster(ctx, employeeDeductionMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeductionMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeDeductionMaster : ""
func (h *Handler) EnableEmployeeDeductionMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeDeductionMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeductionMaster"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeDeductionMaster : ""
func (h *Handler) DisableEmployeeDeductionMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeDeductionMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeductionMaster"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeeDeductionMaster : ""
func (h *Handler) DeleteEmployeeDeductionMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeDeductionMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeDeductionMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeDeductionMaster : ""
func (h *Handler) FilterEmployeeDeductionMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeeDeductionMaster *models.FilterEmployeeDeductionMaster
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeeDeductionMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeeDeductionMasters []models.RefEmployeeDeductionMaster
	log.Println(pagination)
	filterEmployeeDeductionMasters, err = h.Service.FilterEmployeeDeductionMaster(ctx, filterEmployeeDeductionMaster, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeeDeductionMasters) > 0 {
		m["EmployeeDeductionMaster"] = filterEmployeeDeductionMasters
	} else {
		res := make([]models.EmployeeDeductionMaster, 0)
		m["EmployeeDeductionMaster"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
