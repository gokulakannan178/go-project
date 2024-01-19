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

//SaveLandCrop : ""
func (h *Handler) SaveLandCrop(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	landCrop := new(models.LandCrop)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&landCrop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveLandCrop(ctx, landCrop)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCrop"] = landCrop
	response.With200V2(w, "Success", m, platform)
}

//UpdateLandCrop :""
func (h *Handler) UpdateLandCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	landCrop := new(models.LandCrop)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&landCrop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if landCrop.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateLandCrop(ctx, landCrop)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCrop"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableLandCrop : ""
func (h *Handler) EnableLandCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableLandCrop(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCrop"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableLandCrop : ""
func (h *Handler) DisableLandCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableLandCrop(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCrop"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteLandCrop : ""
func (h *Handler) DeleteLandCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteLandCrop(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCrop"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleLandCrop :""
func (h *Handler) GetSingleLandCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	landCrop := new(models.RefLandCrop)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	landCrop, err := h.Service.GetSingleLandCrop(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["landCrop"] = landCrop
	response.With200V2(w, "Success", m, platform)
}

//FilterLandCrop : ""
func (h *Handler) FilterLandCrop(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var landCrop *models.LandCropFilter
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
	err := json.NewDecoder(r.Body).Decode(&landCrop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var landCrops []models.RefLandCrop
	log.Println(pagination)
	landCrops, err = h.Service.FilterLandCrop(ctx, landCrop, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(landCrops) > 0 {
		m["landCrop"] = landCrops
	} else {
		res := make([]models.LandCrop, 0)
		m["landCrop"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
