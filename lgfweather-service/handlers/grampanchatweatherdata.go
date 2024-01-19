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

//SaveGramPanchayatWeatherData : ""
func (h *Handler) SaveGramPanchayatWeatherData(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	GramPanchayatweatherdata := new(models.GramPanchayatWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&GramPanchayatweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveGramPanchayatWeatherData(ctx, GramPanchayatweatherdata)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["GramPanchayatweatherdata"] = GramPanchayatweatherdata
	response.With200V2(w, "Success", m, platform)
}

//UpdateGramPanchayatWeatherData :""
func (h *Handler) UpdateGramPanchayatWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	GramPanchayatweatherdata := new(models.GramPanchayatWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&GramPanchayatweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if GramPanchayatweatherdata.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateGramPanchayatWeatherData(ctx, GramPanchayatweatherdata)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["GramPanchayatweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableGramPanchayatWeatherData : ""
func (h *Handler) EnableGramPanchayatWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableGramPanchayatWeatherData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["GramPanchayatweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableGramPanchayatWeatherData : ""
func (h *Handler) DisableGramPanchayatWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableGramPanchayatWeatherData(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["GramPanchayatweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteGramPanchayatWeatherData : ""
func (h *Handler) DeleteGramPanchayatWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.GramPanchayatWeatherData)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteGramPanchayatWeatherData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vaccine"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleGramPanchayatWeatherData :""
func (h *Handler) GetSingleGramPanchayatWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	GramPanchayatweatherdata := new(models.RefGramPanchayatWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	GramPanchayatweatherdata, err := h.Service.GetSingleGramPanchayatWeatherData(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["GramPanchayatweatherdata"] = GramPanchayatweatherdata
	response.With200V2(w, "Success", m, platform)
}

//FilterGramPanchayatWeatherData : ""
func (h *Handler) FilterGramPanchayatWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var GramPanchayatweatherdata *models.GramPanchayatWeatherDataFilter
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
	err := json.NewDecoder(r.Body).Decode(&GramPanchayatweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var GramPanchayatweatherdatas []models.RefGramPanchayatWeatherData
	log.Println(pagination)
	GramPanchayatweatherdatas, err = h.Service.FilterGramPanchayatWeatherData(ctx, GramPanchayatweatherdata, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(GramPanchayatweatherdatas) > 0 {
		m["GramPanchayatweatherdata"] = GramPanchayatweatherdatas
	} else {
		res := make([]models.GramPanchayatWeatherData, 0)
		m["GramPanchayatweatherdata"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleGramPanchayatWeatherDataWithOpen(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if lat == "" {
		response.With400V2(w, "lat is missing", platform)
	}
	if lon == "" {
		response.With400V2(w, "lon is missing", platform)
	}

	GramPanchayatweatherdata := new(models.WeatherDataMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	GramPanchayatweatherdata, err := h.Service.GetWeatherData(ctx, lat, lon)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	fmt.Println("GramPanchayatweatherdata")
	m := make(map[string]interface{})
	m["GramPanchayatweatherdata"] = GramPanchayatweatherdata
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveGramPanchayatWeatherDataWithOpenWebsite(w http.ResponseWriter, r *http.Request) {
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
	err := h.Service.SaveGramPanchayatWeatherDataWithOpenWebsite(ctx, lat, lon)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["GramPanchayatweatherdata"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleGramPanchayatWeatherDataWithCurrentDate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	GramPanchayatweatherdata := new(models.RefGramPanchayatWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	GramPanchayatweatherdata, err := h.Service.GetSingleGramPanchayatWeatherDataWithCurrentDate(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["GramPanchayatweatherdata"] = GramPanchayatweatherdata
	response.With200V2(w, "Success", m, platform)
}
