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

// SaveMobileTowerRegistrationTax : ""
func (h *Handler) SaveMobileTowerRegistrationTax(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	mtrt := new(models.MobileTowerRegistrationTax)
	err := json.NewDecoder(r.Body).Decode(&mtrt)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveMobileTowerRegistrationTax(ctx, mtrt)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationtax"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleMobileTowerRegistrationTax : ""
func (h *Handler) GetSingleMobileTowerRegistrationTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	mtrt := new(models.RefMobileTowerRegistrationTax)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	mtrt, err := h.Service.GetSingleMobileTowerRegistrationTax(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationtax"] = mtrt
	response.With200V2(w, "Success", m, platform)
}

// GetSingleMobileTowerRegistrationTax : ""
func (h *Handler) GetSingleDefaultMobileTowerRegistrationTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	mtrt := new(models.RefMobileTowerRegistrationTax)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	mtrt, err := h.Service.GetSingleDefaultMobileTowerRegistrationTax(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationtax"] = mtrt
	response.With200V2(w, "Success", m, platform)
}

// UpdateMobileTowerRegistrationTax : ""
func (h *Handler) UpdateMobileTowerRegistrationTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	mtrt := new(models.MobileTowerRegistrationTax)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&mtrt)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if mtrt.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateMobileTowerRegistrationTax(ctx, mtrt)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationtax"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableMobileTowerRegistrationTax : ""
func (h *Handler) EnableMobileTowerRegistrationTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableMobileTowerRegistrationTax(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationtax"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableMobileTowerRegistrationTax : ""
func (h *Handler) DisableMobileTowerRegistrationTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableMobileTowerRegistrationTax(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationtax"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteMobileTowerRegistrationTax : ""
func (h *Handler) DeleteMobileTowerRegistrationTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteMobileTowerRegistrationTax(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobiletowerregistrationtax"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterMobileTowerRegistrationTax : ""
func (h *Handler) FilterMobileTowerRegistrationTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.MobileTowerRegistrationTaxFilter
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

	var MobileTowerRegistrationTaxs []models.RefMobileTowerRegistrationTax
	log.Println(pagination)
	MobileTowerRegistrationTaxs, err = h.Service.FilterMobileTowerRegistrationTax(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(MobileTowerRegistrationTaxs) > 0 {
		m["mobiletowerregistrationtax"] = MobileTowerRegistrationTaxs
	} else {
		res := make([]models.MobileTowerRegistrationTax, 0)
		m["mobiletowerregistrationtax"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
