package handlers

import (
	"encoding/json"
	"log"
	"logikoof-echalan-service/app"
	"logikoof-echalan-service/models"
	"logikoof-echalan-service/response"
	"net/http"
	"strconv"
)

//SaveVehicleChallan : ""
func (h *Handler) SaveVehicleChallan(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	vehicleChallan := new(models.VehicleChallan)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&vehicleChallan)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveVehicleChallan(ctx, vehicleChallan)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleChallan"] = vehicleChallan
	response.With200V2(w, "Success", m, platform)
}

//UpdateVehicleChallan :""
func (h *Handler) UpdateVehicleChallan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	vehicleChallan := new(models.VehicleChallan)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&vehicleChallan)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if vehicleChallan.UniqueID == "" {
		response.With400V2(w, "RegNo is missing", platform)
	}
	err = h.Service.UpdateVehicleChallan(ctx, vehicleChallan)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleChallan"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableVehicleChallan : ""
func (h *Handler) EnableVehicleChallan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableVehicleChallan(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleChallan"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableVehicleChallan : ""
func (h *Handler) DisableVehicleChallan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableVehicleChallan(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleChallan"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteVehicleChallan : ""
func (h *Handler) DeleteVehicleChallan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteVehicleChallan(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleChallan"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleVehicleChallan :""
func (h *Handler) GetSingleVehicleChallan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	vehicleChallan := new(models.RefVehicleChallan)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	vehicleChallan, err := h.Service.GetSingleVehicleChallan(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleChallan"] = vehicleChallan
	response.With200V2(w, "Success", m, platform)
}

//FilterVehicleChallan : ""
func (h *Handler) FilterVehicleChallan(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var vehicleChallan *models.VehicleChallanFilter
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
	err := json.NewDecoder(r.Body).Decode(&vehicleChallan)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var vehicleChallans []models.RefVehicleChallan
	log.Println(pagination)
	vehicleChallans, err = h.Service.FilterVehicleChallan(ctx, vehicleChallan, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(vehicleChallans) > 0 {
		m["vehicleChallan"] = vehicleChallans
	} else {
		res := make([]models.VehicleChallan, 0)
		m["vehicleChallan"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
