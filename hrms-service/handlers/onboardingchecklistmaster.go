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

// SaveOnboardingCheckListMaster : ""
func (h *Handler) SaveOnboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	onboardingchecklistmaster := new(models.OnboardingCheckListMaster)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&onboardingchecklistmaster)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveOnboardingCheckListMaster(ctx, onboardingchecklistmaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingCheckListMaster"] = onboardingchecklistmaster
	response.With200V2(w, "Success", m, platform)
}

// GetSingleOnboardingCheckListMaster : ""
func (h *Handler) GetSingleOnboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefOnboardingCheckListMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleOnboardingCheckListMaster(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingCheckListMaster"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateOnboardingCheckListMaster : ""
func (h *Handler) UpdateOnboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	onboardingchecklistmaster := new(models.OnboardingCheckListMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&onboardingchecklistmaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if onboardingchecklistmaster.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateOnboardingCheckListMaster(ctx, onboardingchecklistmaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingCheckListMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableOnboardingCheckListMaster : ""
func (h *Handler) EnableOnboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableOnboardingCheckListMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingCheckListMaster"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableOnboardingCheckListMaster : ""
func (h *Handler) DisableOnboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableOnboardingCheckListMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingCheckListMaster"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteOnboardingCheckListMaster : ""
func (h *Handler) DeleteOnboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteOnboardingCheckListMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OnboardingCheckListMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterOnboardingCheckListMaster : ""
func (h *Handler) FilterOnboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterOnboardingCheckListMaster
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

	var fts []models.RefOnboardingCheckListMaster
	log.Println(pagination)
	fts, err = h.Service.FilterOnboardingCheckListMaster(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["OnboardingCheckListMaster"] = fts
	} else {
		res := make([]models.OnboardingCheckListMaster, 0)
		m["OnboardingCheckListMaster"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
