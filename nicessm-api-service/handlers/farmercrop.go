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

//SaveFarmerCrop : ""
func (h *Handler) SaveFarmerCrop(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	farmerCrop := new(models.FarmerCrop)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&farmerCrop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveFarmerCrop(ctx, farmerCrop)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCrop"] = farmerCrop
	response.With200V2(w, "Success", m, platform)
}

//UpdateFarmerCrop :""
func (h *Handler) UpdateFarmerCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	farmerCrop := new(models.FarmerCrop)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&farmerCrop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if farmerCrop.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateFarmerCrop(ctx, farmerCrop)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["FarmerCrop"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableFarmerCrop : ""
func (h *Handler) EnableFarmerCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableFarmerCrop(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCrop"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableFarmerCrop : ""
func (h *Handler) DisableFarmerCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableFarmerCrop(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCrop"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteFarmerCrop : ""
func (h *Handler) DeleteFarmerCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteFarmerCrop(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCrop"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleFarmerCrop :""
func (h *Handler) GetSingleFarmerCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	farmerCrop := new(models.RefFarmerCrop)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	farmerCrop, err := h.Service.GetSingleFarmerCrop(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerCrop"] = farmerCrop
	response.With200V2(w, "Success", m, platform)
}

//FilterFarmerCrop : ""
func (h *Handler) FilterFarmerCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var farmerCrop *models.FarmerCropFilter
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
	err := json.NewDecoder(r.Body).Decode(&farmerCrop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var farmerCrops []models.RefFarmerCrop
	log.Println(pagination)
	farmerCrops, err = h.Service.FilterFarmerCrop(ctx, farmerCrop, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(farmerCrops) > 0 {
		m["farmerCrop"] = farmerCrops
	} else {
		res := make([]models.FarmerCrop, 0)
		m["farmerCrop"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//UpdateFarmerCropDone :""
func (h *Handler) UpdateFarmerCropDone(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	farmerCrop := new(models.FarmerCrop)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&farmerCrop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if farmerCrop.ID.String() == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateFarmerCropDone(ctx, farmerCrop)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["FarmerCrop"] = "success"
	response.With200V2(w, "Success", m, platform)
}
