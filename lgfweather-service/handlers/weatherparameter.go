package handlers

import (
	"encoding/json"
	"lgfweather-service/app"
	"lgfweather-service/models"
	"lgfweather-service/response"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveWeatherParameter : ""
func (h *Handler) SaveWeatherParameter(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	WeatherParameter := new(models.WeatherParameter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&WeatherParameter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveWeatherParameter(ctx, WeatherParameter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherParameter"] = WeatherParameter
	response.With200V2(w, "Success", m, platform)
}

//UpdateWeatherParameter :""
func (h *Handler) UpdateWeatherParameter(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	WeatherParameter := new(models.WeatherParameter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&WeatherParameter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if WeatherParameter.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateWeatherParameter(ctx, WeatherParameter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherParameter"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableWeatherParameter : ""
func (h *Handler) EnableWeatherParameter(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableWeatherParameter(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherParameter"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableWeatherParameter : ""
func (h *Handler) DisableWeatherParameter(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableWeatherParameter(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherParameter"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteWeatherParameter(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.WeatherParameter)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteWeatherParameter(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherParameter"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleWeatherParameter :""
func (h *Handler) GetSingleWeatherParameter(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	WeatherParameter := new(models.RefWeatherParameter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	WeatherParameter, err := h.Service.GetSingleWeatherParameter(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WeatherParameter"] = WeatherParameter
	response.With200V2(w, "Success", m, platform)
}

//FilterWeatherParameter : ""
func (h *Handler) FilterWeatherParameter(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var WeatherParameter *models.WeatherParameterFilter
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
	err := json.NewDecoder(r.Body).Decode(&WeatherParameter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var WeatherParameters []models.RefWeatherParameter
	log.Println(pagination)
	WeatherParameters, err = h.Service.FilterWeatherParameter(ctx, WeatherParameter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(WeatherParameters) > 0 {
		m["WeatherParameter"] = WeatherParameters
	} else {
		res := make([]models.WeatherParameter, 0)
		m["WeatherParameter"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
