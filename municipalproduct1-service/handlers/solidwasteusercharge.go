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

// SaveSolidWasteUserCharge : ""
func (h *Handler) SaveSolidWasteUserCharge(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	solidwasteusercharge := new(models.SolidWasteUserCharge)
	err := json.NewDecoder(r.Body).Decode(&solidwasteusercharge)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveSolidWasteUserCharge(ctx, solidwasteusercharge)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteusercharge"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleSolidWasteUserCharge : ""
func (h *Handler) GetSingleSolidWasteUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	solidwasteusercharge := new(models.RefSolidWasteUserCharge)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	solidwasteusercharge, err := h.Service.GetSingleSolidWasteUserCharge(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteusercharge"] = solidwasteusercharge
	response.With200V2(w, "Success", m, platform)
}

// UpdateSolidWasteUserCharge : ""
func (h *Handler) UpdateSolidWasteUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	solidwasteusercharge := new(models.SolidWasteUserCharge)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&solidwasteusercharge)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if solidwasteusercharge.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateSolidWasteUserCharge(ctx, solidwasteusercharge)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteusercharge"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableSolidWasteUserCharge : ""
func (h *Handler) EnableSolidWasteUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableSolidWasteUserCharge(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteusercharge"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableSolidWasteUserCharge : ""
func (h *Handler) DisableSolidWasteUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableSolidWasteUserCharge(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteusercharge"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteSolidWasteUserCharge : ""
func (h *Handler) DeleteSolidWasteUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteSolidWasteUserCharge(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["solidwasteusercharge"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterSolidWasteUserCharge : ""
func (h *Handler) FilterSolidWasteUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.SolidWasteUserChargeFilter
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

	var solidwasteusercharges []models.RefSolidWasteUserCharge
	log.Println(pagination)
	solidwasteusercharges, err = h.Service.FilterSolidWasteUserCharge(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(solidwasteusercharges) > 0 {
		m["solidwasteusercharge"] = solidwasteusercharges
	} else {
		res := make([]models.SolidWasteUserCharge, 0)
		m["solidwasteusercharge"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
