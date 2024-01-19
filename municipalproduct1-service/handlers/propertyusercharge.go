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

// SavePropertyUserCharge : ""
func (h *Handler) SavePropertyUserCharge(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	propertyusercharge := new(models.PropertyUserCharge)
	err := json.NewDecoder(r.Body).Decode(&propertyusercharge)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyUserCharge(ctx, propertyusercharge)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyusercharge"] = propertyusercharge
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePropertyUserCharge : ""
func (h *Handler) GetSinglePropertyUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	propertyusercharge := new(models.RefPropertyUserCharge)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	propertyusercharge, err := h.Service.GetSinglePropertyUserCharge(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyusercharge"] = propertyusercharge
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyUserCharge : ""
func (h *Handler) UpdatePropertyUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	propertyusercharge := new(models.PropertyUserCharge)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&propertyusercharge)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if propertyusercharge.PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyUserCharge(ctx, propertyusercharge)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyusercharge"] = propertyusercharge
	response.With200V2(w, "Success", m, platform)
}

//EnablePropertyUserCharge : ""
func (h *Handler) EnablePropertyUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnablePropertyUserCharge(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyusercharge"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisablePropertyUserCharge : ""
func (h *Handler) DisablePropertyUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisablePropertyUserCharge(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyusercharge"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeletePropertyUserCharge : ""
func (h *Handler) DeletePropertyUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var vp models.UserChargeAction
	if err := json.NewDecoder(r.Body).Decode(&vp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := h.Service.DeletePropertyUserCharge(ctx, &vp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}

	m := make(map[string]interface{})
	m["verifiedProperty"] = "Suceess"
	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyUserCharge : ""
func (h *Handler) FilterPropertyUserCharge(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PropertyUserChargeFilter
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

	var PropertyUserCharges []models.RefPropertyUserCharge
	log.Println(pagination)
	PropertyUserCharges, err = h.Service.FilterPropertyUserCharge(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(PropertyUserCharges) > 0 {
		m["propertyusercharge"] = PropertyUserCharges
	} else {
		res := make([]models.RefPropertyUserCharge, 0)
		m["propertyusercharge"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// VerifyPayment : ""
func (h *Handler) VerifyPropertyUserCharge(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var vp models.UserChargeAction
	if err := json.NewDecoder(r.Body).Decode(&vp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	propertyId, err := h.Service.VerifyPropertyUserCharge(ctx, &vp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}

	m := make(map[string]interface{})
	m["verifiedProperty"] = propertyId
	response.With200V2(w, "Success", m, platform)
}

// RejectPayment : ""
func (h *Handler) RejectPropertyUserCharge(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var rp models.UserChargeAction
	if err := json.NewDecoder(r.Body).Decode(&rp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	//	fmt.Println("tnx id - ", rp.TnxID)
	propertyId, err := h.Service.RejectPropertyUserCharge(ctx, &rp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}

	m := make(map[string]interface{})
	m["verifiedProperty"] = propertyId
	response.With200V2(w, "Success", m, platform)
}
