package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//PaytmPaymentInitiateTransaction : ""
func (h *Handler) PaytmPaymentInitiateTransaction(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	uppi := new(models.UserPaytmPaymentInit)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&uppi)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.PaytmPaymentInitiateTransaction(ctx, uppi)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["uppi"] = data
	response.With200V2(w, "Success", m, platform)
}

//PaytmPaymentChecksum : ""
func (h *Handler) PaytmPaymentChecksum(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var body interface{}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.CreateChecksum(ctx, body)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["uppi"] = data
	response.With200V2(w, "Success", m, platform)
}
