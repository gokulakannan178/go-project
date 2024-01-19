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

// SavePropertyMobileTower : ""
func (h *Handler) SavePropertyMobileTower(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	mobile := new(models.PropertyMobileTower)
	err := json.NewDecoder(r.Body).Decode(&mobile)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyMobileTower(ctx, mobile)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyMobileTower"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSinglePropertyMobileTower : ""
func (h *Handler) GetSinglePropertyMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	mobile := new(models.RefPropertyMobileTower)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	mobile, err := h.Service.GetSinglePropertyMobileTower(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyMobileTower"] = mobile
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyMobileTower : ""
func (h *Handler) UpdatePropertyMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	mobile := new(models.PropertyMobileTower)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&mobile)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if mobile.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyMobileTower(ctx, mobile)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyMobileTower"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePropertyMobileTower : ""
func (h *Handler) EnablePropertyMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnablePropertyMobileTower(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyMobileTower"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisablePropertyMobileTower : ""
func (h *Handler) DisablePropertyMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisablePropertyMobileTower(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyMobileTower"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeletePropertyMobileTower : ""
func (h *Handler) DeletePropertyMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeletePropertyMobileTower(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyMobileTower"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// RejectPropertyMobileTower : ""
func (h *Handler) RejectPropertyMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.RejectPropertyMobileTower(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyMobileTower"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterPropertyMobileTower : ""
func (h *Handler) FilterPropertyMobileTower(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.PropertyMobileTowerFilter
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

	var users []models.RefPropertyMobileTower
	log.Println(pagination)
	users, err = h.Service.FilterPropertyMobileTower(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(users) > 0 {
		m["propertyMobileTower"] = users
	} else {
		res := make([]models.MobileTowerTax, 0)
		m["propertyMobileTower"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// MobileTowerWithMobileNo : ""
func (h *Handler) MobileTowerWithMobileNo(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	// mobileNo := r.URL.Query().Get("mobileNo")
	var filter *models.MobileTowerWithMobileNoReq
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
	if filter.MobileNo == "" {
		response.With400V2(w, "mobileNo missing", platform)
		return
	}

	var report []models.MobileTowerWithMobileNoRes
	log.Println(pagination)
	report, err = h.Service.MobileTowerWithMobileNo(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(report) > 0 {
		m["mobileTowers"] = report
	} else {
		res := make([]models.MobileTowerWithMobileNoRes, 0)
		m["mobileTowers"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// MobileTowerPenaltyUpdate : ""
func (h *Handler) MobileTowerPenaltyUpdate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	// uniqueID := r.URL.Query().Get("id")

	mobile := new(models.MobileTowerPenaltyUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&mobile)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.MobileTowerPenaltyUpdate(ctx, mobile)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobileTowerPenalty"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// UpdateMobileTowerPropertyID : ""
func (h *Handler) UpdateMobileTowerPropertyID(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.UpdatePropertyUniqueID)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if len(property.UniqueIDs) == 0 {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateMobileTowerPropertyID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["updateMobileTowerPropertyID"] = "success"
	response.With200V2(w, "Success", m, platform)
}
