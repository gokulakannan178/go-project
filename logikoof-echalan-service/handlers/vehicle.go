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

//SaveVehicle : ""
func (h *Handler) SaveVehicle(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	vehicle := new(models.Vehicle)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveVehicle(ctx, vehicle)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicle"] = vehicle
	response.With200V2(w, "Success", m, platform)
}

//UpdateVehicle :""
func (h *Handler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	vehicle := new(models.Vehicle)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&vehicle)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if vehicle.RegNo == "" {
		response.With400V2(w, "RegNo is missing", platform)
	}
	err = h.Service.UpdateVehicle(ctx, vehicle)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicle"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableVehicle : ""
func (h *Handler) EnableVehicle(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableVehicle(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicle"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableVehicle : ""
func (h *Handler) DisableVehicle(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableVehicle(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicle"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteVehicle : ""
func (h *Handler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteVehicle(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicle"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleVehicle :""
func (h *Handler) GetSingleVehicle(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	vehicle := new(models.RefVehicle)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	vehicle, err := h.Service.GetSingleVehicle(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicle"] = vehicle
	response.With200V2(w, "Success", m, platform)
}

//FilterVehicle : ""
func (h *Handler) FilterVehicle(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var vehicle *models.VehicleFilter
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
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var vehicles []models.RefVehicle
	log.Println(pagination)
	vehicles, err = h.Service.FilterVehicle(ctx, vehicle, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(vehicles) > 0 {
		m["vehicle"] = vehicles
	} else {
		res := make([]models.Vehicle, 0)
		m["vehicle"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
