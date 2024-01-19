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

//SaveWorkSchedule : ""
func (h *Handler) SaveWorkSchedule(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	workSchedule := new(models.WorkSchedule)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&workSchedule)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveWorkSchedule(ctx, workSchedule)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["workSchedule"] = workSchedule
	response.With200V2(w, "Success", m, platform)
}

//UpdateWorkSchedule :""
func (h *Handler) UpdateWorkSchedule(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	workSchedule := new(models.WorkSchedule)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&workSchedule)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if workSchedule.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateWorkSchedule(ctx, workSchedule)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["workSchedule"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableWorkSchedule : ""
func (h *Handler) EnableWorkSchedule(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableWorkSchedule(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["workSchedule"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableWorkSchedule : ""
func (h *Handler) DisableWorkSchedule(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableWorkSchedule(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["workSchedule"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteWorkSchedule : ""
func (h *Handler) DeleteWorkSchedule(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteWorkSchedule(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["workSchedule"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleWorkSchedule :""
func (h *Handler) GetSingleWorkSchedule(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	workSchedule := new(models.RefWorkSchedule)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	workSchedule, err := h.Service.GetSingleWorkSchedule(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["workSchedule"] = workSchedule
	response.With200V2(w, "Success", m, platform)
}

//FilterWorkSchedule : ""
func (h *Handler) FilterWorkSchedule(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.WorkScheduleFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var workSchedules []models.RefWorkSchedule
	log.Println(pagination)
	workSchedules, err = h.Service.FilterWorkSchedule(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(workSchedules) > 0 {
		m["workSchedule"] = workSchedules
	} else {
		res := make([]models.WorkSchedule, 0)
		m["workSchedule"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
