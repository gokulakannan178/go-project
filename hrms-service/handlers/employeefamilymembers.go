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

// SaveEmployeeFamilyMembers : ""
func (h *Handler) SaveEmployeeFamilyMembers(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeFamilyMembers := new(models.EmployeeFamilyMembers)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeFamilyMembers)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeFamilyMembers(ctx, employeeFamilyMembers)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeFamilyMembers"] = employeeFamilyMembers
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeFamilyMembers : ""
func (h *Handler) GetSingleEmployeeFamilyMembers(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeeFamilyMembers := new(models.RefEmployeeFamilyMembers)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeFamilyMembers, err := h.Service.GetSingleEmployeeFamilyMembers(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeFamilyMembers"] = employeeFamilyMembers
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeFamilyMembers : ""
func (h *Handler) UpdateEmployeeFamilyMembers(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeFamilyMembers := new(models.EmployeeFamilyMembers)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeFamilyMembers)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeFamilyMembers.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeFamilyMembers(ctx, employeeFamilyMembers)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeFamilyMembers"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeFamilyMembers : ""
func (h *Handler) EnableEmployeeFamilyMembers(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeFamilyMembers(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeFamilyMembers"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeFamilyMembers : ""
func (h *Handler) DisableEmployeeFamilyMembers(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeFamilyMembers(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeFamilyMembers"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeeFamilyMembers : ""
func (h *Handler) DeleteEmployeeFamilyMembers(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeFamilyMembers(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeFamilyMembers"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeFamilyMembers : ""
func (h *Handler) FilterEmployeeFamilyMembers(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeeFamilyMembers *models.FilterEmployeeFamilyMembers
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeeFamilyMembers)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeeFamilyMemberss []models.RefEmployeeFamilyMembers
	log.Println(pagination)
	filterEmployeeFamilyMemberss, err = h.Service.FilterEmployeeFamilyMembers(ctx, filterEmployeeFamilyMembers, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeeFamilyMemberss) > 0 {
		m["EmployeeFamilyMembers"] = filterEmployeeFamilyMemberss
	} else {
		res := make([]models.EmployeeFamilyMembers, 0)
		m["EmployeeFamilyMembers"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
