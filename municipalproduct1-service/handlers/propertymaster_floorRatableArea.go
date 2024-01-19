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

//SaveFloorRatableArea : ""
func (h *Handler) SaveFloorRatableArea(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	floorRatableArea := new(models.FloorRatableArea)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&floorRatableArea)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveFloorRatableArea(ctx, floorRatableArea)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["floorRatableArea"] = floorRatableArea
	response.With200V2(w, "Success", m, platform)
}

//UpdateFloorRatableArea :""
func (h *Handler) UpdateFloorRatableArea(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	floorRatableArea := new(models.FloorRatableArea)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&floorRatableArea)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if floorRatableArea.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateFloorRatableArea(ctx, floorRatableArea)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["floorRatableArea"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableFloorRatableArea : ""
func (h *Handler) EnableFloorRatableArea(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableFloorRatableArea(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["floorRatableArea"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableFloorRatableArea : ""
func (h *Handler) DisableFloorRatableArea(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableFloorRatableArea(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["floorRatableArea"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteFloorRatableArea : ""
func (h *Handler) DeleteFloorRatableArea(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteFloorRatableArea(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["floorRatableArea"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleFloorRatableArea :""
func (h *Handler) GetSingleFloorRatableArea(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	floorRatableArea := new(models.RefFloorRatableArea)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	floorRatableArea, err := h.Service.GetSingleFloorRatableArea(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["floorRatableArea"] = floorRatableArea
	response.With200V2(w, "Success", m, platform)
}

//FilterFloorRatableArea : ""
func (h *Handler) FilterFloorRatableArea(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var floorRatableArea *models.FloorRatableAreaFilter
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
	err := json.NewDecoder(r.Body).Decode(&floorRatableArea)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var floorRatableAreas []models.RefFloorRatableArea
	log.Println(pagination)
	floorRatableAreas, err = h.Service.FilterFloorRatableArea(ctx, floorRatableArea, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(floorRatableAreas) > 0 {
		m["floorRatableArea"] = floorRatableAreas
	} else {
		res := make([]models.FloorRatableArea, 0)
		m["floorRatableArea"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
