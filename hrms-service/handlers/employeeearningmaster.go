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

// SaveEmployeeEarningMaster : ""
func (h *Handler) SaveEmployeeEarningMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeEarningMaster := new(models.EmployeeEarningMaster)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeEarningMaster)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeEarningMaster(ctx, employeeEarningMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEarningMaster"] = employeeEarningMaster
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeEarningMaster : ""
func (h *Handler) GetSingleEmployeeEarningMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeeEarningMaster := new(models.RefEmployeeEarningMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeEarningMaster, err := h.Service.GetSingleEmployeeEarningMaster(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEarningMaster"] = employeeEarningMaster
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeEarningMaster : ""
func (h *Handler) UpdateEmployeeEarningMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeEarningMaster := new(models.EmployeeEarningMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeEarningMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeEarningMaster.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeEarningMaster(ctx, employeeEarningMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEarningMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeEarningMaster : ""
func (h *Handler) EnableEmployeeEarningMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeEarningMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEarningMaster"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeEarningMaster : ""
func (h *Handler) DisableEmployeeEarningMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeEarningMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEarningMaster"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeeEarningMaster : ""
func (h *Handler) DeleteEmployeeEarningMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeEarningMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeEarningMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeEarningMaster : ""
func (h *Handler) FilterEmployeeEarningMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeeEarningMaster *models.FilterEmployeeEarningMaster
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeeEarningMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeeEarningMasters []models.RefEmployeeEarningMaster
	log.Println(pagination)
	filterEmployeeEarningMasters, err = h.Service.FilterEmployeeEarningMaster(ctx, filterEmployeeEarningMaster, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeeEarningMasters) > 0 {
		m["EmployeeEarningMaster"] = filterEmployeeEarningMasters
	} else {
		res := make([]models.EmployeeEarningMaster, 0)
		m["EmployeeEarningMaster"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
