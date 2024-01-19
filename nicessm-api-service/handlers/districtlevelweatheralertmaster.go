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

//SaveDistrictWeatherAlertMaster : ""
func (h *Handler) SaveDistrictWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	DistrictWeatherAlertMaster := new(models.DistrictWeatherAlertMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDistrictWeatherAlertMaster(ctx, DistrictWeatherAlertMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertMaster"] = DistrictWeatherAlertMaster
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) UpdateDistrictWeatherAlertMasterUpsertwithMin(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	DistrictWeatherAlertMaster := new(models.DistrictWeatherAlertMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.UpdateDistrictWeatherAlertMasterUpsertwithMin(ctx, DistrictWeatherAlertMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateDistrictWeatherAlertMasterUpsertwithMax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	DistrictWeatherAlertMaster := new(models.DistrictWeatherAlertMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.UpdateDistrictWeatherAlertMasterUpsertwithMax(ctx, DistrictWeatherAlertMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//UpdateDistrictWeatherAlertMaster :""
func (h *Handler) UpdateDistrictWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	DistrictWeatherAlertMaster := new(models.DistrictWeatherAlertMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if DistrictWeatherAlertMaster.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateDistrictWeatherAlertMaster(ctx, DistrictWeatherAlertMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDistrictWeatherAlertMaster : ""
func (h *Handler) EnableDistrictWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDistrictWeatherAlertMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDistrictWeatherAlertMaster : ""
func (h *Handler) DisableDistrictWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDistrictWeatherAlertMaster(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//v : ""
func (h *Handler) DeleteDistrictWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.DistrictWeatherAlertMaster)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDistrictWeatherAlertMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDistrictWeatherAlertMaster :""
func (h *Handler) GetSingleDistrictWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	DistrictWeatherAlertMaster := new(models.RefDistrictWeatherAlertMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	DistrictWeatherAlertMaster, err := h.Service.GetSingleDistrictWeatherAlertMaster(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["DistrictWeatherAlertMaster"] = DistrictWeatherAlertMaster
	response.With200V2(w, "Success", m, platform)
}

//FilterDistrictWeatherAlertMaster : ""
func (h *Handler) FilterDistrictWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var DistrictWeatherAlertMaster *models.DistrictWeatherAlertMasterFilter
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
	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var DistrictWeatherAlertMasters []models.RefDistrictWeatherAlertMaster
	log.Println(pagination)
	DistrictWeatherAlertMasters, err = h.Service.FilterDistrictWeatherAlertMaster(ctx, DistrictWeatherAlertMaster, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(DistrictWeatherAlertMasters) > 0 {
		m["DistrictWeatherAlertMaster"] = DistrictWeatherAlertMasters
	} else {
		res := make([]models.DistrictWeatherAlertMaster, 0)
		m["DistrictWeatherAlertMaster"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetDistrictWeatherAlertMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var DistrictWeatherAlertMaster *models.DistrictWeatherAlertMasterFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&DistrictWeatherAlertMaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var DistrictWeatherAlertMasters []models.GetWeatherAlertMaster

	DistrictWeatherAlertMasters, err = h.Service.GetDistrictWeatherAlertMaster(ctx, DistrictWeatherAlertMaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(DistrictWeatherAlertMasters) > 0 {
		m["DistrictWeatherAlertMaster"] = DistrictWeatherAlertMasters
	} else {
		res := make([]models.DistrictWeatherAlertMaster, 0)
		m["DistrictWeatherAlertMaster"] = res
	}

	response.With200V2(w, "Success", m, platform)
}
