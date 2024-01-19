package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/constants"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

// SaveEmployeeTimeOff : ""
func (h *Handler) SaveEmployeeTimeOff(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeTimeOff := new(models.EmployeeTimeOff)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeTimeOff)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeTimeOff(ctx, employeeTimeOff)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = employeeTimeOff
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeTimeOff : ""
func (h *Handler) GetSingleEmployeeTimeOff(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	EmployeeTimeOff := new(models.RefEmployeeTimeOff)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	EmployeeTimeOff, err := h.Service.GetSingleEmployeeTimeOff(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = EmployeeTimeOff
	response.With200V2(w, "Success", m, platform)
}

// UpdateEmployeeTimeOff : ""
func (h *Handler) UpdateEmployeeTimeOff(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeTimeOff := new(models.EmployeeTimeOff)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeTimeOff)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeTimeOff.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeTimeOff(ctx, employeeTimeOff)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EmployeeTimeOffCount : ""
func (h *Handler) EmployeeTimeOffCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeTimeOffCount := new(models.EmployeeTimeOffCount)
	//employeeTimeOffRef := new(models.RefEmployeeTimeOffCount)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeTimeOffCount)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeTimeOffCount.EmployeeId == "" {
		response.With400V2(w, "Employee Id is missing", platform)
	}

	timeoffcount, err := h.Service.EmployeeTimeOffCount(ctx, employeeTimeOffCount)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOffCount"] = timeoffcount
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeTimeOff : ""
func (h *Handler) EnableEmployeeTimeOff(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeTimeOff(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeTimeOff : ""
func (h *Handler) DisableEmployeeTimeOff(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeTimeOff(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeeTimeOff : ""
func (h *Handler) DeleteEmployeeTimeOff(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeTimeOff(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeTimeOff : ""
func (h *Handler) FilterEmployeeTimeOff(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeeTimeOff *models.FilterEmployeeTimeOff
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeeTimeOff)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeeTimeOffs []models.RefEmployeeTimeOff
	log.Println(pagination)
	filterEmployeeTimeOffs, err = h.Service.FilterEmployeeTimeOff(ctx, filterEmployeeTimeOff, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeeTimeOffs) > 0 {
		m["EmployeeTimeOff"] = filterEmployeeTimeOffs
	} else {
		res := make([]models.EmployeeTimeOff, 0)
		m["EmployeeTimeOff"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// EmployeeTimeOffRequest : ""
func (h *Handler) EmployeeTimeOffRequest(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeTimeOff := new(models.EmployeeTimeOff)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeTimeOff)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.EmployeeTimeOffRequest(ctx, employeeTimeOff)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = employeeTimeOff
	response.With200V2(w, "Success", m, platform)
}

// EmployeeTimeOffApprove : ""
func (h *Handler) EmployeeTimeOffApprove(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeTimeOff := new(models.ReviewedEmployeeTimeOff)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeTimeOff)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeTimeOff.EmployeeTimeOff == "" {
		response.With400V2(w, "EmployeeTimeOff is missing", platform)
	}
	err = h.Service.EmployeeTimeOffApprove(ctx, employeeTimeOff)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EmployeeTimeOffRevoke : ""
func (h *Handler) EmployeeTimeOffRevoke(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeTimeOff := new(models.ReviewedEmployeeTimeOff)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeTimeOff)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeTimeOff.EmployeeTimeOff == "" {
		response.With400V2(w, "EmployeeTimeOff is missing", platform)
	}
	err = h.Service.EmployeeTimeOffRevoke(ctx, employeeTimeOff)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EmployeeTimeOffReject : ""
func (h *Handler) EmployeeTimeOffReject(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeTimeOff := new(models.ReviewedEmployeeTimeOff)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeTimeOff)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeTimeOff.EmployeeTimeOff == "" {
		response.With400V2(w, "EmployeeTimeOff is missing", platform)
	}
	err = h.Service.EmployeeTimeOffReject(ctx, employeeTimeOff)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) RevokeRequestEmployeeTimeOff(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	if status == constants.EMPLOYEETIMEOFFSTATUSREQUEST {

		err := h.Service.CancelEmployeeTimeOff(ctx, UniqueID)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
	} else if status == constants.EMPLOYEETIMEOFFSTATUSAPPROVE {
		err := h.Service.RevokeRequestEmployeeTimeOff(ctx, UniqueID)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
	} else {
		response.With400V2(w, "status is missing", platform)
		return
	}

	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EmployeeTimeoffDateCheck : ""
func (h *Handler) EmployeeTimeoffDateCheck(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeTimeOff := new(models.DayWiseAttendanceReportFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeTimeOff)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeTimeOff.EmployeeId == "" {
		response.With400V2(w, "Employee is missing", platform)
		return
	}
	if employeeTimeOff.StartDate == nil {
		response.With400V2(w, "Date is missing", platform)
		return
	}
	err = h.Service.EmployeeTimeoffDateCheck(ctx, employeeTimeOff)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeTimeOff"] = "success"
	response.With200V2(w, "Success", m, platform)
}
