package handlers

import (
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//SavePaymentGateway : ""
func (h *Handler) CalcShopRentMonthlyDemand(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	resType := r.URL.Query().Get("resType")
	if ID == "" {
		response.With400V2(w, "ID missing", platform)
		return
	}
	filter := new(models.ShopRentMonthlyCalcQueryFilter)
	filter.ShopRentID = ID
	filter.OmitPayedMonths = true
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	data, err := h.Service.CalcShopRentMonthlyDemandForParticulars(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	if resType == "pdf" {
		data, err := h.Service.SHopRentMonthlyGetOutstandingDemandPDF(ctx, ID)
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
		w.Header().Set("Content-Disposition", "attachment; filename=demandreceipt.pdf")

	}

	m := make(map[string]interface{})
	m["demand"] = data
	response.With200V2(w, "Success", m, platform)
}

//SavePaymentGateway : ""
func (h *Handler) CalcShopRentMonthlyDemandWithoutPaidYears(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	resType := r.URL.Query().Get("resType")
	if ID == "" {
		response.With400V2(w, "ID missing", platform)
		return
	}
	filter := new(models.ShopRentMonthlyCalcQueryFilter)
	filter.ShopRentID = ID
	filter.OmitPayedMonths = true
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	data, err := h.Service.CalcShopRentMonthlyDemandForParticulars(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	if resType == "pdf" {
		data, err := h.Service.SHopRentMonthlyGetOutstandingDemandPDF(ctx, ID)
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
		w.Header().Set("Content-Disposition", "attachment; filename=demandreceipt.pdf")

	}

	m := make(map[string]interface{})
	m["demand"] = data
	response.With200V2(w, "Success", m, platform)
}
