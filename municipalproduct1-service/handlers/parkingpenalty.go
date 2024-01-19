package handlers

import (
	"encoding/json"
	"log"

	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

// SaveParkingPenalty : ""
func (h *Handler) SaveParkingPenalty(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	parkingPenalty := new(models.ParkingPenalty)
	err := json.NewDecoder(r.Body).Decode(&parkingPenalty)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveParkingPenalty(ctx, parkingPenalty)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["parkingPenalty"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleParkingPenalty : ""
func (h *Handler) GetSingleParkingPenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	parkingPenalty := new(models.RefParkingPenalty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	parkingPenalty, err := h.Service.GetSingleParkingPenalty(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["parkingPenalty"] = parkingPenalty
	response.With200V2(w, "Success", m, platform)
}

// UpdateParkingPenalty : ""
func (h *Handler) UpdateParkingPenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	parkingPenalty := new(models.ParkingPenalty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&parkingPenalty)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if parkingPenalty.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateParkingPenalty(ctx, parkingPenalty)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["parkingPenalty"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableParkingPenalty : ""
func (h *Handler) EnableParkingPenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableParkingPenalty(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["parkingPenalty"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableParkingPenalty : ""
func (h *Handler) DisableParkingPenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableParkingPenalty(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["parkingPenalty"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteParkingPenalty : ""
func (h *Handler) DeleteParkingPenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteParkingPenalty(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["parkingPenalty"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterParkingPenalty : ""
func (h *Handler) FilterParkingPenalty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ParkingPenaltyFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ParkingPenaltys []models.RefParkingPenalty
	log.Println(pagination)
	ParkingPenaltys, err = h.Service.FilterParkingPenalty(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(ParkingPenaltys) > 0 {
		m["parkingPenalty"] = ParkingPenaltys
	} else {
		res := make([]models.RefParkingPenalty, 0)
		m["parkingPenalty"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
