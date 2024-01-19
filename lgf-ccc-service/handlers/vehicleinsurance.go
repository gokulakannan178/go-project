package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"log"
	"net/http"
	"strconv"
)

//SaveVehicleInsurance : ""
func (h *Handler) SaveVehicleInsurance(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	vehicleinsurance := new(models.VehicleInsurance)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&vehicleinsurance)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveVehicleInsurance(ctx, vehicleinsurance)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleinsurance"] = vehicleinsurance
	response.With200V2(w, "Success", m, platform)
}

//GetSingleVehicleInsurance :""
func (h *Handler) GetSingleVehicleInsurance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var vehicleinsurance *models.RefVehicleInsurance
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	vehicleinsurance, err := h.Service.GetSingleVehicleInsurance(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleinsurance"] = vehicleinsurance
	response.With200V2(w, "Success", m, platform)
}

//UpdateVehicleInsurance :""
func (h *Handler) UpdateVehicleInsurance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var vehicleinsurance *models.VehicleInsurance
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&vehicleinsurance)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if vehicleinsurance.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateVehicleInsurance(ctx, vehicleinsurance)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleinsurance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableVehicleInsurance : ""
func (h *Handler) EnableVehicleInsurance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableVehicleInsurance(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleinsurance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableVehicleInsurance : ""
func (h *Handler) DisableVehicleInsurance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableVehicleInsurance(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleinsurance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteVehicleInsurance : ""
func (h *Handler) DeleteVehicleInsurance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteVehicleInsurance(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["vehicleinsurance"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterVehicleInsurance : ""
func (h *Handler) FilterVehicleInsurance(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.VehicleInsuranceFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
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

	var vehicleinsurances []models.VehicleInsurance
	log.Println(pagination)
	vehicleinsurances, err = h.Service.VehicleInsuranceFilter(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(vehicleinsurances) > 0 {
		m["vehicleinsurance"] = vehicleinsurances
	} else {
		res := make([]models.User, 0)
		m["vehicleinsurance"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
