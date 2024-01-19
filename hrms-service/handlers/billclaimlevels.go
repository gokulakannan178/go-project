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

//SaveBillclaimLevels : ""
func (h *Handler) SaveBillclaimLevels(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	billclaimlevels := new(models.BillclaimLevels)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&billclaimlevels)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveBillclaimLevels(ctx, billclaimlevels)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimLevels"] = billclaimlevels
	response.With200V2(w, "Success", m, platform)
}

//GetSingleBillclaimLevels :""
func (h *Handler) GetSingleBillclaimLevels(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	billclaimlevels := new(models.RefBillclaimLevels)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	billclaimlevels, err := h.Service.GetSingleBillclaimLevels(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimLevels"] = billclaimlevels
	response.With200V2(w, "Success", m, platform)
}

//UpdateBillclaimLevels :""
func (h *Handler) UpdateBillclaimLevels(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	billclaimlevels := new(models.BillclaimLevels)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&billclaimlevels)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if billclaimlevels.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateBillclaimLevels(ctx, billclaimlevels)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimLevels"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableBillclaimLevels : ""
func (h *Handler) EnableBillclaimLevels(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableBillclaimLevels(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimLevels"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableBillclaimLevels : ""
func (h *Handler) DisableBillclaimLevels(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableBillclaimLevels(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimLevels"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteBillclaimLevels : ""
func (h *Handler) DeleteBillclaimLevels(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteBillclaimLevels(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimLevels"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterBillclaimLevels : ""
func (h *Handler) FilterBillclaimLevels(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var billclaimlevels *models.BillclaimLevelsFilter
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
	err := json.NewDecoder(r.Body).Decode(&billclaimlevels)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var billclaimlevelss []models.RefBillclaimLevels
	log.Println(pagination)
	billclaimlevelss, err = h.Service.FilterBillclaimLevels(ctx, billclaimlevels, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(billclaimlevelss) > 0 {
		m["BillclaimLevels"] = billclaimlevelss
	} else {
		res := make([]models.BillclaimLevels, 0)
		m["BillclaimLevels"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) ApprovedBillclaimLevels(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	billclaimlevels := new(models.BillclaimLevels)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&billclaimlevels)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if billclaimlevels.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.ApprovedBillclaimLevels(ctx, billclaimlevels)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimLevels"] = "success"
	response.With200V2(w, "Success", m, platform)
}
