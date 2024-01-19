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

//SaveBlockWeatherData : ""
func (h *Handler) SaveBlockWeatherData(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Blockweatherdata := new(models.BlockWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Blockweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveBlockWeatherData(ctx, Blockweatherdata)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Blockweatherdata"] = Blockweatherdata
	response.With200V2(w, "Success", m, platform)
}

//UpdateBlockWeatherData :""
func (h *Handler) UpdateBlockWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	Blockweatherdata := new(models.BlockWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&Blockweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if Blockweatherdata.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateBlockWeatherData(ctx, Blockweatherdata)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Blockweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableBlockWeatherData : ""
func (h *Handler) EnableBlockWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableBlockWeatherData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Blockweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableBlockWeatherData : ""
func (h *Handler) DisableBlockWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableBlockWeatherData(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Blockweatherdata"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteBlockWeatherData : ""
func (h *Handler) DeleteBlockWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.BlockWeatherData)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteBlockWeatherData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vaccine"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleBlockWeatherData :""
func (h *Handler) GetSingleBlockWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	Blockweatherdata := new(models.RefBlockWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Blockweatherdata, err := h.Service.GetSingleBlockWeatherData(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Blockweatherdata"] = Blockweatherdata
	response.With200V2(w, "Success", m, platform)
}

//FilterBlockWeatherData : ""
func (h *Handler) FilterBlockWeatherData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var Blockweatherdata *models.BlockWeatherDataFilter
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
	err := json.NewDecoder(r.Body).Decode(&Blockweatherdata)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var Blockweatherdatas []models.RefBlockWeatherData
	log.Println(pagination)
	Blockweatherdatas, err = h.Service.FilterBlockWeatherData(ctx, Blockweatherdata, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Blockweatherdatas) > 0 {
		m["Blockweatherdata"] = Blockweatherdatas
	} else {
		res := make([]models.BlockWeatherData, 0)
		m["Blockweatherdata"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleBlockWeatherDataWithOpen(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	if lat == "" {
		response.With400V2(w, "lat is missing", platform)
	}
	if lon == "" {
		response.With400V2(w, "lon is missing", platform)
	}

	Blockweatherdata := new(models.WeatherDataMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Blockweatherdata, err := h.Service.GetWeatherData(ctx, lat, lon)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	fmt.Println("Blockweatherdata")
	m := make(map[string]interface{})
	m["Blockweatherdata"] = Blockweatherdata
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveBlockWeatherDataWithOpenWebsite(w http.ResponseWriter, r *http.Request) {
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
	err := h.Service.SaveBlockWeatherDataWithOpenWebsite(ctx, lat, lon)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Blockweatherdata"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleBlockWeatherDataWithCurrentDate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	Blockweatherdata := new(models.RefBlockWeatherData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Blockweatherdata, err := h.Service.GetSingleBlockWeatherDataWithCurrentDate(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Blockweatherdata"] = Blockweatherdata
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetBlockWeatherDataByBlockId(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var blockweatherdata []models.RefBlockWeatherData
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	blockweatherdata, err := h.Service.GetBlockWeatherDataByBlockId(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["blockweatherdata"] = blockweatherdata
	response.With200V2(w, "Success", m, platform)
}
