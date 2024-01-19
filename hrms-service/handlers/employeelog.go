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

// SaveEmployeeLog : ""
func (h *Handler) SaveEmployeeLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Elog := new(models.EmployeeLog)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Elog)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveEmployeeLog(ctx, Elog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLog"] = Elog
	response.With200V2(w, "Success", m, platform)
}

// GetSingleEmployeeLog : ""
func (h *Handler) GetSingleEmployeeLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefEmployeeLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleEmployeeLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLog"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateEmployeeLog : ""
func (h *Handler) UpdateEmployeeLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	EmployeeLog := new(models.EmployeeLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&EmployeeLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if EmployeeLog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateEmployeeLog(ctx, EmployeeLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableEmployeeLog : ""
func (h *Handler) EnableEmployeeLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableEmployeeLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableEmployeeLog : ""
func (h *Handler) DisableEmployeeLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableEmployeeLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteEmployeeLog : ""
func (h *Handler) DeleteEmployeeLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteEmployeeLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["EmployeeLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterEmployeeLog : ""
func (h *Handler) FilterEmployeeLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterEmployeeLog
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

	var fts []models.RefEmployeeLog
	log.Println(pagination)
	fts, err = h.Service.FilterEmployeeLog(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["EmployeeLog"] = fts
	} else {
		res := make([]models.EmployeeLog, 0)
		m["EmployeeLog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
