package handlers

import (
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

// TradeLicenseDemandPart2 : ""
func (h *Handler) TradeLicenseDemandPart2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	TradeLicenseID := r.URL.Query().Get("id")
	resType := r.URL.Query().Get("resType")
	fyID := r.URL.Query().Get("fy")
	tlPart2 := new(models.TradeLicenseDemandPart2Filter)
	if TradeLicenseID != "" {
		tlPart2.TradeLicenseID = TradeLicenseID
	}

	if fyID != "" {
		tlPart2.FyID = append(tlPart2.FyID, fyID)
	}
	if TradeLicenseID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	var tradeLicenseDemand *models.RefTradeLicenseDemandPart2

	tradeLicenseDemand, err := h.Service.TradeLicenseDemandCalcPart2(ctx, tlPart2)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if resType == "pdf" {
		data, err := h.Service.GetTradeLicensePaymentDemandPDFPart2(ctx, tlPart2)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				response.With500mV2(w, "failed no data for this id", platform)
				return
			}
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}

		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=tradelicensedemandpart2.pdf")

	}
	m := make(map[string]interface{})
	m["tradeLicenseDemand"] = tradeLicenseDemand
	response.With200V2(w, "Success", m, platform)
}
