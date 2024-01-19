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

//SaveartPayment : ""
func (h *Handler) SavePropertyPartPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	propertyPayment := new(models.PropertyPayment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&propertyPayment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.MakePaymentV2(ctx, propertyPayment)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	propertyID, err := h.Service.ValidatePartPayments(ctx, propertyPayment.TnxID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	err = h.Service.ValidateMainPayment(ctx, propertyID)
	if err != nil {
		response.With500mV2(w, "failed ValidateMainPayment- "+err.Error(), platform)
		return
	}
	// err = h.Service.UpdatePropertyPartPaymentDemandAndCollections(ctx, shoprent.UniqueID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["propertyPayment"] = propertyPayment
	response.With200V2(w, "Success", m, platform)
}

//SaveartPayment : ""
func (h *Handler) SavePropertyPartPaymentAdditional(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	propertyPayment := new(models.PropertyPayment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&propertyPayment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.MakePaymentV2AdditionalPayment(ctx, propertyPayment)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	propertyID, err := h.Service.ValidatePartPayments(ctx, propertyPayment.TnxID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	err = h.Service.ValidateMainPayment(ctx, propertyID)
	if err != nil {
		response.With500mV2(w, "failed ValidateMainPayment- "+err.Error(), platform)
		return
	}
	// err = h.Service.UpdatePropertyPartPaymentDemandAndCollections(ctx, shoprent.UniqueID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["propertyPayment"] = propertyPayment
	response.With200V2(w, "Success", m, platform)
}

//SaveartPayment : ""
func (h *Handler) ValidatePartPayments(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var TnxID string
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&TnxID)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	propertyID, err := h.Service.ValidatePartPayments(ctx, TnxID)
	if err != nil {
		response.With500mV2(w, "failed ValidatePartPayments- "+err.Error(), platform)
		return
	}
	err = h.Service.ValidateMainPayment(ctx, propertyID)
	if err != nil {
		response.With500mV2(w, "failed ValidateMainPayment- "+err.Error(), platform)
		return
	}

	// err = h.Service.UpdatePropertyPartPaymentDemandAndCollections(ctx, shoprent.UniqueID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["partPayment Validated"] = TnxID
	response.With200V2(w, "Success", m, platform)
}

//VerifyPropertyPartPayment : ""
func (h *Handler) VerifyPropertyPartPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakePropertyPartPaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	_, err = h.Service.VerifyPropertyPartPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	propertyID, err := h.Service.ValidatePartPayments(ctx, action.TnxID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	err = h.Service.ValidateMainPayment(ctx, propertyID)
	if err != nil {
		response.With500mV2(w, "failed ValidateMainPayment- "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["verifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//NotVerifyPropertyPartPayment : ""
func (h *Handler) NotVerifyPropertyPartPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakePropertyPartPaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	_, err = h.Service.NotVerifyPropertyPartPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	propertyID, err := h.Service.ValidatePartPayments(ctx, action.TnxID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	err = h.Service.ValidateMainPayment(ctx, propertyID)
	if err != nil {
		response.With500mV2(w, "failed ValidateMainPayment- "+err.Error(), platform)
		return
	}

	// err = h.Service.UpdatePropertyPartDemandAndCollections(ctx, shoprentID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["notVerifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//RejectPropertyPartPayment : ""
func (h *Handler) RejectPropertyPartPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakePropertyPartPaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	_, err = h.Service.RejectPropertyPartPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	propertyID, err := h.Service.ValidatePartPayments(ctx, action.TnxID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	err = h.Service.ValidateMainPayment(ctx, propertyID)
	if err != nil {
		response.With500mV2(w, "failed ValidateMainPayment- "+err.Error(), platform)
		return
	}

	// err = h.Service.UpdatePropertyPartDemandAndCollections(ctx, shoprentID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["rejectPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterPropertyPartPayment : ""
func (h *Handler) FilterPropertyPartPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PropertyPartPaymentFilter
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

	var partPayments []models.RefPropertyPartPayment
	log.Println(pagination)
	partPayments, err = h.Service.FilterPropertyPartPayment(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(partPayments) > 0 {
		m["data"] = partPayments
	} else {
		res := make([]models.PropertyWallet, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// GetPropertyPaymentsWithPartPayments :""
func (h *Handler) GetPropertyPaymentsWithPartPayments(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	tnxID := r.URL.Query().Get("id")
	if tnxID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.GetPropertyPaymentsWithPartPayments(ctx, tnxID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyPayments"] = data
	response.With200V2(w, "Success", m, platform)
}
