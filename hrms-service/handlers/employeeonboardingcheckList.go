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

// SaveEmployeeOnboardingCheckList : ""
func (h *Handler) SaveEmployeeOnboardingCheckList(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeonboardingchecklist := new(models.EmployeeOnboardingCheckList)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeonboardingchecklist)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeOnboardingCheckList(ctx, employeeonboardingchecklist)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOnboardingCheckList"] = employeeonboardingchecklist
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeOnboardingCheckList : ""
func (h *Handler) GetSingleEmployeeOnboardingCheckList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefEmployeeOnboardingCheckList)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleEmployeeOnboardingCheckList(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOnboardingCheckList"] = task
	response.With200V2(w, "Success", m, platform)
}

// EmployeeOnboardingCheckListFinal : ""
func (h *Handler) EmployeeOnboardingCheckListFinal(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	EmployeeID := r.URL.Query().Get("employeeid")

	if EmployeeID == "" {
		response.With400V2(w, "EmployeeID is missing", platform)
		return
	}

	task := new(models.RefEmployeeOnboardingCheckListv2)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.EmployeeOnboardingCheckListFinal(ctx, EmployeeID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOnboardingCheckList"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeOnboardingCheckList : ""
func (h *Handler) UpdateEmployeeOnboardingCheckList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeonboardingchecklist := new(models.EmployeeOnboardingCheckList)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeonboardingchecklist)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeonboardingchecklist.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeOnboardingCheckList(ctx, employeeonboardingchecklist)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOnboardingCheckList"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeOnboardingCheckList : ""
func (h *Handler) EnableEmployeeOnboardingCheckList(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeOnboardingCheckList(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOnboardingCheckList"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeOnboardingCheckList : ""
func (h *Handler) DisableEmployeeOnboardingCheckList(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeOnboardingCheckList(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOnboardingCheckList"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteEmployeeOnboardingCheckList : ""
func (h *Handler) DeleteEmployeeOnboardingCheckList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeOnboardingCheckList(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOnboardingCheckList"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeOnboardingCheckList : ""
func (h *Handler) FilterEmployeeOnboardingCheckList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterEmployeeOnboardingCheckList
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
	err := json.NewDecoder(r.Body).Decode(&ft)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var fts []models.RefEmployeeOnboardingCheckList
	log.Println(pagination)
	fts, err = h.Service.FilterEmployeeOnboardingCheckList(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["EmployeeOnboardingCheckList"] = fts
	} else {
		res := make([]models.EmployeeOnboardingCheckList, 0)
		m["EmployeeOnboardingCheckList"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
