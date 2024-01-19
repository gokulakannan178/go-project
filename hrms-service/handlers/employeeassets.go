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

// SaveEmployeeAssets : ""
func (h *Handler) SaveEmployeeAssets(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeAssets := new(models.EmployeeAssets)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeAssets)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeAssets(ctx, employeeAssets)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAssets"] = employeeAssets
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeAssets : ""
func (h *Handler) GetSingleEmployeeAssets(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeeAssets := new(models.RefEmployeeAssets)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeAssets, err := h.Service.GetSingleEmployeeAssets(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAssets"] = employeeAssets
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeAssets : ""
func (h *Handler) UpdateEmployeeAssets(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeAssets := new(models.EmployeeAssets)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeAssets)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeAssets.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeAssets(ctx, employeeAssets)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAssets"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeAssets : ""
func (h *Handler) EnableEmployeeAssets(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeAssets(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAssets"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeAssets : ""
func (h *Handler) DisableEmployeeAssets(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeAssets(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAssets"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeeAssets : ""
func (h *Handler) DeleteEmployeeAssets(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeAssets(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAssets"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeAssets : ""
func (h *Handler) FilterEmployeeAssets(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeeAssets *models.FilterEmployeeAssets
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeeAssets)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeeAssetss []models.RefEmployeeAssets
	log.Println(pagination)
	filterEmployeeAssetss, err = h.Service.FilterEmployeeAssets(ctx, filterEmployeeAssets, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeeAssetss) > 0 {
		m["EmployeeAssets"] = filterEmployeeAssetss
	} else {
		res := make([]models.EmployeeAssets, 0)
		m["EmployeeAssets"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
