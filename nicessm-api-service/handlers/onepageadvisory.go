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

//SaveOnePageAdvisory : ""
func (h *Handler) SaveOnePageAdvisory(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	OnePageAdvisory := new(models.OnePageAdvisory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&OnePageAdvisory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveOnePageAdvisory(ctx, OnePageAdvisory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisory"] = OnePageAdvisory
	response.With200V2(w, "Success", m, platform)
}

//UpdateOnePageAdvisory :""
func (h *Handler) UpdateOnePageAdvisory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	OnePageAdvisory := new(models.OnePageAdvisory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&OnePageAdvisory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if OnePageAdvisory.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateOnePageAdvisory(ctx, OnePageAdvisory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableOnePageAdvisory : ""
func (h *Handler) EnableOnePageAdvisory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableOnePageAdvisory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableOnePageAdvisory : ""
func (h *Handler) DisableOnePageAdvisory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableOnePageAdvisory(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteOnePageAdvisory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.OnePageAdvisory)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteOnePageAdvisory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleOnePageAdvisory :""
func (h *Handler) GetSingleOnePageAdvisory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	OnePageAdvisory := new(models.RefOnePageAdvisory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	OnePageAdvisory, err := h.Service.GetSingleOnePageAdvisory(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnePageAdvisory"] = OnePageAdvisory
	response.With200V2(w, "Success", m, platform)
}

//FilterOnePageAdvisory : ""
func (h *Handler) FilterOnePageAdvisory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var OnePageAdvisory *models.OnePageAdvisoryFilter
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
	err := json.NewDecoder(r.Body).Decode(&OnePageAdvisory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var OnePageAdvisorys []models.RefOnePageAdvisory
	log.Println(pagination)
	OnePageAdvisorys, err = h.Service.FilterOnePageAdvisory(ctx, OnePageAdvisory, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(OnePageAdvisorys) > 0 {
		m["OnePageAdvisory"] = OnePageAdvisorys
	} else {
		res := make([]models.OnePageAdvisory, 0)
		m["OnePageAdvisory"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
