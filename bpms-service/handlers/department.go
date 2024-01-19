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

//SaveDepartment : ""
func (h *Handler) SaveDepartment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	department := new(models.Department)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&department)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDepartment(ctx, department)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["department"] = department
	response.With200V2(w, "Success", m, platform)
}

//UpdateDepartment :""
func (h *Handler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	department := new(models.Department)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&department)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if department.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateDepartment(ctx, department)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["department"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDepartment : ""
func (h *Handler) EnableDepartment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableDepartment(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["department"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDepartment : ""
func (h *Handler) DisableDepartment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableDepartment(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["department"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDepartment : ""
func (h *Handler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteDepartment(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["department"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDepartment :""
func (h *Handler) GetSingleDepartment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	department := new(models.RefDepartment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	department, err := h.Service.GetSingleDepartment(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["department"] = department
	response.With200V2(w, "Success", m, platform)
}

//FilterDepartment : ""
func (h *Handler) FilterDepartment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var department *models.DepartmentFilter
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
	err := json.NewDecoder(r.Body).Decode(&department)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var departments []models.RefDepartment
	log.Println(pagination)
	departments, err = h.Service.FilterDepartment(ctx, department, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(departments) > 0 {
		m["department"] = departments
	} else {
		res := make([]models.Department, 0)
		m["department"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
