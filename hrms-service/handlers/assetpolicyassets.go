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

//SaveAssetPolicyAssets : ""
func (h *Handler) SaveAssetPolicyAssets(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	assetPolicyAssets := new(models.AssetPolicyAssets)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&assetPolicyAssets)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveAssetPolicyAssets(ctx, assetPolicyAssets)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicyAssets"] = assetPolicyAssets
	response.With200V2(w, "Success", m, platform)
}

//GetSingleAssetPolicyAssets :""
func (h *Handler) GetSingleAssetPolicyAssets(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	assetPolicyAssets := new(models.RefAssetPolicyAssets)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	assetPolicyAssets, err := h.Service.GetSingleAssetPolicyAssets(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicyAssets"] = assetPolicyAssets
	response.With200V2(w, "Success", m, platform)
}

//UpdateAssetPolicyAssets :""
func (h *Handler) UpdateAssetPolicyAssets(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	assetPolicyAssets := new(models.AssetPolicyAssets)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&assetPolicyAssets)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if assetPolicyAssets.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAssetPolicyAssets(ctx, assetPolicyAssets)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicyAssets"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableAssetPolicyAssets : ""
func (h *Handler) EnableAssetPolicyAssets(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableAssetPolicyAssets(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicyAssets"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableAssetPolicyAssets : ""
func (h *Handler) DisableAssetPolicyAssets(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableAssetPolicyAssets(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicyAssets"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteAssetPolicyAssets : ""
func (h *Handler) DeleteAssetPolicyAssets(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteAssetPolicyAssets(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicyAssets"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterAssetPolicyAssets : ""
func (h *Handler) FilterAssetPolicyAssets(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var assetPolicyAssets *models.FilterAssetPolicyAssets
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
	err := json.NewDecoder(r.Body).Decode(&assetPolicyAssets)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var assetPolicyAssetsFilter []models.RefAssetPolicyAssets
	log.Println(pagination)
	assetPolicyAssetsFilter, err = h.Service.FilterAssetPolicyAssets(ctx, assetPolicyAssets, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(assetPolicyAssetsFilter) > 0 {
		m["AssetPolicyAssets"] = assetPolicyAssetsFilter
	} else {
		res := make([]models.AssetPolicyAssets, 0)
		m["AssetPolicyAssets"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
