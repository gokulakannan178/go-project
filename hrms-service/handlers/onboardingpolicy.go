package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

// SaveOnboardingPolicy : ""
func (h *Handler) SaveOnboardingPolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	onboardingpolicy := new(models.OnboardingPolicy)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&onboardingpolicy)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveOnboardingPolicy(ctx, onboardingpolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingPolicy"] = onboardingpolicy
	response.With200V2(w, "Success", m, platform)
}

// GetSingleOnboardingPolicy : ""
func (h *Handler) GetSingleOnboardingPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefOnboardingPolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleOnboardingPolicy(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingPolicy"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateOnboardingPolicy : ""
func (h *Handler) UpdateOnboardingPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	onboardingPolicy := new(models.OnboardingPolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&onboardingPolicy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if onboardingPolicy.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateOnboardingPolicy(ctx, onboardingPolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableOnboardingPolicy : ""
func (h *Handler) EnableOnboardingPolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableOnboardingPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingPolicy"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableOnboardingPolicy : ""
func (h *Handler) DisableOnboardingPolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableOnboardingPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingPolicy"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteOnboardingPolicy : ""
func (h *Handler) DeleteOnboardingPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteOnboardingPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterOnboardingPolicy : ""
func (h *Handler) FilterOnboardingPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterOnboardingPolicy
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
	err := json.NewDecoder(r.Body).Decode(&ft)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var fts []models.RefOnboardingPolicy
	log.Println(pagination)
	fts, err = h.Service.FilterOnboardingPolicy(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["OnboardingPolicy"] = fts
	} else {
		res := make([]models.OnboardingPolicy, 0)
		m["OnboardingPolicy"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
