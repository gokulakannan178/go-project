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

// SaveNewsLike : ""
func (h *Handler) SaveNewsLike(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	newslike := new(models.NewsLike)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&newslike)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveNewsLike(ctx, newslike)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsLike"] = newslike
	response.With200V2(w, "Success", m, platform)
}

// GetSingleNewsLike : ""
func (h *Handler) GetSingleNewsLike(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	newslike := new(models.RefNewsLike)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	newslike, err := h.Service.GetSingleNewsLike(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsLike"] = newslike
	response.With200V2(w, "Success", m, platform)
}

//UpdateNewsLike : ""
func (h *Handler) UpdateNewsLike(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	newslike := new(models.NewsLike)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&newslike)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if newslike.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateNewsLike(ctx, newslike)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsLike"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableNewsLike : ""
func (h *Handler) EnableNewsLike(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableNewsLike(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsLike"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableNewsLike : ""
func (h *Handler) DisableNewsLike(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableNewsLike(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["News"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteNewsLike : ""
func (h *Handler) DeleteNewsLike(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteNewsLike(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsLike"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterNewsLike : ""
func (h *Handler) FilterNewsLike(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterNewsLike *models.FilterNewsLike
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
	err := json.NewDecoder(r.Body).Decode(&filterNewsLike)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterNewsLikes []models.RefNewsLike
	log.Println(pagination)
	filterNewsLikes, err = h.Service.FilterNewsLike(ctx, filterNewsLike, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterNewsLikes) > 0 {
		m["NewsLike"] = filterNewsLikes
	} else {
		res := make([]models.NewsLike, 0)
		m["NewsLike"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
