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

//SaveDistrictWeatherAlertNotInRange : ""
func (h *Handler) SaveDistrictWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	DistrictWeatherAlertNotInRange := new(models.DistrictWeatherAlertNotInRange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlertNotInRange)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDistrictWeatherAlertNotInRange(ctx, DistrictWeatherAlertNotInRange)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertNotInRange"] = DistrictWeatherAlertNotInRange
	response.With200V2(w, "Success", m, platform)
}

//UpdateDistrictWeatherAlertNotInRange :""
func (h *Handler) UpdateDistrictWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	DistrictWeatherAlertNotInRange := new(models.DistrictWeatherAlertNotInRange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlertNotInRange)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if DistrictWeatherAlertNotInRange.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateDistrictWeatherAlertNotInRange(ctx, DistrictWeatherAlertNotInRange)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertNotInRange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDistrictWeatherAlertNotInRange : ""
func (h *Handler) EnableDistrictWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDistrictWeatherAlertNotInRange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertNotInRange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDistrictWeatherAlertNotInRange : ""
func (h *Handler) DisableDistrictWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDistrictWeatherAlertNotInRange(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertNotInRange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteDistrictWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.DistrictWeatherAlertNotInRange)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDistrictWeatherAlertNotInRange(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertNotInRange"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDistrictWeatherAlertNotInRange :""
func (h *Handler) GetSingleDistrictWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	DistrictWeatherAlertNotInRange := new(models.RefDistrictWeatherAlertNotInRange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	DistrictWeatherAlertNotInRange, err := h.Service.GetSingleDistrictWeatherAlertNotInRange(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertNotInRange"] = DistrictWeatherAlertNotInRange
	response.With200V2(w, "Success", m, platform)
}

//FilterDistrictWeatherAlertNotInRange : ""
func (h *Handler) FilterDistrictWeatherAlertNotInRange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var DistrictWeatherAlertNotInRange *models.DistrictWeatherAlertNotInRangeFilter
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
	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlertNotInRange)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var DistrictWeatherAlertNotInRanges []models.RefDistrictWeatherAlertNotInRange
	log.Println(pagination)
	DistrictWeatherAlertNotInRanges, err = h.Service.FilterDistrictWeatherAlertNotInRange(ctx, DistrictWeatherAlertNotInRange, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(DistrictWeatherAlertNotInRanges) > 0 {
		m["DistrictWeatherAlertNotInRange"] = DistrictWeatherAlertNotInRanges
	} else {
		res := make([]models.DistrictWeatherAlertNotInRange, 0)
		m["DistrictWeatherAlertNotInRange"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
