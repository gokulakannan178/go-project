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

//InitiateTradeLicensePayment : ""
func (h *Handler) InitiateTradeLicensePaymentPart2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	initiate := new(models.TradeLicenseDemandPart2Filter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&initiate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	tnxId, err := h.Service.InitiateTradeLicensePaymentPart2(ctx, initiate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["initiatePayment"] = "success"
	m["tnxId"] = tnxId
	response.With200V2(w, "Success", m, platform)
}

// GetSingleTradeLicensePaymentPart2 : ""
func (h *Handler) GetSingleTradeLicensePaymentPart2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	tnxID := r.URL.Query().Get("tnxId")
	if tnxID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	payment, err := h.Service.GetSingleTradeLicensePaymentPart2(ctx, tnxID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	if payment != nil {
		m["tradeLicensePaymentPart2"] = payment
	}
	response.With200V2(w, "Success", m, platform)
}

// MakeTradeLicensePaymentPart2 : ""
func (h *Handler) MakeTradeLicensePaymentPart2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	var mtlprp2 *models.MakeTradeLicensePaymentReqPart2
	err := json.NewDecoder(r.Body).Decode(&mtlprp2)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	// tradeLicenseID, err := h.Service.MakeTradeLicensePayment(ctx, mtlprp2)
	_, err = h.Service.MakeTradeLicensePaymentPart2(ctx, mtlprp2)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	// err = h.Service.UpdateTradeLicenseDemandAndCollections(ctx, tradeLicenseID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["makeTradeLicensePaymentPart2"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterTradeLicensePaymentPart2 : ""
func (h *Handler) FilterTradeLicensePaymentPart2(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.TradeLicensePaymentsFilterPart2
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
	payments, err := h.Service.FilterTradeLicensePaymentPart2(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["tradelicensePaymentsPart2"] = payments
	} else {
		res := make([]interface{}, 0)
		m["tradelicensePaymentsPart2"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// GetTradeLicensePaymentReceiptsPDFPart2 : ""
func (h *Handler) GetTradeLicensePaymentReceiptsPDFPart2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("tnxId")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	// filter := new(models.PropertyDemandFilter)
	// filter.PropertyID = ID
	data, err := h.Service.GetTradeLicensePaymentReceiptsPDFPart2(ctx, ID)
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

//GetTradeLicensePaymentReceiptsPDFV2Part2 : ""
func (h *Handler) GetTradeLicensePaymentReceiptsPDFV2Part2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID
	data, err := h.Service.GetTradeLicensePaymentReceiptsPDFV2Part2(ctx, ID)
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

//VerifyTradeLicensePayment : ""
func (h *Handler) VerifyTradeLicensePaymentPart2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeTradeLicensePaymentsActionPart2)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.VerifyTradeLicensePaymentPart2(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["verifyPaymentPart2"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// NotVerifyTradeLicensePaymentPart2 : ""
func (h *Handler) NotVerifyTradeLicensePaymentPart2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeTradeLicensePaymentsActionPart2)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.NotVerifyTradeLicensePaymentPart2(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["notVerifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// RejectTradeLicensePaymentPart2 : ""
func (h *Handler) RejectTradeLicensePaymentPart2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeTradeLicensePaymentsActionPart2)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.RejectTradeLicensePaymentPart2(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["rejectPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleTradeLicenseV2Part2 : ""
func (h *Handler) GetSingleTradeLicenseV2Part2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	data, err := h.Service.GetSingleTradeLicenseV2Part2(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	if data != nil {
		m["tradeLicensePaymentPart2"] = data
	}
	response.With200V2(w, "Success", m, platform)
}

// BasicTradeLicenseUpdateGetPaymentsToBeUpdatedPart2 : ""
func (h *Handler) BasicTradeLicenseUpdateGetPaymentsToBeUpdatedPart2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	rbtlulp2 := new(models.RefBasicTradeLicenseUpdateLogV2Part2)
	err := json.NewDecoder(r.Body).Decode(&rbtlulp2)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	payments, err := h.Service.BasicTradeLicenseUpdateGetPaymentsToBeUpdatedPart2(ctx, rbtlulp2)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["paymentsPart2"] = payments
	if len(payments) < 1 {
		m["paymentsPart2"] = []interface{}{}

	}
	response.With200V2(w, "Success", m, platform)
}
