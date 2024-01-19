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

//SaveProperty : ""
func (h *Handler) SaveLegacy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	legacy := new(models.RegLegacyProperty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&legacy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if len(legacy.LegacyPropertyFy) == 0 {
		response.With400V2(w, "need atleast one financial year", platform)
		return
	}
	if legacy.LegacyProperty.PropertyID == "" {
		response.With400V2(w, "property id cannot be empty", platform)
		return
	}
	err = h.Service.SaveLegacy(ctx, legacy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["legacy"] = legacy
	response.With200V2(w, "Success", m, platform)
}

//SaveProperty : ""
func (h *Handler) GetLegacyForAProperty(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	if ID == "" {
		response.With400V2(w, "ID missing", platform)
		return
	}

	legacy, err := h.Service.GetLegacyForAProperty(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["legacy"] = legacy
	response.With200V2(w, "Success", m, platform)
}

//UpdateLegacyForAProperty : ""
func (h *Handler) UpdateLegacyForAProperty(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	legacy := new(models.RegLegacyProperty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&legacy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if len(legacy.LegacyPropertyFy) == 0 {
		response.With400V2(w, "need atleast one financial year", platform)
		return
	}
	if legacy.LegacyProperty.PropertyID == "" {
		response.With400V2(w, "property id cannot be empty", platform)
		return
	}
	err = h.Service.UpdateLegacyForAProperty(ctx, legacy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	// ctx = app.GetApp(r.Context(), h.Service.Daos)
	// defer ctx.Client.Disconnect(r.Context())
	// //err = h.Service.SavePropertyDemand(ctx, property.UniqueID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }

	m := make(map[string]interface{})
	m["legacy"] = legacy
	response.With200V2(w, "Success", m, platform)
}

//GetFinancialYearsForLegacyPayments : ""
func (h *Handler) GetFinancialYearsForLegacyPayments(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("propertyId")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	if ID == "" {
		response.With400V2(w, "ID missing", platform)
		return
	}

	data, err := h.Service.GetFinancialYearsForLegacyPayments(ctx, ID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["legacyyears"] = data
	response.With200V2(w, "Success", m, platform)
}

//GetReqFinancialYearForLegacy : ""
func (h *Handler) GetReqFinancialYearForLegacy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	grfy := new(models.GetReqFinancialYear)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&grfy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if grfy.Doa == nil {
		response.With400V2(w, "need atleast one financial year", platform)
		return
	}
	data, err := h.Service.GetReqFinancialYearForLegacy(ctx, grfy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["legacyyears"] = data
	response.With200V2(w, "Success", m, platform)
}

// GetLegacyForAPropertyV2 : ""
func (h *Handler) GetLegacyForAPropertyV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	PropertyID := r.URL.Query().Get("id")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	if PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	legacy, err := h.Service.GetLegacyForAProperty(ctx, PropertyID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = PropertyID
	propertyDemand, err := h.Service.GetPropertyDemandCalc(ctx, filter, "")
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	propertyDemand.PropertyID = filter.PropertyID
	propertyDemand.OverallPropertyDemand.PropertyID = filter.PropertyID
	//demand.OverallPropertyDemand
	if err := h.Service.UpdateOverallPropertyDemand(ctx, &propertyDemand.OverallPropertyDemand); err != nil {
		fmt.Println("ERR IN UPDATING OVERALL PROPERTY DEMAND - " + err.Error())
	}
	m := make(map[string]interface{})
	m["legacy"] = legacy
	m["propertyDemand"] = propertyDemand
	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyPaymentModeChange : ""
func (h *Handler) FilterLegacy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.LegacyPropertyFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	resType := r.URL.Query().Get("resType")
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

	var requests []models.RefLegacyPropertyPayment
	log.Println(pagination)
	if resType == "excel" {
		file, err := h.Service.LegacyReportExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	requests, err = h.Service.FilterLegacy(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(requests) > 0 {
		m["legacy"] = requests
	} else {
		res := make([]models.RegLegacyProperty, 0)
		m["legacy"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
