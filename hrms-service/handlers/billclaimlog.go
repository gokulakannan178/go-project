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

//SaveBillClaimLog : ""
func (h *Handler) SaveBillClaimLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	billClaimLog := new(models.BillClaimLog)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&billClaimLog)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveBillClaimLog(ctx, billClaimLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillClaimLog"] = billClaimLog
	response.With200V2(w, "Success", m, platform)
}

//UpdateBillClaimLog :""
func (h *Handler) UpdateBillClaimLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	billClaimLog := new(models.BillClaimLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&billClaimLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if billClaimLog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateBillClaimLog(ctx, billClaimLog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableBillClaimLog : ""
func (h *Handler) EnableBillClaimLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableBillClaimLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableBillClaimLog : ""
func (h *Handler) DisableBillClaimLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableBillClaimLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteBillClaimLog : ""
func (h *Handler) DeleteBillClaimLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteBillClaimLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleBillClaimLog :""
func (h *Handler) GetSingleBillClaimLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	billClaimLog := new(models.RefBillClaimLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	billClaimLog, err := h.Service.GetSingleBillClaimLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = billClaimLog
	response.With200V2(w, "Success", m, platform)
}

//FilterBillClaimLog : ""
func (h *Handler) FilterBillClaimLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var billClaimLog *models.FilterBillClaimLog
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
	err := json.NewDecoder(r.Body).Decode(&billClaimLog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var billClaimLogs []models.RefBillClaimLog
	log.Println(pagination)
	billClaimLogs, err = h.Service.FilterBillClaimLog(ctx, billClaimLog, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(billClaimLogs) > 0 {
		m["data"] = billClaimLogs
	} else {
		res := make([]models.BillClaimLog, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
