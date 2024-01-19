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

// SaveEmployeeOffboardingCheckList : ""
func (h *Handler) SaveEmployeeOffboardingCheckList(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeoffboardingchecklist := new(models.EmployeeOffboardingCheckList)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeoffboardingchecklist)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeOffboardingCheckList(ctx, employeeoffboardingchecklist)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOffboardingCheckList"] = employeeoffboardingchecklist
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeOffboardingCheckList : ""
func (h *Handler) GetSingleEmployeeOffboardingCheckList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefEmployeeOffboardingCheckList)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleEmployeeOffboardingCheckList(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOffboardingCheckList"] = task
	response.With200V2(w, "Success", m, platform)
}

// EmployeeOffboardingCheckListFinal : ""
func (h *Handler) EmployeeOffboardingCheckListFinal(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	EmployeeID := r.URL.Query().Get("employeeid")

	if EmployeeID == "" {
		response.With400V2(w, "EmployeeID is missing", platform)
		return
	}

	task := new(models.RefEmployeeOffboardingCheckListv2)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.EmployeeOffboardingCheckListFinal(ctx, EmployeeID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOffboardingCheckList"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeOffboardingCheckList : ""
func (h *Handler) UpdateEmployeeOffboardingCheckList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeoffboardingchecklist := new(models.EmployeeOffboardingCheckList)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeoffboardingchecklist)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeoffboardingchecklist.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeOffboardingCheckList(ctx, employeeoffboardingchecklist)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOffboardingCheckList"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeOffboardingCheckList : ""
func (h *Handler) EnableEmployeeOffboardingCheckList(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeOffboardingCheckList(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOffboardingCheckList"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeOffboardingCheckList : ""
func (h *Handler) DisableEmployeeOffboardingCheckList(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeOffboardingCheckList(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOffboardingCheckList"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteEmployeeOffboardingCheckList : ""
func (h *Handler) DeleteEmployeeOffboardingCheckList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeOffboardingCheckList(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeOffboardingCheckList"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeOffboardingCheckList : ""
func (h *Handler) FilterEmployeeOffboardingCheckList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterEmployeeOffboardingCheckList
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

	var fts []models.RefEmployeeOffboardingCheckList
	log.Println(pagination)
	fts, err = h.Service.FilterEmployeeOffboardingCheckList(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["EmployeeOffboardingCheckList"] = fts
	} else {
		res := make([]models.EmployeeOffboardingCheckList, 0)
		m["EmployeeOffboardingCheckList"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
