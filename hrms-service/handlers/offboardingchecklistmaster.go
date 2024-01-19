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

// SaveOffboardingCheckListMaster : ""
func (h *Handler) SaveOffboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	offboardingchecklistmaster := new(models.OffboardingCheckListMaster)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&offboardingchecklistmaster)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveOffboardingCheckListMaster(ctx, offboardingchecklistmaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingCheckListMaster"] = offboardingchecklistmaster
	response.With200V2(w, "Success", m, platform)
}

// GetSingleOffboardingCheckListMaster : ""
func (h *Handler) GetSingleOffboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefOffboardingCheckListMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleOffboardingCheckListMaster(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingCheckListMaster"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateOffboardingCheckListMaster : ""
func (h *Handler) UpdateOffboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	offboardingchecklistmaster := new(models.OffboardingCheckListMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&offboardingchecklistmaster)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if offboardingchecklistmaster.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateOffboardingCheckListMaster(ctx, offboardingchecklistmaster)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingCheckListMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableOffboardingCheckListMaster : ""
func (h *Handler) EnableOffboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableOffboardingCheckListMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingCheckListMaster"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableOffboardingCheckListMaster : ""
func (h *Handler) DisableOffboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableOffboardingCheckListMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingCheckListMaster"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteOffboardingCheckListMaster : ""
func (h *Handler) DeleteOffboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteOffboardingCheckListMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OffboardingCheckListMaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterOffboardingCheckListMaster : ""
func (h *Handler) FilterOffboardingCheckListMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterOffboardingCheckListMaster
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

	var fts []models.RefOffboardingCheckListMaster
	log.Println(pagination)
	fts, err = h.Service.FilterOffboardingCheckListMaster(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["OffboardingCheckListMaster"] = fts
	} else {
		res := make([]models.OffboardingCheckListMaster, 0)
		m["OffboardingCheckListMaster"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
