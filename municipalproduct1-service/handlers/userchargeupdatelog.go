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

// SaveUserChargeUpdateLog : ""
func (h *Handler) SaveUserChargeUpdateLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UserChargeUpdateLog := new(models.UserChargeUpdateLog)
	err := json.NewDecoder(r.Body).Decode(&UserChargeUpdateLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveUserChargeUpdateLog(ctx, UserChargeUpdateLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserChargeUpdateLog"] = UserChargeUpdateLog
	response.With200V2(w, "Success", m, platform)
}

// GetSingleUserChargeUpdateLog : ""
func (h *Handler) GetSingleUserChargeUpdateLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	UserChargeUpdateLog := new(models.RefUserChargeUpdateLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	UserChargeUpdateLog, err := h.Service.GetSingleUserChargeUpdateLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserChargeUpdateLog"] = UserChargeUpdateLog
	response.With200V2(w, "Success", m, platform)
}

// UpdateUserChargeUpdateLog : ""
func (h *Handler) UpdateUserChargeUpdateLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UserChargeUpdateLog := new(models.UserChargeUpdateLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&UserChargeUpdateLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if UserChargeUpdateLog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateUserChargeUpdateLog(ctx, UserChargeUpdateLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserChargeUpdateLog"] = UserChargeUpdateLog
	response.With200V2(w, "Success", m, platform)
}

//EnableUserChargeUpdateLog : ""
func (h *Handler) EnableUserChargeUpdateLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableUserChargeUpdateLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserChargeUpdateLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableUserChargeUpdateLog : ""
func (h *Handler) DisableUserChargeUpdateLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableUserChargeUpdateLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserChargeUpdateLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteUserChargeUpdateLog : ""
func (h *Handler) DeleteUserChargeUpdateLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteUserChargeUpdateLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserChargeUpdateLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterUserChargeUpdateLog : ""
func (h *Handler) FilterUserChargeUpdateLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.UserChargeUpdateLogFilter
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

	var UserChargeUpdateLogs []models.RefUserChargeUpdateLog
	log.Println(pagination)
	UserChargeUpdateLogs, err = h.Service.FilterUserChargeUpdateLog(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(UserChargeUpdateLogs) > 0 {
		m["UserChargeUpdateLog"] = UserChargeUpdateLogs
	} else {
		res := make([]models.RefUserChargeUpdateLog, 0)
		m["UserChargeUpdateLog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
