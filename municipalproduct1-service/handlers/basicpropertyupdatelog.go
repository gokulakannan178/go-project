package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//GetSingleProperty :""
func (h *Handler) GetSingleBasicPropertyUpdateLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	property := new(models.RefBasicPropertyUpdateLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	property, err := h.Service.GetSingleBasicPropertyUpdateLog(ctx, UniqueID)
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

//CreatedBankDeposit : ""
func (h *Handler) BasicPropertyUpdateGetPaymentsToBeUpdated(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	rbpul := new(models.RefBasicPropertyUpdateLog)
	err := json.NewDecoder(r.Body).Decode(&rbpul)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	payments, err := h.Service.BasicPropertyUpdateGetPaymentsToBeUpdated(ctx, rbpul)
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

// UpdateBasicPropertyUpdateLogPropertyID : ""
func (h *Handler) UpdateBasicPropertyUpdateLogPropertyID(w http.ResponseWriter, r *http.Request) {

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
	err = h.Service.UpdateBasicPropertyUpdateLogPropertyID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["updateBasicPropertyUpdateLogPropertyID"] = "success"
	response.With200V2(w, "Success", m, platform)
}
