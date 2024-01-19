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

// SavePropertyFixedArvLog : ""
func (h *Handler) SavePropertyFixedArvLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	propertyfixedarvlog := new(models.PropertyFixedArvLog)
	err := json.NewDecoder(r.Body).Decode(&propertyfixedarvlog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyFixedArvLog(ctx, propertyfixedarvlog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyfixedarvlog"] = propertyfixedarvlog
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePropertyFixedArvLog : ""
func (h *Handler) GetSinglePropertyFixedArvLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	propertyfixedarvlog := new(models.RefPropertyFixedArvLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	propertyfixedarvlog, err := h.Service.GetSinglePropertyFixedArvLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyfixedarvlog"] = propertyfixedarvlog
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyFixedArvLog : ""
func (h *Handler) UpdatePropertyFixedArvLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	propertyfixedarvlog := new(models.PropertyFixedArvLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&propertyfixedarvlog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if propertyfixedarvlog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyFixedArvLog(ctx, propertyfixedarvlog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyfixedarvlog"] = propertyfixedarvlog
	response.With200V2(w, "Success", m, platform)
}

//EnablePropertyFixedArvLog : ""
func (h *Handler) EnablePropertyFixedArvLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnablePropertyFixedArvLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyfixedarvlog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisablePropertyFixedArvLog : ""
func (h *Handler) DisablePropertyFixedArvLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisablePropertyFixedArvLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyfixedarvlog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeletePropertyFixedArvLog : ""
func (h *Handler) DeletePropertyFixedArvLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeletePropertyFixedArvLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyfixedarvlog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyFixedArvLog : ""
func (h *Handler) FilterPropertyFixedArvLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PropertyFixedArvLogFilter
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

	var propertyfixedarvlogs []models.RefPropertyFixedArvLog
	log.Println(pagination)
	propertyfixedarvlogs, err = h.Service.FilterPropertyFixedArvLog(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(propertyfixedarvlogs) > 0 {
		m["propertyfixedarvlog"] = propertyfixedarvlogs
	} else {
		res := make([]models.PropertyFixedArvLog, 0)
		m["propertyfixedarvlog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
