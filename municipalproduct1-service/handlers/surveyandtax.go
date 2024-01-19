package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

// SaveSurveyAndTax : ""
func (h *Handler) SaveSurveyAndTax(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	sat := new(models.SurveyAndTax)
	err := json.NewDecoder(r.Body).Decode(&sat)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveSurveyAndTax(ctx, sat)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankDeposit"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleSurveyAndTax : ""
func (h *Handler) GetSingleSurveyAndTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	property := new(models.RefSurveyAndTax)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	property, err := h.Service.GetSingleSurveyAndTax(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = property
	response.With200V2(w, "Success", m, platform)
}

// PushNotification : ""
func (h *Handler) PushNotification(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	property := new(models.SurveyAndTax)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	property, err := h.Service.PushNotification(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = property
	response.With200V2(w, "Success", m, platform)
}

// SurveyAndTaxFilter : ""
func (h *Handler) SurveyAndTaxFilter(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	//uniqueID := r.URL.Query().Get("id")
	//var user *models.UserFilter
	satf := new(models.SurveyAndTaxFilter)
	err := json.NewDecoder(r.Body).Decode(&satf)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	var pagination *models.Pagination
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
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

	data, err := h.Service.SurveyAndTaxFilter(ctx, satf, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankdeposit"] = data
	m["pagination"] = pagination
	response.With200V2(w, "Success", m, platform)
}
