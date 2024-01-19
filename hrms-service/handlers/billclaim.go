package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

//SaveBillClaim : ""
func (h *Handler) SaveBillClaim(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	billClaim := new(models.BillClaim)
	version := r.URL.Query().Get("version")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&billClaim)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	if version == "V2" {
		err = h.Service.SaveBillClaimV2(ctx, billClaim)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		m := make(map[string]interface{})
		m["BillClaim"] = billClaim
		response.With200V2(w, "Success", m, platform)
		return
	}
	err = h.Service.SaveBillClaim(ctx, billClaim)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillClaim"] = billClaim
	response.With200V2(w, "Success", m, platform)
}

//UpdateBillClaim :""
func (h *Handler) UpdateBillClaim(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	billClaim := new(models.BillClaim)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&billClaim)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if billClaim.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateBillClaim(ctx, billClaim)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableBillClaim : ""
func (h *Handler) EnableBillClaim(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableBillClaim(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableBillClaim : ""
func (h *Handler) DisableBillClaim(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableBillClaim(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteBillClaim : ""
func (h *Handler) DeleteBillClaim(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteBillClaim(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleBillClaim :""
func (h *Handler) GetSingleBillClaim(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	billClaim := new(models.RefBillClaim)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	billClaim, err := h.Service.GetSingleBillClaim(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = billClaim
	response.With200V2(w, "Success", m, platform)
}

//FilterBillClaim : ""
func (h *Handler) FilterBillClaim(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var billClaim *models.FilterBillClaim
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
	err := json.NewDecoder(r.Body).Decode(&billClaim)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var billClaims []models.RefBillClaim
	log.Println(pagination)
	billClaims, err = h.Service.FilterBillClaim(ctx, billClaim, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(billClaims) > 0 {
		m["data"] = billClaims
	} else {
		res := make([]models.BillClaim, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) ApprovedBillClaim(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	billClaim := new(models.ReviewedBillClaim)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&billClaim)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if billClaim.BillClaim == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.ApprovedBillClaim(ctx, billClaim)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) RejectedBillClaim(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	billClaim := new(models.ReviewedBillClaim)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&billClaim)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if billClaim.BillClaim == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.RejectedBillClaim(ctx, billClaim)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
