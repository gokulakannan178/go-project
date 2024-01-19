package handlers

import (
	"encoding/json"
	"fmt"
	"municipalproduct1-service/app"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//SaveProperty : ""
func (h *Handler) SaveEstimatedPropertyDemand(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	property := new(models.Property)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var collectionName string
	collectionName = constants.COLLECTIONESTIMATEDPROPERTYDEMAND
	err = h.Service.SaveProperty(ctx, property, collectionName)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	// err = h.Service.SavePropertyDemand(ctx, property.UniqueID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	// err = h.Service.SaveOverAllPropertyDemandToProperty(ctx, property.UniqueID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }

	m := make(map[string]interface{})
	m["property"] = property
	response.With200V2(w, "Success", m, platform)
}

//GetPropertyDemandCalc : ""
func (h *Handler) GetEstimatedPropertyDemandCalc(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID
	var collectionName string
	collectionName = constants.COLLECTIONESTIMATEDPROPERTYDEMAND
	propertyDemand, err := h.Service.GetPropertyDemandCalc(ctx, filter, collectionName)
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
	m["propertyDemand"] = propertyDemand
	response.With200V2(w, "Success", m, platform)
}
