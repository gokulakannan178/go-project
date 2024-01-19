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

//SaveWeatherAlertNotInRange : ""
func (h *Handler) SaveWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	WeatherAlertNotInRange := new(models.WeatherAlertNotInRange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&WeatherAlertNotInRange)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveWeatherAlertNotInRange(ctx, WeatherAlertNotInRange)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherAlertNotInRange"] = WeatherAlertNotInRange
	response.With200V2(w, "Success", m, platform)
}

//UpdateWeatherAlertNotInRange :""
func (h *Handler) UpdateWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	WeatherAlertNotInRange := new(models.WeatherAlertNotInRange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&WeatherAlertNotInRange)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if WeatherAlertNotInRange.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateWeatherAlertNotInRange(ctx, WeatherAlertNotInRange)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherAlertNotInRange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableWeatherAlertNotInRange : ""
func (h *Handler) EnableWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableWeatherAlertNotInRange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherAlertNotInRange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableWeatherAlertNotInRange : ""
func (h *Handler) DisableWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableWeatherAlertNotInRange(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherAlertNotInRange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.WeatherAlertNotInRange)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteWeatherAlertNotInRange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherAlertNotInRange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleWeatherAlertNotInRange :""
func (h *Handler) GetSingleWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	WeatherAlertNotInRange := new(models.RefWeatherAlertNotInRange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	WeatherAlertNotInRange, err := h.Service.GetSingleWeatherAlertNotInRange(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherAlertNotInRange"] = WeatherAlertNotInRange
	response.With200V2(w, "Success", m, platform)
}

//FilterWeatherAlertNotInRange : ""
func (h *Handler) FilterWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var WeatherAlertNotInRange *models.WeatherAlertNotInRangeFilter
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
	err := json.NewDecoder(r.Body).Decode(&WeatherAlertNotInRange)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var WeatherAlertNotInRanges []models.RefWeatherAlertNotInRange
	log.Println(pagination)
	WeatherAlertNotInRanges, err = h.Service.FilterWeatherAlertNotInRange(ctx, WeatherAlertNotInRange, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(WeatherAlertNotInRanges) > 0 {
		m["WeatherAlertNotInRange"] = WeatherAlertNotInRanges
	} else {
		res := make([]models.WeatherAlertNotInRange, 0)
		m["WeatherAlertNotInRange"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
