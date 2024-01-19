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

// BasicMobileTowerUpdate : ""
func (h *Handler) BasicMobileTowerUpdate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	mtu := new(models.BasicMobileTowerUpdateData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&mtu)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	fmt.Println(mtu)
	err = h.Service.BasicMobileTowerUpdate(ctx, mtu)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basicmobiletowerupdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// AcceptBasicMobileTowerUpdate : ""
func (h *Handler) AcceptBasicMobileTowerUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.AcceptBasicMobileTowerUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.AcceptBasicMobileTowerUpdate(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basicmobiletowerupdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// RejectBasicMobileTowerUpdate : ""
func (h *Handler) RejectBasicMobileTowerUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.RejectBasicMobileTowerUpdate)

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.RejectBasicMobileTowerUpdate(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BasicMobileTowerUpdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterBasicMobileTowerUpdateLog : ""
func (h *Handler) FilterBasicMobileTowerUpdateLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.FilterBasicMobileTowerUpdateLog
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

	var refs []models.RefBasicMobileTowerUpdateLog
	log.Println(pagination)
	refs, err = h.Service.FilterBasicMobileTowerUpdateLog(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(refs) > 0 {
		m["data"] = refs
	} else {
		res := make([]models.BasicMobileTowerUpdateLog, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// GetSingleBasicMobileTowerUpdateLogV2 : ""
func (h *Handler) GetSingleBasicMobileTowerUpdateLogV2(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	mobile, err := h.Service.GetSingleBasicMobileTowerUpdateLogV2(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basicmobiletowerupdate"] = mobile
	response.With200V2(w, "Success", m, platform)
}

// BasicMobileTowerUpdateGetPaymentsToBeUpdated : ""
func (h *Handler) BasicMobileTowerUpdateGetPaymentsToBeUpdated(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	rbmtul := new(models.RefBasicMobileTowerUpdateLogV2)
	err := json.NewDecoder(r.Body).Decode(&rbmtul)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	payments, err := h.Service.BasicMobileTowerUpdateGetPaymentsToBeUpdated(ctx, rbmtul)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payments"] = payments
	if len(payments) < 1 {
		m["payments"] = []interface{}{}

	}
	response.With200V2(w, "Success", m, platform)
}
