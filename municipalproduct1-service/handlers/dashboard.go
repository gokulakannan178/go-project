package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

func (h *Handler) DashboardDemandAndCollection(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ddac models.DashboardDemandAndCollectionFilter
	if err := json.NewDecoder(r.Body).Decode(&ddac); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	res, err := h.Service.DashboardDemandAndCollection(ctx, &ddac)
	if err != nil {
		response.With500mV2(w, "failed no data in this id "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ddac"] = res
	response.With200V2(w, "Success", m, platform)

}

func (h *Handler) DashboardDemandAndCollectionV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ddac models.DashboardDemandAndCollectionFilter
	if err := json.NewDecoder(r.Body).Decode(&ddac); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	res, err := h.Service.DashboardDemandAndCollectionV2(ctx, &ddac)
	if err != nil {
		response.With500mV2(w, "failed no data in this id "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ddac"] = res
	response.With200V2(w, "Success", m, platform)

}
