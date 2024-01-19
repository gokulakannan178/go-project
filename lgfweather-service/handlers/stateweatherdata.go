package handlers

import (
	"encoding/json"
	"fmt"
	"lgfweather-service/app"
	"lgfweather-service/models"
	"lgfweather-service/response"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveStateWeatherData : ""
func (h *Handler) SaveStateWeatherData(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	stateweatherdata := new(models.StateWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&stateweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveStateWeatherData(ctx, stateweatherdata)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["stateweatherdata"] = stateweatherdata
	response.With200V2(w, "Success", m, platform)
}

//UpdateStateWeatherData :""
func (h *Handler) UpdateStateWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	stateweatherdata := new(models.StateWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&stateweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if stateweatherdata.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateStateWeatherData(ctx, stateweatherdata)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["stateweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableStateWeatherData : ""
func (h *Handler) EnableStateWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableStateWeatherData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["stateweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableStateWeatherData : ""
func (h *Handler) DisableStateWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableStateWeatherData(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["stateweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteStateWeatherData : ""
func (h *Handler) DeleteStateWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.StateWeatherData)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteStateWeatherData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vaccine"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleStateWeatherData :""
func (h *Handler) GetSingleStateWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	stateweatherdata := new(models.RefStateWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	stateweatherdata, err := h.Service.GetSingleStateWeatherData(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["stateweatherdata"] = stateweatherdata
	response.With200V2(w, "Success", m, platform)
}

//FilterStateWeatherData : ""
func (h *Handler) FilterStateWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var stateweatherdata *models.StateWeatherDataFilter
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
	err := json.NewDecoder(r.Body).Decode(&stateweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var stateweatherdatas []models.RefStateWeatherData
	log.Println(pagination)
	stateweatherdatas, err = h.Service.FilterStateWeatherData(ctx, stateweatherdata, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(stateweatherdatas) > 0 {
		m["stateweatherdata"] = stateweatherdatas
	} else {
		res := make([]models.StateWeatherData, 0)
		m["stateweatherdata"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleStateWeatherDataWithOpen(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if lat == "" {
		response.With400V2(w, "lat is missing", platform)
	}
	if lon == "" {
		response.With400V2(w, "lon is missing", platform)
	}

	stateweatherdata := new(models.WeatherDataMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	stateweatherdata, err := h.Service.GetWeatherData(ctx, lat, lon)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	fmt.Println("stateweatherdata")
	m := make(map[string]interface{})
	m["stateweatherdata"] = stateweatherdata
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveStateWeatherDataWithOpenWebsite(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if lat == "" {
		response.With400V2(w, "lat is missing", platform)
	}
	if lon == "" {
		response.With400V2(w, "lon is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := h.Service.SaveStateWeatherDataWithOpenWebsite(ctx, lat, lon)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["stateweatherdata"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleStateWeatherDataWithCurrentDate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	stateweatherdata := new(models.RefStateWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	stateweatherdata, err := h.Service.GetSingleStateWeatherDataWithCurrentDate(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["stateweatherdata"] = stateweatherdata
	response.With200V2(w, "Success", m, platform)
}
