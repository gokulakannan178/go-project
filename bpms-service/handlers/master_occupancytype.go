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

//SaveOccupancyType : ""
func (h *Handler) SaveOccupancyType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	occupancyType := new(models.OccupancyType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&occupancyType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveOccupancyType(ctx, occupancyType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["occupancyType"] = occupancyType
	response.With200V2(w, "Success", m, platform)
}

//UpdateOccupancyType :""
func (h *Handler) UpdateOccupancyType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	occupancyType := new(models.OccupancyType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&occupancyType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if occupancyType.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateOccupancyType(ctx, occupancyType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["occupancyType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableOccupancyType : ""
func (h *Handler) EnableOccupancyType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableOccupancyType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["occupancyType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableOccupancyType : ""
func (h *Handler) DisableOccupancyType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableOccupancyType(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["occupancyType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteOccupancyType : ""
func (h *Handler) DeleteOccupancyType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteOccupancyType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["occupancyType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleOccupancyType :""
func (h *Handler) GetSingleOccupancyType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	occupancyType := new(models.RefOccupancyType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	occupancyType, err := h.Service.GetSingleOccupancyType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["occupancyType"] = occupancyType
	response.With200V2(w, "Success", m, platform)
}

//FilterOccupancyType : ""
func (h *Handler) FilterOccupancyType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var occupancyType *models.OccupancyTypeFilter
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
	err := json.NewDecoder(r.Body).Decode(&occupancyType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var occupancyTypes []models.RefOccupancyType
	log.Println(pagination)
	occupancyTypes, err = h.Service.FilterOccupancyType(ctx, occupancyType, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(occupancyTypes) > 0 {
		m["occupancyType"] = occupancyTypes
	} else {
		res := make([]models.OccupancyType, 0)
		m["occupancyType"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
