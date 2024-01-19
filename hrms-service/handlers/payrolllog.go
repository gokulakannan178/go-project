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

//SavePayrollLog : ""
func (h *Handler) SavePayrollLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	payrollLog := new(models.PayrollLog)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&payrollLog)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SavePayrollLog(ctx, payrollLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PayrollLog"] = payrollLog
	response.With200V2(w, "Success", m, platform)
}

//UpdatePayrollLog :""
func (h *Handler) UpdatePayrollLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	payrollLog := new(models.PayrollLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&payrollLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if payrollLog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePayrollLog(ctx, payrollLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePayrollLog : ""
func (h *Handler) EnablePayrollLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnablePayrollLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisablePayrollLog : ""
func (h *Handler) DisablePayrollLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisablePayrollLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePayrollLog : ""
func (h *Handler) DeletePayrollLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeletePayrollLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePayrollLog :""
func (h *Handler) GetSinglePayrollLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	payrollLog := new(models.RefPayrollLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	payrollLog, err := h.Service.GetSinglePayrollLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = payrollLog
	response.With200V2(w, "Success", m, platform)
}

//FilterPayrollLog : ""
func (h *Handler) FilterPayrollLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var payrollLog *models.FilterPayrollLog
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
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
	err := json.NewDecoder(r.Body).Decode(&payrollLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var payrollLogs []models.RefPayrollLog
	log.Println(pagination)
	payrollLogs, err = h.Service.FilterPayrollLog(ctx, payrollLog, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payrollLogs) > 0 {
		m["data"] = payrollLogs
	} else {
		res := make([]models.PayrollLog, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetEmployeeSalarySlip(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var payrollLog *models.EmployeeSalarySlip
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&payrollLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	resType := r.URL.Query().Get("resType")
	if resType == "reportexcel" {
		file, err := h.Service.EmployeePayrollExcel(ctx, payrollLog)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=SalarySlip.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	if resType == "pdf" {
		data, err := h.Service.EmployeePayrollPdf(ctx, payrollLog)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				response.With500mV2(w, "failed no data for this id", platform)
				return
			}
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=payslip.pdf")
		w.Write(data)
		return
	}
	if resType == "pdfv2" {
		data, err := h.Service.EmployeePayrollPdfV2(ctx, payrollLog)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				response.With500mV2(w, "failed no data for this id", platform)
				return
			}
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		m := make(map[string]interface{})
		m["data"] = data
		response.With200V2(w, "Success", m, platform)
		return
	}
	var payrollLogs *models.RefPayrollLog

	payrollLogs, err = h.Service.GetEmployeeSalarySlip(ctx, payrollLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = payrollLogs
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) PayrollLogList(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var payrollLog *models.FilterPayrollLog
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
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
	err := json.NewDecoder(r.Body).Decode(&payrollLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var payrollLogs []models.RefPayrollLog
	log.Println(pagination)
	payrollLogs, err = h.Service.PayrollLogList(ctx, payrollLog, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payrollLogs) > 0 {
		m["data"] = payrollLogs
	} else {
		res := make([]models.RefPayrollLog, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
