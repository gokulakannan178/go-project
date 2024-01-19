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

// SaveOffboardingPolicy : ""
func (h *Handler) SaveOffboardingPolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	offboardingpolicy := new(models.OffboardingPolicy)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&offboardingpolicy)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveOffboardingPolicy(ctx, offboardingpolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingPolicy"] = offboardingpolicy
	response.With200V2(w, "Success", m, platform)
}

// GetSingleOffboardingPolicy : ""
func (h *Handler) GetSingleOffboardingPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefOffboardingPolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleOffboardingPolicy(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingPolicy"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateOffboardingPolicy : ""
func (h *Handler) UpdateOffboardingPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	OffboardingPolicy := new(models.OffboardingPolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&OffboardingPolicy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if OffboardingPolicy.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateOffboardingPolicy(ctx, OffboardingPolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableOffboardingPolicy : ""
func (h *Handler) EnableOffboardingPolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableOffboardingPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingPolicy"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableOffboardingPolicy : ""
func (h *Handler) DisableOffboardingPolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableOffboardingPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingPolicy"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteOffboardingPolicy : ""
func (h *Handler) DeleteOffboardingPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteOffboardingPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterOffboardingPolicy : ""
func (h *Handler) FilterOffboardingPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterOffboardingPolicy
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

	var fts []models.RefOffboardingPolicy
	log.Println(pagination)
	fts, err = h.Service.FilterOffboardingPolicy(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["OffboardingPolicy"] = fts
	} else {
		res := make([]models.OffboardingPolicy, 0)
		m["OffboardingPolicy"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
