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

//SaveVillageWeatherData : ""
func (h *Handler) SaveVillageWeatherData(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Villageweatherdata := new(models.VillageWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Villageweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveVillageWeatherData(ctx, Villageweatherdata)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Villageweatherdata"] = Villageweatherdata
	response.With200V2(w, "Success", m, platform)
}

//UpdateVillageWeatherData :""
func (h *Handler) UpdateVillageWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	Villageweatherdata := new(models.VillageWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&Villageweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if Villageweatherdata.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateVillageWeatherData(ctx, Villageweatherdata)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Villageweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableVillageWeatherData : ""
func (h *Handler) EnableVillageWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableVillageWeatherData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Villageweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableVillageWeatherData : ""
func (h *Handler) DisableVillageWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableVillageWeatherData(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Villageweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteVillageWeatherData : ""
func (h *Handler) DeleteVillageWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.VillageWeatherData)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteVillageWeatherData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vaccine"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleVillageWeatherData :""
func (h *Handler) GetSingleVillageWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	Villageweatherdata := new(models.RefVillageWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Villageweatherdata, err := h.Service.GetSingleVillageWeatherData(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Villageweatherdata"] = Villageweatherdata
	response.With200V2(w, "Success", m, platform)
}

//FilterVillageWeatherData : ""
func (h *Handler) FilterVillageWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var Villageweatherdata *models.VillageWeatherDataFilter
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
	err := json.NewDecoder(r.Body).Decode(&Villageweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var Villageweatherdatas []models.RefVillageWeatherData
	log.Println(pagination)
	Villageweatherdatas, err = h.Service.FilterVillageWeatherData(ctx, Villageweatherdata, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Villageweatherdatas) > 0 {
		m["Villageweatherdata"] = Villageweatherdatas
	} else {
		res := make([]models.VillageWeatherData, 0)
		m["Villageweatherdata"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleVillageWeatherDataWithOpen(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if lat == "" {
		response.With400V2(w, "lat is missing", platform)
	}
	if lon == "" {
		response.With400V2(w, "lon is missing", platform)
	}

	Villageweatherdata := new(models.WeatherDataMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Villageweatherdata, err := h.Service.GetWeatherData(ctx, lat, lon)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	fmt.Println("Villageweatherdata")
	m := make(map[string]interface{})
	m["Villageweatherdata"] = Villageweatherdata
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveVillageWeatherDataWithOpenWebsite(w http.ResponseWriter, r *http.Request) {
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
	err := h.Service.SaveVillageWeatherDataWithOpenWebsite(ctx, lat, lon)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Villageweatherdata"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleVillageWeatherDataWithCurrentDate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	Villageweatherdata := new(models.RefVillageWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Villageweatherdata, err := h.Service.GetSingleVillageWeatherDataWithCurrentDate(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Villageweatherdata"] = Villageweatherdata
	response.With200V2(w, "Success", m, platform)
}
