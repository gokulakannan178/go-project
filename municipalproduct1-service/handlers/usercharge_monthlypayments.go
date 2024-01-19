package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

// InitiateUserChargePayment : ""
func (h *Handler) InitiateUserChargeMonthlyPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	initiate := new(models.InitiateUserChargeMonthlyPaymentReq)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&initiate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	tnxId, err := h.Service.InitiateUserChargeMonthlyPayment(ctx, initiate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["initiatePayment"] = "success"
	m["tnxId"] = tnxId
	response.With200V2(w, "Success", m, platform)
}

// GetUserChargePaymentReceiptsPDF : ""
func (h *Handler) GetUserChargeMonthlyPaymentReceiptsPDF(w http.ResponseWriter, r *http.Request) {
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
	data, err := h.Service.GetUserChargeMonthlyPaymentReceiptsPDF(ctx, ID)
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
	w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")

}

// GetUserChargePaymentReceiptsPDF : ""
func (h *Handler) GetUserChargeMonthlyPayment(w http.ResponseWriter, r *http.Request) {
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
	data, err := h.Service.GetSingleUserChargeMonthlyPayment(ctx, ID)
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

// FilterUserChargePayment : ""
func (h *Handler) FilterUserChargeMonthlyPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.UserChargeMonthlyPaymentsFilter
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

	if resType == "excel" {
		file, err := h.Service.FilterUserChargeMonthlyPaymentExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=userchargepayment.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	log.Println(pagination)
	payments, err := h.Service.FilterUserChargeMonthlyPayment(ctx, filter, pagination)
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

func (h *Handler) MakeUserChargePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	var mmtpr *models.MakeUserChargePaymentReq
	err := json.NewDecoder(r.Body).Decode(&mmtpr)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	// tradeLicenseID, err := h.Service.MakeTradeLicensePayment(ctx, mmtpr)
	_, err = h.Service.MakeUserChargePayment(ctx, mmtpr)
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
	m["makePayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// BouncePayment : ""
func (h *Handler) UserChargeBouncePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var bp models.BouncePayment
	if err := json.NewDecoder(r.Body).Decode(&bp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	fmt.Println("tnx id - ", bp.TnxID)
	propertyId, err := h.Service.UserChargeBouncePayment(ctx, &bp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyDemand(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.PropertyUpdateCollection(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["bouncedPayment"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// VerifyPayment : ""
func (h *Handler) UserChargVerifyPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var vp models.VerifyPayment
	if err := json.NewDecoder(r.Body).Decode(&vp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	propertyId, err := h.Service.UserChargVerifyPayment(ctx, &vp)
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
	m["verifiedProperty"] = propertyId
	response.With200V2(w, "Success", m, platform)
}

// NotVerifiedPayment : ""
func (h *Handler) UserChargNotVerifiedPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var vp models.NotVerifiedPayment
	if err := json.NewDecoder(r.Body).Decode(&vp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))

		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := h.Service.UserChargNotVerifiedPayment(ctx, &vp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}
	m := make(map[string]interface{})
	m["notVerifiedProperty"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// RejectPayment : ""
func (h *Handler) UserChargRejectPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var rp models.RejectPayment
	if err := json.NewDecoder(r.Body).Decode(&rp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	fmt.Println("tnx id - ", rp.TnxID)
	propertyId, err := h.Service.UserChargRejectPayment(ctx, &rp)
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
	m["verifiedProperty"] = propertyId
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetSingleUserChargePaymentTxtID(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.GetSingleUserChargePaymentTxtID(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyPayment"] = data
	response.With200V2(w, "Success", m, platform)
}

//DateRangeWiseTradeLisencePaymentReport : ""
func (h *Handler) DateRangeWiseUserchargePaymentReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	filter := new(models.DateWiseUserchargeReportFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	report, err := h.Service.DateRangeWiseUserchargePaymentReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["report"] = report
	response.With200V2(w, "Success", m, platform)
}
