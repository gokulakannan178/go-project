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

// SaveShoprentPaymentModeChange : ""
func (h *Handler) SaveShoprentPaymentModeChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	request := new(models.ShoprentPaymentModeChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.SaveShoprentPaymentModeChange(ctx, request)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayment"] = request
	response.With200V2(w, "Success", m, platform)
}

// GetSingleReassessmentRequest :""
func (h *Handler) GetSingleShoprentPaymentModeChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "uniqueId is missing", platform)
	}

	rppmcr := new(models.RefShoprentPaymentModeChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	rppmcr, err := h.Service.GetSingleShoprentPaymentModeChange(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayment"] = rppmcr
	response.With200V2(w, "Success", m, platform)
}

// AcceptShoprentPaymentModeChange : ""
func (h *Handler) AcceptShoprentPaymentModeChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.AcceptShoprentPaymentModeChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.AcceptShoprentPaymentModeChange(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayment"] = req
	response.With200V2(w, "Success", m, platform)
}

// RejectShoprentPaymentModeChange : ""
func (h *Handler) RejectShoprentPaymentModeChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.RejectShoprentPaymentModeChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.RejectShoprentPaymentModeChange(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterShoprentPaymentModeChange : ""
func (h *Handler) FilterShoprentPaymentModeChange(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ShoprentPaymentModeChangeFilter
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

	var requests []models.RefShoprentPaymentModeChange
	log.Println(pagination)
	requests, err = h.Service.FilterShoprentPaymentModeChange(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(requests) > 0 {
		m["shoprentpayment"] = requests
	} else {
		res := make([]models.RefShoprentPaymentModeChange, 0)
		m["shoprentpayment"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// UpdateShoprentPaymentModeChangePropertyID : ""
func (h *Handler) UpdateShoprentPaymentModeChangePropertyID(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.UpdatePropertyUniqueID)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if len(property.UniqueIDs) == 0 {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateShoprentPaymentModeChangePropertyID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["shoprentpayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}
