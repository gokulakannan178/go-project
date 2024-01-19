package handlers

import (
	"encoding/json"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

// InitiateSolidWasteUserChargePayment : ""
func (h *Handler) InitiateSolidWasteUserChargeMonthlyPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	initiate := new(models.InitiateSolidWasteChargeMonthlyPaymentReq)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&initiate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	tnxId, err := h.Service.InitiateSolidWasteChargeMonthlyPayment(ctx, initiate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["initiatePayment"] = "success"
	m["filter.SolidWasteChargeID"] = initiate.SolidWasteChargeID
	m["tnxId"] = tnxId
	response.With200V2(w, "Success", m, platform)
}

// GetSingleSolidWasteUserChargePayment : ""
func (h *Handler) GetSingleSolidWasteUserChargePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("tnxId")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID
	data, err := h.Service.GetSingleSolidWasteUserChargePayment(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	m["payment"] = data
	response.With200V2(w, "Success", m, platform)
}

// MakeSolidWasteUserChargePayment : ""
func (h *Handler) MakeSolidWasteUserChargePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	var mmtpr *models.MakeSolidWasteUserChargePaymentReq
	err := json.NewDecoder(r.Body).Decode(&mmtpr)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	_, err = h.Service.MakeSolidWasteUserChargePayment(ctx, mmtpr)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	// err = h.Service.UpdateSolidWasteUserChargeDemandAndCollections(ctx, shopRentID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["makePayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterSolidWasteUserChargePayment : ""
func (h *Handler) FilterSolidWasteUserChargePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var filter *models.SolidWasteUserChargePaymentsFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	var pagination *models.Pagination
	if pageNo != "no" {
		pagination = new(models.Pagination)
		if pagination.PageNum = 1; pageNo != "" {
			page, err := strconv.Atoi(pageNo)
			if pagination.PageNum = 1; err == nil {
				pagination.PageNum = page
			}
		}
		if pagination.Limit = 10; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	log.Println(pagination)
	payments, err := h.Service.FilterSolidWasteUserChargePayment(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["payments"] = payments
	} else {
		res := make([]interface{}, 0)
		m["payments"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}
	response.With200V2(w, "Success", m, platform)
}

// GetSolidWasteUserChargePaymentReceiptsPDF : ""
func (h *Handler) GetSolidWasteUserChargePaymentReceiptsPDF(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("tnxId")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.GetSolidWasteUserChargePaymentReceiptsPDF(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
	w.Write(data)

}

// VerifySolidWasteUserChargePayment : ""
func (h *Handler) VerifySolidWasteUserChargePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeSolidWasteUserChargePaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	_, err = h.Service.VerifySolidWasteUserChargePayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	// err = h.Service.UpdateSolidWasteUserChargeDemandAndCollections(ctx, shoprentID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["verifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//NotVerifySolidWasteUserChargePayment : ""
func (h *Handler) NotVerifySolidWasteUserChargePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeSolidWasteUserChargePaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	_, err = h.Service.NotVerifySolidWasteUserChargePayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	// err = h.Service.UpdateSolidWasteUserChargeDemandAndCollections(ctx, shoprentID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["notVerifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//RejectSolidWasteUserChargePayment : ""
func (h *Handler) RejectSolidWasteUserChargePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeSolidWasteUserChargePaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	_, err = h.Service.RejectSolidWasteUserChargePayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	// err = h.Service.UpdateSolidWasteUserChargeDemandAndCollections(ctx, shoprentID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["rejectPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}
