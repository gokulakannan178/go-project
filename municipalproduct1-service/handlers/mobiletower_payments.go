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

//InitiateMobileTowerPayment : ""
func (h *Handler) InitiateMobileTowerPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	imtpr := new(models.InitiateMobileTowerPaymentReq)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&imtpr)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	tnxId, err := h.Service.InitiateMobileTowerPayment(ctx, imtpr)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["initiatePayment"] = "success"
	m["tnxId"] = tnxId
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetSingleMobileTowerPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	tnxID := r.URL.Query().Get("tnxId")
	if tnxID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	payment, err := h.Service.GetSingleMobileTowerPayment(ctx, tnxID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payment"] = payment
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) MakeMobileTowerPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	var mmtpr *models.MakeMobileTowerPaymentReq
	err := json.NewDecoder(r.Body).Decode(&mmtpr)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.MakeMobileTowerPayment(ctx, mmtpr)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["makePayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterMobileTowerPayment : ""
func (h *Handler) FilterMobileTowerPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.MobileTowerPaymentsFilter
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
	payments, err := h.Service.FilterMobileTowerPayment(ctx, filter, pagination)
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

//VerifyMobileTowerPayment : ""
func (h *Handler) VerifyMobileTowerPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeMobileTowerPaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.VerifyMobileTowerPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["verifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//NotVerifyMobileTowerPayment : ""
func (h *Handler) NotVerifyMobileTowerPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeMobileTowerPaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.NotVerifyMobileTowerPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["notVerifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//RejectMobileTowerPayment : ""
func (h *Handler) RejectMobileTowerPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeMobileTowerPaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.RejectMobileTowerPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["rejectPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetMobileTowerPaymentReceiptsPDF : ""
func (h *Handler) GetMobileTowerPaymentReceiptsPDF(w http.ResponseWriter, r *http.Request) {
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
	data, err := h.Service.GetMobileTowerPaymentReceiptsPDF(ctx, ID)
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
