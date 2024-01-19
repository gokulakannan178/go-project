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

//SaveUserLoginLog : ""
func (h *Handler) SaveUserLoginLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UserLoginLog := new(models.UserLoginLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&UserLoginLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveUserLoginLog(ctx, UserLoginLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLoginLog"] = UserLoginLog
	response.With200V2(w, "Success", m, platform)
}

//UpdateUserLoginLog :""
func (h *Handler) UpdateUserLoginLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UserLoginLog := new(models.UserLoginLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&UserLoginLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if UserLoginLog.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateUserLoginLog(ctx, UserLoginLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLoginLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableUserLoginLog : ""
func (h *Handler) EnableUserLoginLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableUserLoginLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLoginLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableUserLoginLog : ""
func (h *Handler) DisableUserLoginLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableUserLoginLog(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLoginLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteUserLoginLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.UserLoginLog)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteUserLoginLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLoginLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleUserLoginLog :""
func (h *Handler) GetSingleUserLoginLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	UserLoginLog := new(models.RefUserLoginLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	UserLoginLog, err := h.Service.GetSingleUserLoginLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLoginLog"] = UserLoginLog
	response.With200V2(w, "Success", m, platform)
}

//FilterUserLoginLog : ""
func (h *Handler) FilterUserLoginLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var UserLoginLog *models.UserLoginLogFilter
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
	err := json.NewDecoder(r.Body).Decode(&UserLoginLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var UserLoginLogs []models.RefUserLoginLog
	log.Println(pagination)
	UserLoginLogs, err = h.Service.FilterUserLoginLog(ctx, UserLoginLog, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(UserLoginLogs) > 0 {
		m["UserLoginLog"] = UserLoginLogs
	} else {
		res := make([]models.UserLoginLog, 0)
		m["UserLoginLog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UserLogin(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UserLoginLog := new(models.UserLoginLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&UserLoginLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.UserLogin(ctx, UserLoginLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLoginLog"] = UserLoginLog
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateUserLogout(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UserLoginLog := new(models.UserLoginLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&UserLoginLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if UserLoginLog.UserId.IsZero() {
		response.With400V2(w, "UserId is missing", platform)
		return
	}
	err = h.Service.UpdateUserLogout(ctx, UserLoginLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLoginLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}
