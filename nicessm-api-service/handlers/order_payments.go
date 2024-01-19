package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SaveOrderPayments : ""
func (h *Handler) SaveOrderPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	payment := new(models.OrderPayment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&payment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveOrderPayment(ctx, payment)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payment"] = payment
	response.With200V2(w, "Success", m, platform)
}

//UpdateOrderPayments :""
func (h *Handler) UpdateOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	payment := new(models.OrderPayment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&payment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if payment.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateOrderPayment(ctx, payment)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableOrderPayments : ""
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
	m["payment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableOrderPayments : ""
func (h *Handler) DisableOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableOrderPayment(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteOrderPayments : ""
func (h *Handler) DeleteOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := new(models.OrderPayment)
	UniqueID := r.URL.Query().Get("id")

	if ID.ID != primitive.NilObjectID {
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
	m["payment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleOrderPayments :""
func (h *Handler) GetSingleOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	payment := new(models.RefOrderPayment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	payment, err := h.Service.GetSingleOrderPayment(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payment"] = payment
	response.With200V2(w, "Success", m, platform)
}

//FilterOrderPayments : ""
func (h *Handler) FilterOrderPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var payment *models.OrderPaymentFilter
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
	err := json.NewDecoder(r.Body).Decode(&payment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var payments []models.RefOrderPayment
	log.Println(pagination)
	payments, err = h.Service.FilterOrderPayment(ctx, payment, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["payment"] = payments
	} else {
		res := make([]models.OrderPayment, 0)
		m["payment"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
