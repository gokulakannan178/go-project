package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//InitiateShopRentMonthlyLegacyPayment : ""
func (h *Handler) InitiateShopRentMonthlyLegacyPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	initiate := new(models.InitiateShopRentMonthlyPaymentReq)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&initiate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	tnxId, err := h.Service.InitiateShopRentMonthlyLegacyPayment(ctx, initiate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["initiatemonthlypayment"] = "success"
	m["tnxId"] = tnxId
	response.With200V2(w, "Success", m, platform)
}

//MakeShopRentMonthlyLegacyPayment : ""
func (h *Handler) MakeShopRentMonthlyLegacyPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	var mmtpr *models.MakeShopRentPaymentReq
	err := json.NewDecoder(r.Body).Decode(&mmtpr)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	shopRentID, err := h.Service.MakeShopRentMonthlyLegacyPayment(ctx, mmtpr)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	err = h.Service.UpdateShopRentDemandAndCollections(ctx, shopRentID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["makePayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}
