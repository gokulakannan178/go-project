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

//SaveDeptChecklist : ""
func (h *Handler) SaveDeptChecklist(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	deptChecklist := new(models.DeptChecklist)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&deptChecklist)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDeptChecklist(ctx, deptChecklist)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["deptChecklist"] = deptChecklist
	response.With200V2(w, "Success", m, platform)
}

//UpdateDeptChecklist :""
func (h *Handler) UpdateDeptChecklist(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	deptChecklist := new(models.DeptChecklist)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&deptChecklist)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if deptChecklist.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateDeptChecklist(ctx, deptChecklist)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["deptChecklist"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDeptChecklist : ""
func (h *Handler) EnableDeptChecklist(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableDeptChecklist(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["deptChecklist"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDeptChecklist : ""
func (h *Handler) DisableDeptChecklist(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableDeptChecklist(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["deptChecklist"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDeptChecklist : ""
func (h *Handler) DeleteDeptChecklist(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteDeptChecklist(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["deptChecklist"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDeptChecklist :""
func (h *Handler) GetSingleDeptChecklist(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	deptChecklist := new(models.RefDeptChecklist)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	deptChecklist, err := h.Service.GetSingleDeptChecklist(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["deptChecklist"] = deptChecklist
	response.With200V2(w, "Success", m, platform)
}

//FilterDeptChecklist : ""
func (h *Handler) FilterDeptChecklist(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var deptChecklist *models.DeptChecklistFilter
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
	err := json.NewDecoder(r.Body).Decode(&deptChecklist)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var deptChecklists []models.RefDeptChecklist
	log.Println(pagination)
	deptChecklists, err = h.Service.FilterDeptChecklist(ctx, deptChecklist, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(deptChecklists) > 0 {
		m["deptChecklist"] = deptChecklists
	} else {
		res := make([]models.DeptChecklist, 0)
		m["deptChecklist"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
