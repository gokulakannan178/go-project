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

// SaveEmployeeLeave : ""
func (h *Handler) SaveEmployeeLeave(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Elog := new(models.EmployeeLeave)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Elog)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeLeave(ctx, Elog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeave"] = Elog
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeLeave : ""
func (h *Handler) GetSingleEmployeeLeave(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefEmployeeLeave)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleEmployeeLeave(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeave"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeLeave : ""
func (h *Handler) UpdateEmployeeLeave(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	EmployeeLeave := new(models.EmployeeLeave)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&EmployeeLeave)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if EmployeeLeave.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeLeave(ctx, EmployeeLeave)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeave"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeLeave : ""
func (h *Handler) EnableEmployeeLeave(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeLeave(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeave"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeLeave : ""
func (h *Handler) DisableEmployeeLeave(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeLeave(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeave"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteEmployeeLeave : ""
func (h *Handler) DeleteEmployeeLeave(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeLeave(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeave"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeLeave : ""
func (h *Handler) FilterEmployeeLeave(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterEmployeeLeave
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

	var fts []models.RefEmployeeLeave
	log.Println(pagination)
	fts, err = h.Service.FilterEmployeeLeave(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["EmployeeLeave"] = fts
	} else {
		res := make([]models.EmployeeLeave, 0)
		m["EmployeeLeave"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// GetEmployeeLeaveCount : ""
func (h *Handler) GetEmployeeLeaveCount(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Eleave := new(models.EmployeeLeaveCount)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Eleave)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	var fts []models.RefEmployeeLeaveCount
	fts, err = h.Service.GetEmployeeLeaveCount(ctx, Eleave)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeaveCount"] = fts
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateEmployeeLeaveFromTimeOff(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	EmployeeLeave := new(models.UpdateEmployeeLeave)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&EmployeeLeave)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if EmployeeLeave.EmployeeId == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeLeaveFromTimeOff(ctx, EmployeeLeave)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeave"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) EmployeeLeaveList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var employeeDocuments *models.FilterEmployeeLeave
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeDocuments)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	//var employeeDocumentss []models.EmployeeLeaveListV2
	//log.Println(pagination)
	employeeDocumentss, err := h.Service.EmployeeleaveList(ctx, employeeDocuments)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(employeeDocumentss) > 0 {
		m["EmployeeLeaveList"] = employeeDocumentss
	} else {
		res := make([]models.EmployeeDocuments, 0)
		m["EmployeeLeaveList"] = res
	}
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) EmployeeLeaveListV2(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var employeeDocuments *models.FilterEmployeeLeaveList
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeDocuments)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	//var employeeDocumentss []models.EmployeeLeaveListV2
	//log.Println(pagination)
	employeeDocumentss, err := h.Service.EmployeeLeaveListV2(ctx, employeeDocuments)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeave"] = employeeDocumentss
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetAllEmployeeLeaveList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var employeeDocuments *models.FilterEmployeeLeaveList
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&employeeDocuments)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	//var employeeDocumentss []models.EmployeeLeaveListV2
	//log.Println(pagination)
	employeeDocumentss, err := h.Service.GetAllEmployeeLeaveList(ctx, employeeDocuments)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeave"] = employeeDocumentss
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateEmployeeLeaveWithEmployeeId(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	EmployeeLeave := new(models.EmployeeLeaveLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&EmployeeLeave)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if EmployeeLeave.LeaveType == "" {
		response.With400V2(w, "Leaveid is missing", platform)
	}
	if EmployeeLeave.EmployeeId == "" {
		response.With400V2(w, "Employeeid is missing", platform)
	}
	err = h.Service.UpdateEmployeeLeaveWithEmployeeId(ctx, EmployeeLeave)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLeave"] = "success"
	response.With200V2(w, "Success", m, platform)
}
