package handlers

import (
	"encoding/json"
	"fmt"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

// SaveAttendanceLog : ""
func (h *Handler) SaveAttendanceLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	tk := new(models.AttendanceLog)
	err := json.NewDecoder(r.Body).Decode(&tk)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err = h.Service.SaveAttendanceLog(ctx, tk)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AttendanceLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleAttendanceLog : ""
func (h *Handler) GetSingleAttendanceLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefAttendanceLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleAttendanceLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AttendanceLog"] = task
	response.With200V2(w, "Success", m, platform)
}

// GetSingleAttendanceLoglast : ""
func (h *Handler) GetSingleAttendanceLoglast(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	EmployeeID := r.URL.Query().Get("Employeeid")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}
	if EmployeeID == "" {
		response.With400V2(w, "Employee id is missing", platform)
		return
	}

	task := new(models.RefAttendanceLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleAttendanceLoglast(ctx, EmployeeID, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	//fmt.Println("task", task)

	m["attendanceLogRecent"] = task
	response.With200V2(w, "Success", m, platform)
}

// AttendanceEmployeeTodayStatus : ""
func (h *Handler) AttendanceEmployeeTodayStatus(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//	UniqueID := r.URL.Query().Get("id")
	EmployeeID := r.URL.Query().Get("Employeeid")

	// if UniqueID == "" {
	// 	response.With400V2(w, "id is missing", platform)
	// 	return
	// }
	if EmployeeID == "" {
		response.With400V2(w, "Employee id is missing", platform)
		return
	}

	task := new(models.AttendanceEmployeeTodayStatus)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.AttendanceEmployeeTodayStatus(ctx, EmployeeID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	//fmt.Println("task", task)

	m["AttendanceLog"] = task
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) AttendanceEmployeeTodayLogs(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//	UniqueID := r.URL.Query().Get("id")
	EmployeeID := r.URL.Query().Get("Employeeid")

	if EmployeeID == "" {
		response.With400V2(w, "Employee id is missing", platform)
		return
	}

	//	task := new(models.AttendanceLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.AttendanceEmployeeTodayLogs(ctx, EmployeeID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	//fmt.Println("task", task)
	m["AttendanceLog"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateAttendanceLog : ""
func (h *Handler) UpdateAttendanceLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	attendancelog := new(models.AttendanceLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&attendancelog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if attendancelog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAttendanceLog(ctx, attendancelog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AttendanceLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableAttendanceLog : ""
func (h *Handler) EnableAttendanceLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	fmt.Println(r)
	fmt.Println(r.URL)
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL.Query().Get("platform"))

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableAttendanceLog(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AttendanceLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableAttendanceLog : ""
func (h *Handler) DisableAttendanceLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableAttendanceLog(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AttendanceLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteAttendanceLog : ""
func (h *Handler) DeleteAttendanceLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	fmt.Println(r)
	fmt.Println(r.URL)
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL.Query().Get("platform"))

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DeleteAttendanceLog(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AttendanceLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// FilterAttendanceLog : ""
func (h *Handler) FilterAttendanceLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterAttendanceLog
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

	var fts []models.RefAttendanceLog
	log.Println(pagination)
	fts, err = h.Service.FilterAttendanceLog(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["AttendanceLog"] = fts
	} else {
		res := make([]models.AttendanceLog, 0)
		m["AttendanceLog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
