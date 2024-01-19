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

// SaveSalaryConfigLog : ""
func (h *Handler) SaveSalaryConfigLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	salaryConfigLog := new(models.SalaryConfigLog)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&salaryConfigLog)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveSalaryConfigLog(ctx, salaryConfigLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfigLog"] = salaryConfigLog
	response.With200V2(w, "Success", m, platform)
}

// GetSingleSalaryConfigLog : ""
func (h *Handler) GetSingleSalaryConfigLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	salaryConfigLog := new(models.RefSalaryConfigLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	salaryConfigLog, err := h.Service.GetSingleSalaryConfigLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfigLog"] = salaryConfigLog
	response.With200V2(w, "Success", m, platform)
}

//UpdateSalaryConfigLog : ""
func (h *Handler) UpdateSalaryConfigLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	salaryConfigLog := new(models.SalaryConfigLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&salaryConfigLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if salaryConfigLog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateSalaryConfigLog(ctx, salaryConfigLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfigLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableSalaryConfigLog : ""
func (h *Handler) EnableSalaryConfigLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableSalaryConfigLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfigLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableSalaryConfigLog : ""
func (h *Handler) DisableSalaryConfigLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableSalaryConfigLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfigLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteSalaryConfigLog : ""
func (h *Handler) DeleteSalaryConfigLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteSalaryConfigLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["SalaryConfigLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterSalaryConfigLog : ""
func (h *Handler) FilterSalaryConfigLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterSalaryConfigLog *models.FilterSalaryConfigLog
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
	err := json.NewDecoder(r.Body).Decode(&filterSalaryConfigLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterSalaryConfigLogs []models.RefSalaryConfigLog
	log.Println(pagination)
	filterSalaryConfigLogs, err = h.Service.FilterSalaryConfigLog(ctx, filterSalaryConfigLog, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterSalaryConfigLogs) > 0 {
		m["SalaryConfigLog"] = filterSalaryConfigLogs
	} else {
		res := make([]models.SalaryConfigLog, 0)
		m["SalaryConfigLog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
