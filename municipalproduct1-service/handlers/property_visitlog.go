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

// SavePropertyVisitLog : ""
func (h *Handler) SavePropertyVisitLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	visitlog := new(models.PropertyVisitLog)
	err := json.NewDecoder(r.Body).Decode(&visitlog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyVisitLog(ctx, visitlog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["visitlog"] = visitlog
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePropertyVisitLog : ""
func (h *Handler) GetSinglePropertyVisitLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	visitlog := new(models.RefPropertyVisitLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	visitlog, err := h.Service.GetSinglePropertyVisitLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["visitlog"] = visitlog
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyVisitLog : ""
func (h *Handler) UpdatePropertyVisitLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	visitlog := new(models.PropertyVisitLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&visitlog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if visitlog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyVisitLog(ctx, visitlog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["visitlog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePropertyVisitLog : ""
func (h *Handler) EnablePropertyVisitLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnablePropertyVisitLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["visitlog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisablePropertyVisitLog : ""
func (h *Handler) DisablePropertyVisitLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisablePropertyVisitLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["visitlog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeletePropertyVisitLog : ""
func (h *Handler) DeletePropertyVisitLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeletePropertyVisitLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["visitlog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyVisitLog : ""
func (h *Handler) FilterPropertyVisitLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PropertyVisitLogFilter
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

	var visitlogs []models.RefPropertyVisitLog
	log.Println(pagination)
	visitlogs, err = h.Service.FilterPropertyVisitLog(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(visitlogs) > 0 {
		m["visitlogs"] = visitlogs
	} else {
		res := make([]models.PropertyVisitLog, 0)
		m["visitlog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyVisitLogPropertyID : ""
func (h *Handler) UpdatePropertyVisitLogPropertyID(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.UpdatePropertyUniqueID)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if len(property.UniqueIDs) == 0 {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyVisitLogPropertyID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["updatePropertyVisitLogPropertyID"] = "success"
	response.With200V2(w, "Success", m, platform)
}
