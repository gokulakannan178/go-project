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

//SaveDistrictWeatherAlert : ""
func (h *Handler) SaveDistrictWeatherAlert(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	DistrictWeatherAlert := new(models.DistrictWeatherAlert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDistrictWeatherAlert(ctx, DistrictWeatherAlert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlert"] = DistrictWeatherAlert
	response.With200V2(w, "Success", m, platform)
}

//UpdateDistrictWeatherAlert :""
func (h *Handler) UpdateDistrictWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	DistrictWeatherAlert := new(models.DistrictWeatherAlert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if DistrictWeatherAlert.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateDistrictWeatherAlert(ctx, DistrictWeatherAlert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlert"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateDistrictWeatherAlertMasterwithWeatheralert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	DistrictWeatherAlert := new(models.UpdateDistrictWeatherAlert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if DistrictWeatherAlert.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateDistrictWeatherAlertMasterwithWeatheralert(ctx, DistrictWeatherAlert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlert"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDistrictWeatherAlert : ""
func (h *Handler) EnableDistrictWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDistrictWeatherAlert(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlert"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDistrictWeatherAlert : ""
func (h *Handler) DisableDistrictWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDistrictWeatherAlert(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlert"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteDistrictWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.DistrictWeatherAlert)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDistrictWeatherAlert(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlert"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDistrictWeatherAlert :""
func (h *Handler) GetSingleDistrictWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	DistrictWeatherAlert := new(models.RefDistrictWeatherAlert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	DistrictWeatherAlert, err := h.Service.GetSingleDistrictWeatherAlert(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlert"] = DistrictWeatherAlert
	response.With200V2(w, "Success", m, platform)
}

//FilterDistrictWeatherAlert : ""
func (h *Handler) FilterDistrictWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var DistrictWeatherAlert *models.DistrictWeatherAlertFilter
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
	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var DistrictWeatherAlerts []models.RefDistrictWeatherAlert
	log.Println(pagination)
	DistrictWeatherAlerts, err = h.Service.FilterDistrictWeatherAlert(ctx, DistrictWeatherAlert, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(DistrictWeatherAlerts) > 0 {
		m["DistrictWeatherAlert"] = DistrictWeatherAlerts
	} else {
		res := make([]models.DistrictWeatherAlert, 0)
		m["DistrictWeatherAlert"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
