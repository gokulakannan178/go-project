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

//SaveDepartmentType : ""
func (h *Handler) SaveDepartmentType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	departmentType := new(models.DepartmentType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&departmentType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDepartmentType(ctx, departmentType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["departmentType"] = departmentType
	response.With200V2(w, "Success", m, platform)
}

//UpdateDepartmentType :""
func (h *Handler) UpdateDepartmentType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	departmentType := new(models.DepartmentType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&departmentType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if departmentType.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateDepartmentType(ctx, departmentType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["departmentType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDepartmentType : ""
func (h *Handler) EnableDepartmentType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableDepartmentType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["departmentType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDepartmentType : ""
func (h *Handler) DisableDepartmentType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableDepartmentType(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["departmentType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDepartmentType : ""
func (h *Handler) DeleteDepartmentType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteDepartmentType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["departmentType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDepartmentType :""
func (h *Handler) GetSingleDepartmentType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	departmentType := new(models.RefDepartmentType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	departmentType, err := h.Service.GetSingleDepartmentType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["departmentType"] = departmentType
	response.With200V2(w, "Success", m, platform)
}

//FilterDepartmentType : ""
func (h *Handler) FilterDepartmentType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var departmentType *models.DepartmentTypeFilter
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
	err := json.NewDecoder(r.Body).Decode(&departmentType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var departmentTypes []models.RefDepartmentType
	log.Println(pagination)
	departmentTypes, err = h.Service.FilterDepartmentType(ctx, departmentType, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(departmentTypes) > 0 {
		m["departmentType"] = departmentTypes
	} else {
		res := make([]models.DepartmentType, 0)
		m["departmentType"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
