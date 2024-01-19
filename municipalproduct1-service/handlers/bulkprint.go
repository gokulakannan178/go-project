package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//FilterBulkPrint : ""
func (h *Handler) BulkPrintGetDetailsForProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.BulkPrintFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var BulkPrints []models.BulkPrintDetail

	BulkPrints, err = h.Service.BulkPrintGetDetailsForProperty(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(BulkPrints) > 0 {
		m["results"] = BulkPrints
	} else {
		res := make([]models.BulkPrint, 0)
		m["results"] = res
	}

	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) BulkPrintGetDetailsReceiptsForProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.BulkPrintReceiptsRequest
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	data, err := h.Service.BulkPrintReceiptsRequestForProperty(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	w.Write(data)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
	return

}

//FilterBulkPrint : ""
func (h *Handler) BulkPrintGetDetailsForTradelicense(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.BulkPrintFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var BulkPrints []models.BulkPrintDetail

	BulkPrints, err = h.Service.BulkPrintGetDetailsForTradelicense(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(BulkPrints) > 0 {
		m["results"] = BulkPrints
	} else {
		res := make([]models.BulkPrint, 0)
		m["results"] = res
	}

	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) BulkPrintGetDetailsReceiptsForTradelicense(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.BulkPrintReceiptsRequest
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	data, err := h.Service.BulkPrintGetDetailsReceiptsForTradelicense(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	w.Write(data)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
	return

}
