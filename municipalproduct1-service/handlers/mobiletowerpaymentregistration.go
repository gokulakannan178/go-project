package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//InitiateMobileTowerPayment : ""
func (h *Handler) InitiateMobileTowerRegisterPayment(w http.ResponseWriter, r *http.Request) {
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
	tnxId, err := h.Service.InitiateMobileTowerRegisterPayment(ctx, imtpr)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["initiatePayment"] = "success"
	m["tnxId"] = tnxId
	response.With200V2(w, "Success", m, platform)
}

// MakeMobileTowerPaymentForRegistration : ""
func (h *Handler) MakeMobileTowerPaymentForRegistration(w http.ResponseWriter, r *http.Request) {
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
	err = h.Service.MakeMobileTowerPaymentForRegistration(ctx, mmtpr)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["makePaymentForRegistration"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//VerifyMobileTowerRegistrationPayment : ""
func (h *Handler) VerifyMobileTowerRegistrationPayment(w http.ResponseWriter, r *http.Request) {
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
	err = h.Service.VerifyMobileTowerRegistrationPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["verifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//NotVerifyMobileTowerRegistrationPayment : ""
func (h *Handler) NotVerifyMobileTowerRegistrationPayment(w http.ResponseWriter, r *http.Request) {
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
	err = h.Service.NotVerifyMobileTowerRegistrationPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["notVerifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//RejectMobileTowerRegistrationPayment : ""
func (h *Handler) RejectMobileTowerRegistrationPayment(w http.ResponseWriter, r *http.Request) {
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
	err = h.Service.RejectMobileTowerRegistrationPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["rejectPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetMobileTowerRegistartionPaymentReceiptsPDF : ""
func (h *Handler) GetMobileTowerRegistartionPaymentReceiptsPDF(w http.ResponseWriter, r *http.Request) {
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
	data, err := h.Service.GetMobileTowerRegistartionPaymentReceiptsPDF(ctx, ID)
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
