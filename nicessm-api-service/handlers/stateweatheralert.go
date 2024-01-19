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

//SaveStateWeatherAlert : ""
func (h *Handler) SaveStateWeatherAlert(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	StateWeatherAlert := new(models.StateWeatherAlert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveStateWeatherAlert(ctx, StateWeatherAlert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlert"] = StateWeatherAlert
	response.With200V2(w, "Success", m, platform)
}

//UpdateStateWeatherAlert :""
func (h *Handler) UpdateStateWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	StateWeatherAlert := new(models.StateWeatherAlert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if StateWeatherAlert.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateStateWeatherAlert(ctx, StateWeatherAlert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlert"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateStateWeatherAlertMatser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	StateWeatherAlert := new(models.UpdateStateWeatherAlert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if StateWeatherAlert.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateWeatherAlertMaster(ctx, StateWeatherAlert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlert"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableStateWeatherAlert : ""
func (h *Handler) EnableStateWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableStateWeatherAlert(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlert"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableStateWeatherAlert : ""
func (h *Handler) DisableStateWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableStateWeatherAlert(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlert"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteStateWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.StateWeatherAlert)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteStateWeatherAlert(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlert"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleStateWeatherAlert :""
func (h *Handler) GetSingleStateWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	StateWeatherAlert := new(models.RefStateWeatherAlert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	StateWeatherAlert, err := h.Service.GetSingleStateWeatherAlert(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlert"] = StateWeatherAlert
	response.With200V2(w, "Success", m, platform)
}

//FilterStateWeatherAlert : ""
func (h *Handler) FilterStateWeatherAlert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var StateWeatherAlert *models.StateWeatherAlertFilter
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
	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var StateWeatherAlerts []models.RefStateWeatherAlert
	log.Println(pagination)
	StateWeatherAlerts, err = h.Service.FilterStateWeatherAlert(ctx, StateWeatherAlert, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(StateWeatherAlerts) > 0 {
		m["StateWeatherAlert"] = StateWeatherAlerts
	} else {
		res := make([]models.StateWeatherAlert, 0)
		m["StateWeatherAlert"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
