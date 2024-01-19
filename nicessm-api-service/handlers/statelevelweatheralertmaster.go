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

//SaveStateWeatherAlertMaster : ""
func (h *Handler) SaveStateWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	StateWeatherAlertMaster := new(models.StateWeatherAlertMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveStateWeatherAlertMaster(ctx, StateWeatherAlertMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertMaster"] = StateWeatherAlertMaster
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) UpdateStateWeatherAlertMasterUpsertwithMin(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	StateWeatherAlertMaster := new(models.StateWeatherAlertMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.UpdateStateWeatherAlertMasterUpsertwithMin(ctx, StateWeatherAlertMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateStateWeatherAlertMasterUpsertwithMax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	StateWeatherAlertMaster := new(models.StateWeatherAlertMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.UpdateStateWeatherAlertMasterUpsertwithMax(ctx, StateWeatherAlertMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//UpdateStateWeatherAlertMaster :""
func (h *Handler) UpdateStateWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	StateWeatherAlertMaster := new(models.StateWeatherAlertMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if StateWeatherAlertMaster.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateStateWeatherAlertMaster(ctx, StateWeatherAlertMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableStateWeatherAlertMaster : ""
func (h *Handler) EnableStateWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableStateWeatherAlertMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableStateWeatherAlertMaster : ""
func (h *Handler) DisableStateWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableStateWeatherAlertMaster(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteStateWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.StateWeatherAlertMaster)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteStateWeatherAlertMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleStateWeatherAlertMaster :""
func (h *Handler) GetSingleStateWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	StateWeatherAlertMaster := new(models.RefStateWeatherAlertMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	StateWeatherAlertMaster, err := h.Service.GetSingleStateWeatherAlertMaster(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["StateWeatherAlertMaster"] = StateWeatherAlertMaster
	response.With200V2(w, "Success", m, platform)
}

//FilterStateWeatherAlertMaster : ""
func (h *Handler) FilterStateWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var StateWeatherAlertMaster *models.StateWeatherAlertMasterFilter
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
	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var StateWeatherAlertMasters []models.RefStateWeatherAlertMaster
	log.Println(pagination)
	StateWeatherAlertMasters, err = h.Service.FilterStateWeatherAlertMaster(ctx, StateWeatherAlertMaster, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(StateWeatherAlertMasters) > 0 {
		m["StateWeatherAlertMaster"] = StateWeatherAlertMasters
	} else {
		res := make([]models.StateWeatherAlertMaster, 0)
		m["StateWeatherAlertMaster"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetStateWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var StateWeatherAlertMaster *models.StateWeatherAlertMasterFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&StateWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var StateWeatherAlertMasters []models.GetWeatherAlertMaster

	StateWeatherAlertMasters, err = h.Service.GetStateWeatherAlertMaster(ctx, StateWeatherAlertMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(StateWeatherAlertMasters) > 0 {
		m["StateWeatherAlertMaster"] = StateWeatherAlertMasters
	} else {
		res := make([]models.StateWeatherAlertMaster, 0)
		m["StateWeatherAlertMaster"] = res
	}

	response.With200V2(w, "Success", m, platform)
}
