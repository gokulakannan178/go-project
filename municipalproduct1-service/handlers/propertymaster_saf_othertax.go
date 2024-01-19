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

//SavePropertyOtherTax : ""
func (h *Handler) SavePropertyOtherTax(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	propertyOtherTax := new(models.PropertyOtherTax)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&propertyOtherTax)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePropertyOtherTax(ctx, propertyOtherTax)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyOtherTax"] = propertyOtherTax
	response.With200V2(w, "Success", m, platform)
}

//UpdatePropertyOtherTax :""
func (h *Handler) UpdatePropertyOtherTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	propertyOtherTax := new(models.PropertyOtherTax)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&propertyOtherTax)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if propertyOtherTax.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyOtherTax(ctx, propertyOtherTax)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyOtherTax"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePropertyOtherTax : ""
func (h *Handler) EnablePropertyOtherTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnablePropertyOtherTax(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyOtherTax"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisablePropertyOtherTax : ""
func (h *Handler) DisablePropertyOtherTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisablePropertyOtherTax(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyOtherTax"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePropertyOtherTax : ""
func (h *Handler) DeletePropertyOtherTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeletePropertyOtherTax(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyOtherTax"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePropertyOtherTax :""
func (h *Handler) GetSinglePropertyOtherTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	propertyOtherTax := new(models.RefPropertyOtherTax)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	propertyOtherTax, err := h.Service.GetSinglePropertyOtherTax(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyOtherTax"] = propertyOtherTax
	response.With200V2(w, "Success", m, platform)
}

//FilterPropertyOtherTax : ""
func (h *Handler) FilterPropertyOtherTax(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var propertyOtherTax *models.PropertyOtherTaxFilter
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
	err := json.NewDecoder(r.Body).Decode(&propertyOtherTax)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var propertyOtherTaxs []models.RefPropertyOtherTax
	log.Println(pagination)
	propertyOtherTaxs, err = h.Service.FilterPropertyOtherTax(ctx, propertyOtherTax, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(propertyOtherTaxs) > 0 {
		m["propertyOtherTax"] = propertyOtherTaxs
	} else {
		res := make([]models.PropertyOtherTax, 0)
		m["propertyOtherTax"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
