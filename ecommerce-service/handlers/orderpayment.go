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

// SaveOrderPayment : ""
func (h *Handler) SaveOrderPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	OrderPayment := new(models.OrderPayment)
	err := json.NewDecoder(r.Body).Decode(&OrderPayment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveOrderPayment(ctx, OrderPayment)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OrderPayment"] = OrderPayment
	response.With200V2(w, "Success", m, platform)
}

// GetSingleOrderPayment : ""
func (h *Handler) GetSingleOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	OrderPayment := new(models.RefOrderPayment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	OrderPayment, err := h.Service.GetSingleOrderPayment(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OrderPayment"] = OrderPayment
	response.With200V2(w, "Success", m, platform)
}

// UpdateOrderPayment : ""
func (h *Handler) UpdateOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	OrderPayment := new(models.OrderPayment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&OrderPayment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if OrderPayment.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateOrderPayment(ctx, OrderPayment)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OrderPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableOrderPayment : ""
func (h *Handler) EnableOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableOrderPayment(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OrderPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableOrderPayment : ""
func (h *Handler) DisableOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableOrderPayment(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OrderPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteOrderPayment : ""
func (h *Handler) DeleteOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteOrderPayment(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OrderPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterOrderPayment : ""
func (h *Handler) FilterOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.OrderPaymentFilter
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

	var OrderPayments []models.RefOrderPayment
	log.Println(pagination)
	OrderPayments, err = h.Service.FilterOrderPayment(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(OrderPayments) > 0 {
		m["OrderPayment"] = OrderPayments
	} else {
		res := make([]models.OrderPayment, 0)
		m["OrderPayment"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// MakePayment : ""
func (h *Handler) MakePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	OrderPayment := new(models.OrderPayment)
	err := json.NewDecoder(r.Body).Decode(&OrderPayment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.MakePayment(ctx, OrderPayment)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["OrderPayment"] = OrderPayment
	response.With200V2(w, "Success", m, platform)
}
