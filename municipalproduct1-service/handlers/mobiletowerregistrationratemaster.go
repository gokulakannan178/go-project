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

// SaveMobileTowerRegistrationRateMaster : ""
func (h *Handler) SaveMobileTowerRegistrationRateMaster(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	mobile := new(models.MobileTowerRegistrationRateMaster)
	err := json.NewDecoder(r.Body).Decode(&mobile)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveMobileTowerRegistrationRateMaster(ctx, mobile)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationratemaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleMobileTowerRegistrationRateMaster : ""
func (h *Handler) GetSingleMobileTowerRegistrationRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	mobile := new(models.MobileTowerRegistrationRateMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	mobile, err := h.Service.GetSingleMobileTowerRegistrationRateMaster(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationratemaster"] = mobile
	response.With200V2(w, "Success", m, platform)
}

// UpdateMobileTowerRegistrationRateMaster : ""
func (h *Handler) UpdateMobileTowerRegistrationRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	mobile := new(models.MobileTowerRegistrationRateMaster)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&mobile)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if mobile.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateMobileTowerRegistrationRateMaster(ctx, mobile)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationratemaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableMobileTowerRegistrationRateMaster : ""
func (h *Handler) EnableMobileTowerRegistrationRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableMobileTowerRegistrationRateMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationratemaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableMobileTowerRegistrationRateMaster : ""
func (h *Handler) DisableMobileTowerRegistrationRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableMobileTowerRegistrationRateMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationratemaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteMobileTowerRegistrationRateMaster : ""
func (h *Handler) DeleteMobileTowerRegistrationRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteMobileTowerRegistrationRateMaster(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationratemaster"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterMobileTowerRegistrationRateMaster : ""
func (h *Handler) FilterMobileTowerRegistrationRateMaster(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.MobileTowerRegistrationRateMasterFilter
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

	var users []models.RefMobileTowerRegistrationRateMaster
	log.Println(pagination)
	users, err = h.Service.FilterMobileTowerRegistrationRateMaster(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(users) > 0 {
		m["mobiletowerregistrationratemaster"] = users
	} else {
		res := make([]models.MobileTowerRegistrationRateMaster, 0)
		m["mobiletowerregistrationratemaster"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
