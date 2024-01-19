package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
	"strconv"
)

//SaveLandCropCalendar : ""
func (h *Handler) SaveLandCropCalendar(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	landCropCalendar := new(models.LandCropCalendar)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&landCropCalendar)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveLandCropCalendar(ctx, landCropCalendar)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCropCalendar"] = landCropCalendar
	response.With200V2(w, "Success", m, platform)
}

//UpdateLandCropCalendar :""
func (h *Handler) UpdateLandCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	landCropCalendar := new(models.LandCropCalendar)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&landCropCalendar)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if landCropCalendar.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateLandCropCalendar(ctx, landCropCalendar)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCropCalendar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableLandCropCalendar : ""
func (h *Handler) EnableLandCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableLandCropCalendar(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCropCalendar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableLandCropCalendar : ""
func (h *Handler) DisableLandCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableLandCropCalendar(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCropCalendar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteLandCropCalendar : ""
func (h *Handler) DeleteLandCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteLandCropCalendar(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCropCalendar"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleLandCropCalendar :""
func (h *Handler) GetSingleLandCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	landCropCalendar := new(models.RefLandCropCalendar)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	landCropCalendar, err := h.Service.GetSingleLandCropCalendar(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCropCalendar"] = landCropCalendar
	response.With200V2(w, "Success", m, platform)
}

//FilterLandCropCalendar : ""
func (h *Handler) FilterLandCropCalendar(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var landCropCalendar *models.LandCropCalendarFilter
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
	err := json.NewDecoder(r.Body).Decode(&landCropCalendar)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var landCropCalendars []models.RefLandCropCalendar
	log.Println(pagination)
	landCropCalendars, err = h.Service.FilterLandCropCalendar(ctx, landCropCalendar, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(landCropCalendars) > 0 {
		m["landCropCalendar"] = landCropCalendars
	} else {
		res := make([]models.LandCropCalendar, 0)
		m["landCropCalendar"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
