package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"log"
	"net/http"
	"strconv"
)

//SaveVehicleLog : ""
func (h *Handler) SaveVehicleLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	vehiclelog := new(models.VehicleLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&vehiclelog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveVehicleLog(ctx, vehiclelog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelog"] = vehiclelog
	response.With200V2(w, "Success", m, platform)
}

//GetSingleVehicleLog :""
func (h *Handler) GetSingleVehicleLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var vehiclelog *models.RefVehicleLog
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	vehiclelog, err := h.Service.GetSingleVehicleLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelog"] = vehiclelog
	response.With200V2(w, "Success", m, platform)
}

//UpdateVehicleLog :""
func (h *Handler) UpdateVehicleLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var vehiclelog *models.VehicleLog
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&vehiclelog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if vehiclelog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateVehicleLog(ctx, vehiclelog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableVehicleLog : ""
func (h *Handler) EnableVehicleLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableVehicleLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableVehicleLog : ""
func (h *Handler) DisableVehicleLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableVehicleLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteVehicleLog : ""
func (h *Handler) DeleteVehicleLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteVehicleLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehiclelog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterVehicleLog : ""
func (h *Handler) FilterVehicleLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var vehiclelog *models.VehicleLogFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
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
	err := json.NewDecoder(r.Body).Decode(&vehiclelog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var vehiclelogs []models.VehicleLog
	log.Println(pagination)
	vehiclelogs, err = h.Service.VehicleLogFilter(ctx, vehiclelog, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(vehiclelogs) > 0 {
		m["vehiclelog"] = vehiclelogs
	} else {
		res := make([]models.User, 0)
		m["vehiclelog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
