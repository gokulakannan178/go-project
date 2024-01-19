package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

// SaveUserLocationTracker : ""
func (h *Handler) SaveUserLocationTracker(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	tracker := new(models.UserLocationTracker)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&tracker)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveUserLocationTracker(ctx, tracker)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userLocationTracker"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleUserLocationTracker : ""
func (h *Handler) GetSingleUserLocationTracker(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	tracker := new(models.RefUserLocationTracker)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	tracker, err := h.Service.GetSingleUserLocationTracker(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userLocationTracker"] = tracker
	response.With200V2(w, "Success", m, platform)
}

// UpdateUserLocationTracker : ""
func (h *Handler) UpdateUserLocationTracker(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	tracker := new(models.UserLocationTracker)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&tracker)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if tracker.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateUserLocationTracker(ctx, tracker)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userLocationTracker"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableUserLocationTracker : ""
func (h *Handler) EnableUserLocationTracker(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableUserLocationTracker(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userLocationTracker"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableUserLocationTracker : ""
func (h *Handler) DisableUserLocationTracker(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableUserLocationTracker(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userLocationTracker"] = "success"
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
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteUserLocationTracker(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userLocationTracker"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterUserLocationTracker : ""
func (h *Handler) FilterUserLocationTracker(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.UserLocationTrackerFilter
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

	var UserLocationTrackers []models.RefUserLocationTracker
	log.Println(pagination)
	UserLocationTrackers, err = h.Service.FilterUserLocationTracker(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(UserLocationTrackers) > 0 {
		m["userLocationTracker"] = UserLocationTrackers
	} else {
		res := make([]models.UserLocationTracker, 0)
		m["userLocationTracker"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// GetSingleUserLocationTracker : ""
func (h *Handler) UserLocationTrackerCoordinates(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	//UniqueID := r.URL.Query().Get("id")
	tracker := new(models.UserLocationTrackerCoordinates)

	err := json.NewDecoder(r.Body).Decode(&tracker)

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	trackers, err := h.Service.UserLocationTrackerCoordinates(ctx, tracker)
	fmt.Println("============", trackers)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userLocationTracker"] = trackers
	response.With200V2(w, "Success", m, platform)
}
