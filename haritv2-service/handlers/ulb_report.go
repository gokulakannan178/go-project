package handlers

import (
	"encoding/json"
	"haritv2-service/app"
	"haritv2-service/models"
	"haritv2-service/response"
	"net/http"
)

//FPOMasterReport : ""
func (h *Handler) ULBMasterReportV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.ULBMasterReportV2Filter)

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.ULBMasterReportV2Excel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	data, err := h.Service.ULBMasterReportV2Json(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbmasterreport"] = data
	response.With200V2(w, "Success", m, platform)

}
