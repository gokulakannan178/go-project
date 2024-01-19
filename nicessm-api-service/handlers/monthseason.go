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

//SaveMonthSeason : ""
func (h *Handler) SaveMonthSeason(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	MonthSeason := new(models.MonthSeason)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&MonthSeason)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveMonthSeason(ctx, MonthSeason)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["MonthSeason"] = MonthSeason
	response.With200V2(w, "Success", m, platform)
}

//UpdateMonthSeason :""
func (h *Handler) UpdateMonthSeason(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	MonthSeason := new(models.MonthSeason)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&MonthSeason)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if MonthSeason.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateMonthSeason(ctx, MonthSeason)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["MonthSeason"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableMonthSeason : ""
func (h *Handler) EnableMonthSeason(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableMonthSeason(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["MonthSeason"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableMonthSeason : ""
func (h *Handler) DisableMonthSeason(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableMonthSeason(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["MonthSeason"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteMonthSeason(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.MonthSeason)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteMonthSeason(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["MonthSeason"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleMonthSeason :""
func (h *Handler) GetSingleMonthSeason(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	MonthSeason := new(models.RefMonthSeason)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	MonthSeason, err := h.Service.GetSingleMonthSeason(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["MonthSeason"] = MonthSeason
	response.With200V2(w, "Success", m, platform)
}

//FilterMonthSeason : ""
func (h *Handler) FilterMonthSeason(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var MonthSeason *models.MonthSeasonFilter
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
	err := json.NewDecoder(r.Body).Decode(&MonthSeason)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var MonthSeasons []models.RefMonthSeason
	log.Println(pagination)
	MonthSeasons, err = h.Service.FilterMonthSeason(ctx, MonthSeason, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(MonthSeasons) > 0 {
		m["MonthSeason"] = MonthSeasons
	} else {
		res := make([]models.MonthSeason, 0)
		m["MonthSeason"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
