package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

// GetUserChargeSAFDashboard : ""
func (h *Handler) GetUserChargeSAFDashboard(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	filter := new(models.GetUserChargeSAFDashboardFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	result, err := h.Service.GetUserChargeSAFDashboard(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tlSafResult"] = result
	response.With200V2(w, "Success", m, platform)
}

// GetUserChargeSAFDashboard : ""
func (h *Handler) UserwiseUserChargeReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.UserFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if resType == "excel" {
		file, err := h.Service.UserwiseUserChargeReportExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=userwiseusercharge.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	result, err := h.Service.UserwiseUserChargeReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["tlSafResult"] = result
	response.With200V2(w, "Success", m, platform)
}
