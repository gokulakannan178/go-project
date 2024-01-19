package handlers

import (
	"encoding/json"
	"log"
	"logikoof-echalan-service/app"
	"logikoof-echalan-service/models"
	"logikoof-echalan-service/response"
	"net/http"
	"strconv"
)

//SaveOffenceVideo : ""
func (h *Handler) SaveOffenceVideo(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	offenceVideo := new(models.OffenceVideo)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&offenceVideo)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveOffenceVideo(ctx, offenceVideo)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceVideo"] = offenceVideo
	response.With200V2(w, "Success", m, platform)
}

//UpdateOffenceVideo :""
func (h *Handler) UpdateOffenceVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	offenceVideo := new(models.OffenceVideo)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&offenceVideo)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if offenceVideo.UniqueID == "" {
		response.With400V2(w, "RegNo is missing", platform)
	}
	err = h.Service.UpdateOffenceVideo(ctx, offenceVideo)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceVideo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableOffenceVideo : ""
func (h *Handler) EnableOffenceVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableOffenceVideo(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceVideo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableOffenceVideo : ""
func (h *Handler) DisableOffenceVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableOffenceVideo(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceVideo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteOffenceVideo : ""
func (h *Handler) DeleteOffenceVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteOffenceVideo(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceVideo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleOffenceVideo :""
func (h *Handler) GetSingleOffenceVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	offenceVideo := new(models.RefOffenceVideo)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	offenceVideo, err := h.Service.GetSingleOffenceVideo(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["offenceVideo"] = offenceVideo
	response.With200V2(w, "Success", m, platform)
}

//FilterOffenceVideo : ""
func (h *Handler) FilterOffenceVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var offenceVideo *models.OffenceVideoFilter
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
	err := json.NewDecoder(r.Body).Decode(&offenceVideo)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var offenceVideos []models.RefOffenceVideo
	log.Println(pagination)
	offenceVideos, err = h.Service.FilterOffenceVideo(ctx, offenceVideo, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(offenceVideos) > 0 {
		m["offenceVideo"] = offenceVideos
	} else {
		res := make([]models.OffenceVideo, 0)
		m["offenceVideo"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
