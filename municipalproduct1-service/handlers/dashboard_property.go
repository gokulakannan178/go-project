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

// SaveDashBoardProperty : ""
func (h *Handler) SaveDashBoardProperty(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	property := new(models.PropertyDashBoard)
	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveDashBoardProperty(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyDashBoard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleDashBoardProperty : ""
func (h *Handler) GetSingleDashBoardProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	property := new(models.RefPropertyDashBoard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	property, err := h.Service.GetSingleDashBoardProperty(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyDashBoard"] = property
	response.With200V2(w, "Success", m, platform)
}

// UpdateDashBoardProperty : ""
func (h *Handler) UpdateDashBoardProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.PropertyDashBoard)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if property.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateDashBoardProperty(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyDashBoard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDashBoardProperty : ""
func (h *Handler) EnableDashBoardProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDashBoardProperty(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyDashBoard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableDashBoardProperty : ""
func (h *Handler) DisableDashBoardProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDashBoardProperty(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyDashBoard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteDashBoardProperty : ""
func (h *Handler) DeleteDashBoardProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDashBoardProperty(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyDashBoard"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterDashBoardProperty : ""
func (h *Handler) FilterDashBoardProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PropertyDashBoardFilter
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

	var propertyDashBoards []models.RefPropertyDashBoard
	log.Println(pagination)
	propertyDashBoards, err = h.Service.FilterDashBoardProperty(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(propertyDashBoards) > 0 {
		m["propertyDashBoard"] = propertyDashBoards
	} else {
		res := make([]models.PropertyDashBoard, 0)
		m["propertyDashBoard"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// GetOverAllPropertyDashBoard : ""
func (h *Handler) GetOverAllPropertyDashBoard(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	property, err := h.Service.GetOverAllPropertyDashBoard(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyDashBoard"] = property
	response.With200V2(w, "Success", m, platform)
}
