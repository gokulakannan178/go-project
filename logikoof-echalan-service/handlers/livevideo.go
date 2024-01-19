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

//SaveLiveVideo : ""
func (h *Handler) SaveLiveVideo(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	liveVideo := new(models.LiveVideo)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&liveVideo)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveLiveVideo(ctx, liveVideo)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["liveVideo"] = liveVideo
	response.With200V2(w, "Success", m, platform)
}

//UpdateLiveVideo :""
func (h *Handler) UpdateLiveVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	liveVideo := new(models.LiveVideo)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&liveVideo)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if liveVideo.UniqueID == "" {
		response.With400V2(w, "RegNo is missing", platform)
	}
	err = h.Service.UpdateLiveVideo(ctx, liveVideo)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["liveVideo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableLiveVideo : ""
func (h *Handler) EnableLiveVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableLiveVideo(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["liveVideo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableLiveVideo : ""
func (h *Handler) DisableLiveVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableLiveVideo(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["liveVideo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteLiveVideo : ""
func (h *Handler) DeleteLiveVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteLiveVideo(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["liveVideo"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleLiveVideo :""
func (h *Handler) GetSingleLiveVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	liveVideo := new(models.RefLiveVideo)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	liveVideo, err := h.Service.GetSingleLiveVideo(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["liveVideo"] = liveVideo
	response.With200V2(w, "Success", m, platform)
}

//FilterLiveVideo : ""
func (h *Handler) FilterLiveVideo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var liveVideo *models.LiveVideoFilter
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
	err := json.NewDecoder(r.Body).Decode(&liveVideo)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var liveVideos []models.RefLiveVideo
	log.Println(pagination)
	liveVideos, err = h.Service.FilterLiveVideo(ctx, liveVideo, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(liveVideos) > 0 {
		m["liveVideo"] = liveVideos
	} else {
		res := make([]models.LiveVideo, 0)
		m["liveVideo"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
