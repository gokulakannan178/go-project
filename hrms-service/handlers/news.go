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

// SaveNews : ""
func (h *Handler) SaveNews(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	news := new(models.News)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&news)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveNews(ctx, news)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["News"] = news
	response.With200V2(w, "Success", m, platform)
}

// GetSingleNews : ""
func (h *Handler) GetSingleNews(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	news := new(models.RefNews)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	news, err := h.Service.GetSingleNews(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["News"] = news
	response.With200V2(w, "Success", m, platform)
}

//UpdateNews : ""
func (h *Handler) UpdateNews(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	news := new(models.News)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&news)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if news.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateNews(ctx, news)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["News"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableNews : ""
func (h *Handler) EnableNews(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableNews(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["News"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableNews : ""
func (h *Handler) DisableNews(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableNews(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["News"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteNews : ""
func (h *Handler) DeleteNews(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteNews(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["News"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterNews : ""
func (h *Handler) FilterNews(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterNews *models.FilterNews
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
	err := json.NewDecoder(r.Body).Decode(&filterNews)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterNewss []models.RefNews
	log.Println(pagination)
	filterNewss, err = h.Service.FilterNews(ctx, filterNews, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterNewss) > 0 {
		m["News"] = filterNewss
	} else {
		res := make([]models.News, 0)
		m["News"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) PublishedNews(w http.ResponseWriter, r *http.Request) {

	news := new(models.News)
	platform := r.URL.Query().Get("platform")
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&news)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	//var ctx *models.Context
	//ctx = app.GetApp(r.Context(), h.Service.Daos)

	err = h.Service.PublishedNews(ctx, news)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["News"] = "success"
	response.With200V2(w, "Success", m, platform)
}
