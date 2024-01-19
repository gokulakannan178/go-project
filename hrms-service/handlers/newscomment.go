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

// SaveNewsComment : ""
func (h *Handler) SaveNewsComment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	newscomment := new(models.NewsComment)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&newscomment)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveNewsComment(ctx, newscomment)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsComment"] = newscomment
	response.With200V2(w, "Success", m, platform)
}

// GetSingleNewsComment : ""
func (h *Handler) GetSingleNewsComment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	newscomment := new(models.RefNewsComment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	newscomment, err := h.Service.GetSingleNewsComment(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsComment"] = newscomment
	response.With200V2(w, "Success", m, platform)
}

//UpdateNewsComment : ""
func (h *Handler) UpdateNewsComment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	newscomment := new(models.NewsComment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&newscomment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if newscomment.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateNewsComment(ctx, newscomment)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsComment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableNewsComment : ""
func (h *Handler) EnableNewsComment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableNewsComment(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsComment"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableNewsComment : ""
func (h *Handler) DisableNewsComment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableNewsComment(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsComment"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteNewsComment : ""
func (h *Handler) DeleteNewsComment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteNewsComment(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["NewsComment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterNewsComment : ""
func (h *Handler) FilterNewsComment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterNewscomment *models.FilterNewsComment
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
	err := json.NewDecoder(r.Body).Decode(&filterNewscomment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterNewsComments []models.RefNewsComment
	log.Println(pagination)
	filterNewsComments, err = h.Service.FilterNewsComment(ctx, filterNewscomment, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterNewsComments) > 0 {
		m["NewsComment"] = filterNewsComments
	} else {
		res := make([]models.NewsComment, 0)
		m["NewsComment"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
