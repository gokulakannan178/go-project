package handlers

import (
	"ecommerce-service/app"
	"ecommerce-service/models"
	"ecommerce-service/response"
	"encoding/json"
	"log"

	"net/http"
	"strconv"
)

// SavePaymentMode : ""
func (h *Handler) SavePaymentMode(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	PaymentMode := new(models.PaymentMode)
	err := json.NewDecoder(r.Body).Decode(&PaymentMode)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePaymentMode(ctx, PaymentMode)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PaymentMode"] = PaymentMode
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePaymentMode : ""
func (h *Handler) GetSinglePaymentMode(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	PaymentMode := new(models.RefPaymentMode)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	PaymentMode, err := h.Service.GetSinglePaymentMode(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PaymentMode"] = PaymentMode
	response.With200V2(w, "Success", m, platform)
}

// UpdatePaymentMode : ""
func (h *Handler) UpdatePaymentMode(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	PaymentMode := new(models.PaymentMode)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&PaymentMode)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if PaymentMode.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePaymentMode(ctx, PaymentMode)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PaymentMode"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePaymentMode : ""
func (h *Handler) EnablePaymentMode(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnablePaymentMode(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PaymentMode"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisablePaymentMode : ""
func (h *Handler) DisablePaymentMode(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisablePaymentMode(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PaymentMode"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeletePaymentMode : ""
func (h *Handler) DeletePaymentMode(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeletePaymentMode(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["PaymentMode"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterPaymentMode : ""
func (h *Handler) FilterPaymentMode(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PaymentModeFilter
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

	var PaymentModes []models.RefPaymentMode
	log.Println(pagination)
	PaymentModes, err = h.Service.FilterPaymentMode(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(PaymentModes) > 0 {
		m["PaymentMode"] = PaymentModes
	} else {
		res := make([]models.PaymentMode, 0)
		m["PaymentMode"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
