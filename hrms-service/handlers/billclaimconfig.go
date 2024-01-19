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

//SaveBillclaimConfig : ""
func (h *Handler) SaveBillclaimConfig(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	billclaimconfig := new(models.BillclaimConfig)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&billclaimconfig)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveBillclaimConfig(ctx, billclaimconfig)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimConfig"] = billclaimconfig
	response.With200V2(w, "Success", m, platform)
}

//GetSingleBillclaimConfig :""
func (h *Handler) GetSingleBillclaimConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	billclaimconfig := new(models.RefBillclaimConfig)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	billclaimconfig, err := h.Service.GetSingleBillclaimConfig(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimConfig"] = billclaimconfig
	response.With200V2(w, "Success", m, platform)
}

//UpdateBillclaimConfig :""
func (h *Handler) UpdateBillclaimConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	billclaimconfig := new(models.BillclaimConfig)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&billclaimconfig)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if billclaimconfig.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateBillclaimConfig(ctx, billclaimconfig)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableBillclaimConfig : ""
func (h *Handler) EnableBillclaimConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableBillclaimConfig(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableBillclaimConfig : ""
func (h *Handler) DisableBillclaimConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableBillclaimConfig(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteBillclaimConfig : ""
func (h *Handler) DeleteBillclaimConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteBillclaimConfig(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["BillclaimConfig"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterBillclaimConfig : ""
func (h *Handler) FilterBillclaimConfig(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var billclaimconfig *models.BillclaimConfigFilter
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
	err := json.NewDecoder(r.Body).Decode(&billclaimconfig)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var billclaimconfigs []models.RefBillclaimConfig
	log.Println(pagination)
	billclaimconfigs, err = h.Service.FilterBillclaimConfig(ctx, billclaimconfig, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(billclaimconfigs) > 0 {
		m["BillclaimConfig"] = billclaimconfigs
	} else {
		res := make([]models.BillclaimConfig, 0)
		m["BillclaimConfig"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
