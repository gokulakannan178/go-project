package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"log"
	"net/http"
	"strconv"
)

//SaveAreaAssignLog : ""
func (h *Handler) SaveAreaAssignLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	areaassignlog := new(models.AreaAssignLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&areaassignlog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveAreaAssignLog(ctx, areaassignlog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["areaassignlog"] = areaassignlog
	response.With200V2(w, "Success", m, platform)
}

//GetSingleAreaAssignLog :""
func (h *Handler) GetSingleAreaAssignLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var areaassignlog *models.RefAreaAssignLog
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	areaassignlog, err := h.Service.GetSingleAreaAssignLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["areaassignlog"] = areaassignlog
	response.With200V2(w, "Success", m, platform)
}

//UpdateAreaAssignLog :""
func (h *Handler) UpdateAreaAssignLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var areaassignlog *models.AreaAssignLog
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&areaassignlog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if areaassignlog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAreaAssignLog(ctx, areaassignlog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["areaassignlog"] = areaassignlog
	response.With200V2(w, "Success", m, platform)
}

//EnableAreaAssignLog : ""
func (h *Handler) EnableAreaAssignLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableAreaAssignLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["areaassignlog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableAreaAssignLog : ""
func (h *Handler) DisableAreaAssignLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableAreaAssignLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["areaassignlog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteAreaAssignLog : ""
func (h *Handler) DeleteAreaAssignLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteAreaAssignLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["areaassignlog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterAreaAssignLog : ""
func (h *Handler) FilterAreaAssignLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var AreaAssignLog *models.AreaAssignLogFilter
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
	err := json.NewDecoder(r.Body).Decode(&AreaAssignLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var AreaAssignLogs []models.AreaAssignLog
	log.Println(pagination)
	AreaAssignLogs, err = h.Service.AreaAssignLogFilter(ctx, AreaAssignLog, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(AreaAssignLogs) > 0 {
		m["areaassignlog"] = AreaAssignLogs
	} else {
		res := make([]models.User, 0)
		m["areaassignlog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
