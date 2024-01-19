package handlers

import (
	"haritv2-service/app"
	"haritv2-service/models"
	"haritv2-service/response"
	"net/http"
)

//CreateFPOPurchaseULBSale : ""
func (h *Handler) CreateFPOPurchaseULBSale(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	orderID := r.URL.Query().Get("id")

	if orderID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.CreateFPOPurchaseULBSale(ctx, orderID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpoPurchaseSale"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) PlaceAndDeliverOrder(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	orderID := r.URL.Query().Get("id")

	if orderID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.PlaceAndDeliverOrder(ctx, orderID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpoPurchaseSale"] = "success"
	response.With200V2(w, "Success", m, platform)
}
