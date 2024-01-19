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

//InitiateShopRentPayment : ""
func (h *Handler) InitiateShopRentPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	initiate := new(models.InitiateShopRentPaymentReq)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&initiate)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	tnxId, err := h.Service.InitiateShopRentPayment(ctx, initiate)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["initiatePayment"] = "success"
	m["tnxId"] = tnxId
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetSingleShopRentPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	tnxID := r.URL.Query().Get("tnxId")
	if tnxID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	payment, err := h.Service.GetSingleShopRentPayment(ctx, tnxID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payment"] = payment
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) MakeShopRentPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	var mmtpr *models.MakeShopRentPaymentReq
	err := json.NewDecoder(r.Body).Decode(&mmtpr)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	shopRentID, err := h.Service.MakeShopRentPayment(ctx, mmtpr)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	err = h.Service.UpdateShopRentDemandAndCollections(ctx, shopRentID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["makePayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterShopRentPayment : ""
func (h *Handler) FilterShopRentPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ShopRentPaymentsFilter
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
	payments, err := h.Service.FilterShopRentPayment(ctx, filter, pagination)
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

//GetShopRentPaymentReceiptsPDF : ""
func (h *Handler) GetShopRentPaymentReceiptsPDF(w http.ResponseWriter, r *http.Request) {
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
	data, err := h.Service.GetShopRentPaymentReceiptsPDF(ctx, ID)
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

// BouncePayment : ""
func (h *Handler) ShoprentBouncePayment(w http.ResponseWriter, r *http.Request) {
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
	propertyId, err := h.Service.ShoprentBouncePayment(ctx, &bp)
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
	m["bouncedPayment"] = propertyId
	response.With200V2(w, "Success", m, platform)
}
