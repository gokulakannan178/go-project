package handlers

import (
	"encoding/json"
	"log"

	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

// SaveCitizenGraviansLog : ""
func (h *Handler) SaveCitizenGraviansLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	citizengravianslog := new(models.CitizenGraviansLog)
	err := json.NewDecoder(r.Body).Decode(&citizengravianslog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveCitizenGraviansLog(ctx, citizengravianslog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizengravianslog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleCitizenGraviansLog : ""
func (h *Handler) GetSingleCitizenGraviansLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	citizengravianslog := new(models.RefCitizenGraviansLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	citizengravianslog, err := h.Service.GetSingleCitizenGraviansLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizengravianslog"] = citizengravianslog
	response.With200V2(w, "Success", m, platform)
}

// UpdateCitizenGraviansLog : ""
func (h *Handler) UpdateCitizenGraviansLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	citizengravianslog := new(models.CitizenGraviansLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&citizengravianslog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if citizengravianslog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateCitizenGraviansLog(ctx, citizengravianslog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizengravianslog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableCitizenGraviansLog : ""
func (h *Handler) EnableCitizenGraviansLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableCitizenGraviansLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizengravianslog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableCitizenGraviansLog : ""
func (h *Handler) DisableCitizenGraviansLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableCitizenGraviansLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizengravianslog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteCitizenGraviansLog : ""
func (h *Handler) DeleteCitizenGraviansLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteCitizenGraviansLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["citizengravianslog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterCitizenGraviansLog : ""
func (h *Handler) FilterCitizenGraviansLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.CitizenGraviansLogFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var citizengravianslogs []models.RefCitizenGraviansLog
	log.Println(pagination)
	citizengravianslogs, err = h.Service.FilterCitizenGraviansLog(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(citizengravianslogs) > 0 {
		m["citizengravianslog"] = citizengravianslogs
	} else {
		res := make([]models.CitizenGraviansLog, 0)
		m["citizengravianslog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
