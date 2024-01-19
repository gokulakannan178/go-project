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

// SaveEmployeeLeaveLog : ""
func (h *Handler) SaveEmployeeLeaveLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeLeaveLog := new(models.EmployeeLeaveLog)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeLeaveLog)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeLeaveLog(ctx, employeeLeaveLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeaveLog"] = employeeLeaveLog
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeLeaveLog : ""
func (h *Handler) GetSingleEmployeeLeaveLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeeLeaveLog := new(models.RefEmployeeLeaveLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeLeaveLog, err := h.Service.GetSingleEmployeeLeaveLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeaveLog"] = employeeLeaveLog
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeLeaveLog : ""
func (h *Handler) UpdateEmployeeLeaveLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeLeaveLog := new(models.EmployeeLeaveLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeLeaveLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeLeaveLog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeLeaveLog(ctx, employeeLeaveLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeaveLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EmployeeLeaveLogCount : ""
func (h *Handler) EmployeeLeaveLogCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeLeaveLogCount := new(models.EmployeeLeaveLogCount)
	//EmployeeLeaveLogRef := new(models.RefEmployeeLeaveLogCount)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeLeaveLogCount)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeLeaveLogCount.EmployeeId == "" {
		response.With400V2(w, "Employee Id is missing", platform)
	}

	leavecount, err := h.Service.EmployeeLeaveLogCount(ctx, employeeLeaveLogCount)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeaveLogCount"] = leavecount
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeLeaveLog : ""
func (h *Handler) EnableEmployeeLeaveLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeLeaveLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeaveLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeLeaveLog : ""
func (h *Handler) DisableEmployeeLeaveLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeLeaveLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeaveLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeeLeaveLog : ""
func (h *Handler) DeleteEmployeeLeaveLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeLeaveLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeaveLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeLeaveLog : ""
func (h *Handler) FilterEmployeeLeaveLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeeLeaveLog *models.FilterEmployeeLeaveLog
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeeLeaveLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeeLeaveLogs []models.RefEmployeeLeaveLog
	log.Println(pagination)
	filterEmployeeLeaveLogs, err = h.Service.FilterEmployeeLeaveLog(ctx, filterEmployeeLeaveLog, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeeLeaveLogs) > 0 {
		m["EmployeeLeaveLog"] = filterEmployeeLeaveLogs
	} else {
		res := make([]models.EmployeeLeaveLog, 0)
		m["EmployeeLeaveLog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
