package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
	"strconv"
)

//SaveAssetTypePropertys : ""
func (h *Handler) SaveAssetTypePropertys(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	assetTypePropertys := new(models.AssetTypePropertys)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&assetTypePropertys)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveAssetTypePropertys(ctx, assetTypePropertys)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetTypePropertys"] = assetTypePropertys
	response.With200V2(w, "Success", m, platform)
}

//GetSingleAssetTypePropertys :""
func (h *Handler) GetSingleAssetTypePropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	assetTypePropertys := new(models.RefAssetTypePropertys)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	assetTypePropertys, err := h.Service.GetSingleAssetTypePropertys(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetTypePropertys"] = assetTypePropertys
	response.With200V2(w, "Success", m, platform)
}

//UpdateAssetTypePropertys :""
func (h *Handler) UpdateAssetTypePropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	assetTypePropertys := new(models.AssetTypePropertys)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&assetTypePropertys)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if assetTypePropertys.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAssetTypePropertys(ctx, assetTypePropertys)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetTypePropertys"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableAssetTypePropertys : ""
func (h *Handler) EnableAssetTypePropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableAssetTypePropertys(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetTypePropertys"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableAssetTypePropertys : ""
func (h *Handler) DisableAssetTypePropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableAssetTypePropertys(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetTypePropertys"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteAssetTypePropertys : ""
func (h *Handler) DeleteAssetTypePropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteAssetTypePropertys(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetTypePropertys"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterAssetTypePropertys : ""
func (h *Handler) FilterAssetTypePropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var assetTypePropertys *models.FilterAssetTypePropertys
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
	err := json.NewDecoder(r.Body).Decode(&assetTypePropertys)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var assetTypePropertysFilter []models.RefAssetTypePropertys
	log.Println(pagination)
	assetTypePropertysFilter, err = h.Service.FilterAssetTypePropertys(ctx, assetTypePropertys, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(assetTypePropertysFilter) > 0 {
		m["AssetTypePropertys"] = assetTypePropertysFilter
	} else {
		res := make([]models.AssetTypePropertys, 0)
		m["AssetTypePropertys"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
