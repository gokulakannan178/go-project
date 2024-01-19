package handlers

import (
	"encoding/json"
	"logikoof-echalan-service/app"
	"logikoof-echalan-service/models"
	"logikoof-echalan-service/response"
	"net/http"
)

//PaymentWidget : ""
func (h *Handler) PaymentWidget(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	paymentWidgetFilter := new(models.PaymentWidgetFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&paymentWidgetFilter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	paymentWidget, err := h.Service.PaymentWidget(ctx, paymentWidgetFilter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if paymentWidget == nil {
		paymentWidget = new(models.PaymentWidget)
	}
	m := make(map[string]interface{})
	m["paymentWidget"] = paymentWidget
	response.With200V2(w, "Success", m, platform)
}

//TodaysOffenceWidget : ""
func (h *Handler) TodaysOffenceWidget(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	filter := new(models.TodaysOffenceWidgetFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	todaysOffenceWidget, err := h.Service.TodaysOffenceWidget(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if todaysOffenceWidget == nil {
		todaysOffenceWidget = new(models.TodaysOffenceWidget)
	}
	m := make(map[string]interface{})
	m["todaysOffenceWidget"] = todaysOffenceWidget
	response.With200V2(w, "Success", m, platform)
}

//TopOffencesWidget : ""
func (h *Handler) TopOffencesWidget(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	filter := new(models.TopOffencesWidgetFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	todaysOffenceWidgets, err := h.Service.TopOffencesWidget(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["topOffencesWidgets"] = todaysOffenceWidgets
	response.With200V2(w, "Success", m, platform)
}
