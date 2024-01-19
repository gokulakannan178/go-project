package handlers

import (
	"encoding/json"
	"fmt"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

// InitiatePropertyOtherDemandPayment : ""
func (h *Handler) InitiatePropertyOtherDemandPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	filter := new(models.InitiatePropertyOtherDemandFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	txnID, err := h.Service.InitiatePropertyOtherDemandPayment(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["txtnId"] = txnID
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePropertyOtherDemandPaymentTxtID : ""
func (h *Handler) GetSinglePropertyOtherDemandPaymentTxtID(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.GetSinglePropertyOtherDemandPaymentTxtID(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyOtherDemandPayment"] = data
	response.With200V2(w, "Success", m, platform)
}

// PropertyOtherDemandMakePayment : ""
func (h *Handler) PropertyOtherDemandMakePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	payment := new(models.PropertyOtherDemandMakePayment)
	err := json.NewDecoder(r.Body).Decode(&payment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	propertyId, err := h.Service.PropertyOtherDemandMakePayment(ctx, payment)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if propertyId == "" {
		response.With500mV2(w, "failed to get property id- ", platform)
		return
	}
	// if payment.Details.MOP.Mode != constants.MOPCHEQUE {
	// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
	// 	defer ctx.Client.Disconnect(r.Context())
	// 	err = h.Service.SavePropertyDemand(ctx, propertyId)
	// 	if err != nil {
	// 		response.With500mV2(w, "failed - "+err.Error(), platform)
	// 		return
	// 	}
	// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
	// 	defer ctx.Client.Disconnect(r.Context())
	// 	err = h.Service.PropertyUpdateCollection(ctx, propertyId)
	// 	if err != nil {
	// 		response.With500mV2(w, "failed - "+err.Error(), platform)
	// 		return
	// 	}
	// }
	m := make(map[string]interface{})
	m["payment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// PropertyOtherDemandVerifyPayment : ""
func (h *Handler) PropertyOtherDemandVerifyPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var vp models.PropertyOtherDemandVerifyPayment
	if err := json.NewDecoder(r.Body).Decode(&vp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	_, err := h.Service.PropertyOtherDemandVerifyPayment(ctx, &vp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}

	// ctx = app.GetApp(r.Context(), h.Service.Daos)
	// defer ctx.Client.Disconnect(r.Context())
	// err = h.Service.SavePropertyDemand(ctx, propertyId)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	// ctx = app.GetApp(r.Context(), h.Service.Daos)
	// defer ctx.Client.Disconnect(r.Context())
	// err = h.Service.PropertyUpdateCollection(ctx, propertyId)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }

	m := make(map[string]interface{})
	m["verifiedProperty"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// PropertyOtherDemandNotVerifiedPayment : ""
func (h *Handler) PropertyOtherDemandNotVerifiedPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var vp models.PropertyOtherDemandNotVerifiedPayment
	if err := json.NewDecoder(r.Body).Decode(&vp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))

		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := h.Service.PropertyOtherDemandNotVerifiedPayment(ctx, &vp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}
	m := make(map[string]interface{})
	m["notVerifiedProperty"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// PropertyOtherDemandRejectPayment : ""
func (h *Handler) PropertyOtherDemandRejectPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var rp models.PropertyOtherDemandRejectPayment
	if err := json.NewDecoder(r.Body).Decode(&rp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	fmt.Println("tnx id - ", rp.TnxID)
	_, err := h.Service.PropertyOtherDemandRejectPayment(ctx, &rp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}

	// ctx = app.GetApp(r.Context(), h.Service.Daos)
	// defer ctx.Client.Disconnect(r.Context())
	// err = h.Service.SavePropertyDemand(ctx, propertyId)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	// ctx = app.GetApp(r.Context(), h.Service.Daos)
	// defer ctx.Client.Disconnect(r.Context())
	// err = h.Service.PropertyUpdateCollection(ctx, propertyId)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }

	m := make(map[string]interface{})
	m["rejectedProperty"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// GetPropertyOtherDemandPaymentReceiptsPDF : ""
func (h *Handler) GetPropertyOtherDemandPaymentReceiptsPDF(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID
	// data, err := h.Service.GetPaymentReceiptsPDF(ctx, ID)
	// if err != nil {
	// 	if err.Error() == "mongo: no documents in result" {
	// 		response.With500mV2(w, "failed no data for this id", platform)
	// 		return
	// 	}
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
	// w.Write(data)

}
