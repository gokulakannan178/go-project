package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
	"strconv"
)

//SaveFarmerSoilData : ""
func (h *Handler) SaveFarmerSoilData(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	farmerSoilData := new(models.FarmerSoilData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&farmerSoilData)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveFarmerSoilData(ctx, farmerSoilData)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerSoilData"] = farmerSoilData
	response.With200V2(w, "Success", m, platform)
}

//UpdateFarmerSoilData :""
func (h *Handler) UpdateFarmerSoilData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	farmerSoilData := new(models.FarmerSoilData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&farmerSoilData)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if farmerSoilData.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateFarmerSoilData(ctx, farmerSoilData)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["FarmerSoilData"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableFarmerSoilData : ""
func (h *Handler) EnableFarmerSoilData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableFarmerSoilData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerSoilData"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableFarmerSoilData : ""
func (h *Handler) DisableFarmerSoilData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableFarmerSoilData(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerSoilData"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteFarmerSoilData : ""
func (h *Handler) DeleteFarmerSoilData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteFarmerSoilData(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerSoilData"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleFarmerSoilData :""
func (h *Handler) GetSingleFarmerSoilData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	farmerSoilData := new(models.RefFarmerSoilData)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	farmerSoilData, err := h.Service.GetSingleFarmerSoilData(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerSoilData"] = farmerSoilData
	response.With200V2(w, "Success", m, platform)
}

//FilterFarmerSoilData : ""
func (h *Handler) FilterFarmerSoilData(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var farmerSoilData *models.FarmerSoilDataFilter
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
	err := json.NewDecoder(r.Body).Decode(&farmerSoilData)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var farmerSoilDatas []models.RefFarmerSoilData
	log.Println(pagination)
	farmerSoilDatas, err = h.Service.FilterFarmerSoilData(ctx, farmerSoilData, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(farmerSoilDatas) > 0 {
		m["farmerSoilData"] = farmerSoilDatas
	} else {
		res := make([]models.FarmerSoilData, 0)
		m["farmerSoilData"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
