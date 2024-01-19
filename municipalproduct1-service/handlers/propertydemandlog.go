package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

// SaveOverAllPropertyDemand : ""
func (h *Handler) SaveOverAllPropertyDemand(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	PropertyID := r.URL.Query().Get("id")

	if PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	// pmTarget := new(models.RefPmTarget)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.SaveOverAllPropertyDemandToProperty(ctx, PropertyID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["overallPropertyDemand"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetOverAllPropertyDemand : ""
func (h *Handler) GetOverAllPropertyDemand(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	PropertyID := r.URL.Query().Get("id")

	if PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	// pmTarget := new(models.RefPmTarget)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	demand, err := h.Service.GetOverAllPropertyDemandToProperty(ctx, PropertyID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["overallPropertyDemand"] = demand
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyDemandLogPropertyID : ""
func (h *Handler) UpdatePropertyDemandLogPropertyID(w http.ResponseWriter, r *http.Request) {

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
	err = h.Service.UpdatePropertyDemandLogPropertyID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["updateMobileTowerPropertyID"] = "success"
	response.With200V2(w, "Success", m, platform)
}
