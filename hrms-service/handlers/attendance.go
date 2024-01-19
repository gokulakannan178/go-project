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

// SaveAttendance : ""
func (h *Handler) SaveAttendance(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	tk := new(models.Attendance)
	err := json.NewDecoder(r.Body).Decode(&tk)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err = h.Service.SaveAttendance(ctx, tk)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveAttendanceWithEditEmployee(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	attandeace := new(models.Attendance)
	err := json.NewDecoder(r.Body).Decode(&attandeace)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if attandeace.EmployeeId == "" {
		response.With400V2(w, "EmployeeId missing", platform)
	}
	if attandeace.Date == nil {
		response.With400V2(w, "date missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err = h.Service.SaveAttendanceWithEditEmployee(ctx, attandeace)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = attandeace
	response.With200V2(w, "Success", m, platform)
}

// GetSingleAttendance : ""
func (h *Handler) GetSingleAttendance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefAttendance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleAttendance(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateAttendance : ""
func (h *Handler) UpdateAttendance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	attendance := new(models.Attendance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&attendance)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if attendance.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAttendance(ctx, attendance)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Attendance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableAttendance : ""
func (h *Handler) EnableAttendance(w http.ResponseWriter, r *http.Request) {
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
	err := h.Service.EnableAttendance(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableAttendance : ""
func (h *Handler) DisableAttendance(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if ID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableAttendance(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// FilterAttendance : ""
func (h *Handler) FilterAttendance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterAttendance
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

	var fts []models.RefAttendance
	log.Println(pagination)

	fts, err = h.Service.FilterAttendance(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["attendance"] = fts
	} else {
		res := make([]models.Attendance, 0)
		m["attendance"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// ClockinAttendance : ""
func (h *Handler) ClockinAttendance(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	tk := new(models.Attendance)
	err := json.NewDecoder(r.Body).Decode(&tk)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err = h.Service.ClockinAttendance(ctx, tk)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//ClockoutAttendance : ""
func (h *Handler) ClockoutAttendance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	attendance := new(models.ClockoutAttendance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&attendance)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if attendance.EmployeeId == "" {
		response.With400V2(w, "Employee id is missing", platform)
	}
	err = h.Service.ClockoutAttendance(ctx, attendance)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Attendance"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// EmployeeAttendanceTodayStatus : ""
func (h *Handler) EmployeeAttendanceTodayStatus(w http.ResponseWriter, r *http.Request) {

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

	task := new(models.EmployeeAttendanceTodayStatus)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.EmployeeAttendanceTodayStatus(ctx, EmployeeID, UniqueID)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	// if task == nil {
	// 	m["AttendanceTodayStatus"] = nil
	// 	response.With200V2(w, "NOT FOUND", m, platform)
	// 	return
	// }
	//fmt.Println("task", task)

	m["AttendanceTodayStatus"] = task
	response.With200V2(w, "Success", m, platform)
}

// AttendanceEmployeeLeaveRequest : ""
func (h *Handler) AttendanceEmployeeLeaveRequest(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	attendance := new(models.Attendance)
	err := json.NewDecoder(r.Body).Decode(&attendance)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if attendance.EmployeeId == "" {
		response.With400V2(w, "Employee id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err = h.Service.AttendanceEmployeeLeaveRequest(ctx, attendance)

	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// // AttendanceEmployeeLeaveApprove : ""
// func (h *Handler) AttendanceEmployeeLeaveApprove(w http.ResponseWriter, r *http.Request) {
// 	platform := r.URL.Query().Get("platform")
// 	attendance := new(models.Attendance)
// 	err := json.NewDecoder(r.Body).Decode(&attendance)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	err = h.Service.AttendanceEmployeeLeaveApprove(ctx, attendance)
// 	if attendance.EmployeeId == "" {
// 		response.With400V2(w, "Employee id is missing", platform)
// 	}
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["attendance"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }

// // AttendanceEmployeeLeaveReject : ""
// func (h *Handler) AttendanceEmployeeLeaveReject(w http.ResponseWriter, r *http.Request) {
// 	platform := r.URL.Query().Get("platform")
// 	attendance := new(models.Attendance)
// 	err := json.NewDecoder(r.Body).Decode(&attendance)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	err = h.Service.AttendanceEmployeeLeaveReject(ctx, attendance)
// 	if attendance.EmployeeId == "" {
// 		response.With400V2(w, "Employee id is missing", platform)
// 	}
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["attendance"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }
func (h *Handler) DayWiseAttendanceReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.DayWiseAttendanceReportFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	resType := r.URL.Query().Get("resType")
	if resType == "reportexcel" {
		file, err := h.Service.DayWiseAttendanceReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=EmployeeDaywiseReport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	//var attendance models.DayWiseAttendanceReport
	attendance := new(models.DayWiseAttendanceReport)
	attendance, err = h.Service.DayWiseAttendanceReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = attendance
	response.With200V2(w, "Success", m, platform)
}

// AttendanceEmployeeStatistics : ""
func (h *Handler) AttendanceEmployeeStatistics(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//UniqueID := r.URL.Query().Get("id")
	EmployeeID := r.URL.Query().Get("Employeeid")
	if EmployeeID == "" {
		response.With400V2(w, "Employee id is missing", platform)
		return
	}
	task := new(models.AttendanceEmployeeStatistics)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.AttendanceEmployeeStatistics(ctx, EmployeeID)
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
func (h *Handler) TodayEmployessLeave(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//UniqueID := r.URL.Query().Get("id")
	EmployeeID := r.URL.Query().Get("orgId")
	// if EmployeeID == "" {
	// 	response.With400V2(w, "Employee id is missing", platform)
	// 	return
	// }

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.TodayEmployessLeave(ctx, EmployeeID)
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

	m["TodayEmployessLeave"] = task
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) EmployeeAttendanceApprove(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var attendance *models.EmployeeAttendanceApprove
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&attendance)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.EmployeeAttendanceApprove(ctx, attendance)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) AllEmployeeAttendanceApprove(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var attendance *models.AllEmployeeAttendanceApprove
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&attendance)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.AllEmployeeAttendanceApprove(ctx, attendance)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) AllEmployeeAttendanceReject(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var attendance *models.AllEmployeeAttendanceApprove
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&attendance)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.AllEmployeeAttendanceReject(ctx, attendance)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetTodayEmployeeTimeOff(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//	task := new(models.RefAttendance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetTodayEmployeeTimeOff(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = task
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) EmployeeAttendanceRejected(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var attendance *models.EmployeeAttendanceApprove
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&attendance)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.EmployeeAttendanceRejected(ctx, attendance)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetTodayEmployeePunchin(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//	task := new(models.RefAttendance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetTodayEmployeePunchin(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = task
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetTodayEmployeeUplannedLeave(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//	task := new(models.RefAttendance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetTodayEmployeeUplannedLeave(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = task
	response.With200V2(w, "Success", m, platform)
}

//Create a approved employee leave and Request pending leaves
func (h *Handler) GetTodayEmployeeAbsent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//	task := new(models.RefAttendance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	var filter *models.DayWiseAttendanceReportFilter
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	TodayEmployeeAbsent, err := h.Service.GetTodayEmployeeAbsent(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["attendance"] = TodayEmployeeAbsent
	response.With200V2(w, "Success", m, platform)
}
