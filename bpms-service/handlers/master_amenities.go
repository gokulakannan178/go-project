package handlers

import (
	"bpms-service/app"
	"bpms-service/models"
	"bpms-service/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

//SaveAmenities : ""
func (h *Handler) SaveAmenities(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	amenities := new(models.Amenities)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&amenities)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveAmenities(ctx, amenities)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["amenities"] = amenities
	response.With200V2(w, "Success", m, platform)
}

//UpdateAmenities :""
func (h *Handler) UpdateAmenities(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	amenities := new(models.Amenities)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&amenities)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if amenities.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAmenities(ctx, amenities)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["amenities"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableAmenities : ""
func (h *Handler) EnableAmenities(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableAmenities(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["amenities"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableAmenities : ""
func (h *Handler) DisableAmenities(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableAmenities(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["amenities"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteAmenities : ""
func (h *Handler) DeleteAmenities(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteAmenities(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["amenities"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleAmenities :""
func (h *Handler) GetSingleAmenities(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	amenities := new(models.RefAmenities)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	amenities, err := h.Service.GetSingleAmenities(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["amenities"] = amenities
	response.With200V2(w, "Success", m, platform)
}

//FilterAmenities : ""
func (h *Handler) FilterAmenities(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var amenities *models.AmenitiesFilter
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
	err := json.NewDecoder(r.Body).Decode(&amenities)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var amenitiess []models.RefAmenities
	log.Println(pagination)
	amenitiess, err = h.Service.FilterAmenities(ctx, amenities, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(amenitiess) > 0 {
		m["amenities"] = amenitiess
	} else {
		res := make([]models.Amenities, 0)
		m["amenities"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
