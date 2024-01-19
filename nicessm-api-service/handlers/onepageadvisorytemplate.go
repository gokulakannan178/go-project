package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveOnePageAdvisoryTemplate : ""
func (h *Handler) SaveOnePageAdvisoryTemplate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	OnePageAdvisoryTemplate := new(models.OnePageAdvisoryTemplate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&OnePageAdvisoryTemplate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveOnePageAdvisoryTemplate(ctx, OnePageAdvisoryTemplate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisoryTemplate"] = OnePageAdvisoryTemplate
	response.With200V2(w, "Success", m, platform)
}

//UpdateOnePageAdvisoryTemplate :""
func (h *Handler) UpdateOnePageAdvisoryTemplate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	OnePageAdvisoryTemplate := new(models.OnePageAdvisoryTemplate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&OnePageAdvisoryTemplate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if OnePageAdvisoryTemplate.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateOnePageAdvisoryTemplate(ctx, OnePageAdvisoryTemplate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisoryTemplate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableOnePageAdvisoryTemplate : ""
func (h *Handler) EnableOnePageAdvisoryTemplate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableOnePageAdvisoryTemplate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisoryTemplate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableOnePageAdvisoryTemplate : ""
func (h *Handler) DisableOnePageAdvisoryTemplate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableOnePageAdvisoryTemplate(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisoryTemplate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteOnePageAdvisoryTemplate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.OnePageAdvisoryTemplate)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteOnePageAdvisoryTemplate(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisoryTemplate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleOnePageAdvisoryTemplate :""
func (h *Handler) GetSingleOnePageAdvisoryTemplate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	OnePageAdvisoryTemplate := new(models.RefOnePageAdvisoryTemplate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	OnePageAdvisoryTemplate, err := h.Service.GetSingleOnePageAdvisoryTemplate(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisoryTemplate"] = OnePageAdvisoryTemplate
	response.With200V2(w, "Success", m, platform)
}

//FilterOnePageAdvisoryTemplate : ""
func (h *Handler) FilterOnePageAdvisoryTemplate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var OnePageAdvisoryTemplate *models.OnePageAdvisoryTemplateFilter
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
	err := json.NewDecoder(r.Body).Decode(&OnePageAdvisoryTemplate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var OnePageAdvisoryTemplates []models.RefOnePageAdvisoryTemplate
	log.Println(pagination)
	OnePageAdvisoryTemplates, err = h.Service.FilterOnePageAdvisoryTemplate(ctx, OnePageAdvisoryTemplate, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(OnePageAdvisoryTemplates) > 0 {
		m["OnePageAdvisoryTemplate"] = OnePageAdvisoryTemplates
	} else {
		res := make([]models.OnePageAdvisoryTemplate, 0)
		m["OnePageAdvisoryTemplate"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
