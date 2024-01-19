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

// SaveEmployeeAttendanceCalendar : ""
func (h *Handler) SaveEmployeeAttendanceCalendar(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	employeeAttendanceCalendar := new(models.EmployeeAttendanceCalendar)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&employeeAttendanceCalendar)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeAttendanceCalendar(ctx, employeeAttendanceCalendar)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAttendanceCalendar"] = employeeAttendanceCalendar
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeAttendanceCalendar : ""
func (h *Handler) GetSingleEmployeeAttendanceCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	employeeAttendanceCalendar := new(models.RefEmployeeAttendanceCalendar)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeAttendanceCalendar, err := h.Service.GetSingleEmployeeAttendanceCalendar(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAttendanceCalendar"] = employeeAttendanceCalendar
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeAttendanceCalendar : ""
func (h *Handler) UpdateEmployeeAttendanceCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeAttendanceCalendar := new(models.EmployeeAttendanceCalendar)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&employeeAttendanceCalendar)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if employeeAttendanceCalendar.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeAttendanceCalendar(ctx, employeeAttendanceCalendar)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAttendanceCalendar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeAttendanceCalendar : ""
func (h *Handler) EnableEmployeeAttendanceCalendar(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeAttendanceCalendar(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAttendanceCalendar"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeAttendanceCalendar : ""
func (h *Handler) DisableEmployeeAttendanceCalendar(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeAttendanceCalendar(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAttendanceCalendar"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteEmployeeAttendanceCalendar : ""
func (h *Handler) DeleteEmployeeAttendanceCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeAttendanceCalendar(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAttendanceCalendar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeAttendanceCalendar : ""
func (h *Handler) FilterEmployeeAttendanceCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterEmployeeAttendanceCalendar *models.FilterEmployeeAttendanceCalendar
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
	err := json.NewDecoder(r.Body).Decode(&filterEmployeeAttendanceCalendar)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterEmployeeAttendanceCalendars []models.RefEmployeeAttendanceCalendar
	log.Println(pagination)
	filterEmployeeAttendanceCalendars, err = h.Service.FilterEmployeeAttendanceCalendar(ctx, filterEmployeeAttendanceCalendar, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterEmployeeAttendanceCalendars) > 0 {
		m["EmployeeAttendanceCalendar"] = filterEmployeeAttendanceCalendars
	} else {
		res := make([]models.EmployeeAttendanceCalendar, 0)
		m["EmployeeAttendanceCalendar"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleEmployeeAttendanceCalendarWithCurrentMonth(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	employeeAttendanceCalendar := new(models.RefEmployeeAttendanceCalendar)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	employeeAttendanceCalendar, err := h.Service.GetSingleEmployeeAttendanceCalendarWithCurrentMonth(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeAttendanceCalendar"] = employeeAttendanceCalendar
	response.With200V2(w, "Success", m, platform)
}
