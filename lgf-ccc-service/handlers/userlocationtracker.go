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

// SaveUserLocationTracker : ""
func (h *Handler) SaveUserLocationTracker(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	userLocationTracker := new(models.UserLocationTracker)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&userLocationTracker)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveUserLocationTracker(ctx, userLocationTracker)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLocationTracker"] = userLocationTracker
	response.With200V2(w, "Success", m, platform)
}

// GetSingleUserLocationTracker : ""
func (h *Handler) GetSingleUserLocationTracker(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	userLocationTracker := new(models.RefUserLocationTracker)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	userLocationTracker, err := h.Service.GetSingleUserLocationTracker(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLocationTracker"] = userLocationTracker
	response.With200V2(w, "Success", m, platform)
}

//UpdateUserLocationTracker : ""
func (h *Handler) UpdateUserLocationTracker(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	userLocationTracker := new(models.UserLocationTracker)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&userLocationTracker)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if userLocationTracker.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateUserLocationTracker(ctx, userLocationTracker)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLocationTracker"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableUserLocationTracker : ""
func (h *Handler) EnableUserLocationTracker(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableUserLocationTracker(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLocationTracker"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableUserLocationTracker : ""
func (h *Handler) DisableUserLocationTracker(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableUserLocationTracker(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLocationTracker"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteUserLocationTracker : ""
func (h *Handler) DeleteUserLocationTracker(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteUserLocationTracker(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["UserLocationTracker"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterUserLocationTracker : ""
func (h *Handler) FilterUserLocationTracker(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filterUserLocationTracker *models.FilterUserLocationTracker
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
	err := json.NewDecoder(r.Body).Decode(&filterUserLocationTracker)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var filterUserLocationTrackers []models.RefUserLocationTracker
	log.Println(pagination)
	filterUserLocationTrackers, err = h.Service.FilterUserLocationTracker(ctx, filterUserLocationTracker, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(filterUserLocationTrackers) > 0 {
		m["UserLocationTracker"] = filterUserLocationTrackers
	} else {
		res := make([]models.UserLocationTracker, 0)
		m["UserLocationTracker"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
