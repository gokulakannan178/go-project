package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"net/http"
)

//EmployeeReport : ""
func (h *Handler) EmployeeReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	respType := r.URL.Query().Get("resType")
	filter := new(models.FilterEmployee)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	if respType == "excel" {
		file, err := h.Service.EmployeeReport(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=EmployeeReport.xlsx")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	data, err := h.Service.EmployeeReport(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["collection"] = data
	response.With200V2(w, "Success", m, platform)
}
