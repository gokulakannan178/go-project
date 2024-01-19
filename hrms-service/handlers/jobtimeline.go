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

//SaveJobTimeline : ""
func (h *Handler) SaveJobTimeline(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	jobTimeline := new(models.JobTimeline)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&jobTimeline)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveJobTimeline(ctx, jobTimeline)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["jobTimeline"] = jobTimeline
	response.With200V2(w, "Success", m, platform)
}

//UpdateJobTimeline :""
func (h *Handler) UpdateJobTimeline(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	jobTimeline := new(models.JobTimeline)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&jobTimeline)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if jobTimeline.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateJobTimeline(ctx, jobTimeline)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["jobTimeline"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableJobTimeline : ""
func (h *Handler) EnableJobTimeline(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableJobTimeline(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["jobTimeline"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableJobTimeline : ""
func (h *Handler) DisableJobTimeline(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableJobTimeline(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["jobTimeline"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteJobTimeline : ""
func (h *Handler) DeleteJobTimeline(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteJobTimeline(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["jobTimeline"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleJobTimeline :""
func (h *Handler) GetSingleJobTimeline(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	jobTimeline := new(models.RefJobTimeline)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	jobTimeline, err := h.Service.GetSingleJobTimeline(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["jobTimeline"] = jobTimeline
	response.With200V2(w, "Success", m, platform)
}

//FilterJobTimeline : ""
func (h *Handler) FilterJobTimeline(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.JobTimelineFilter
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

	var JobTimelines []models.RefJobTimeline
	log.Println(pagination)
	JobTimelines, err = h.Service.FilterJobTimeline(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(JobTimelines) > 0 {
		m["jobTimeline"] = JobTimelines
	} else {
		res := make([]models.JobTimeline, 0)
		m["jobTimeline"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
