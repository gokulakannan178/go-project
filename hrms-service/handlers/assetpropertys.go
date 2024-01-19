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

//SaveAssetPropertys : ""
func (h *Handler) SaveAssetPropertys(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	assetPropertys := new(models.AssetPropertys)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&assetPropertys)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveAssetPropertys(ctx, assetPropertys)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPropertys"] = assetPropertys
	response.With200V2(w, "Success", m, platform)
}

//GetSingleAssetPropertys :""
func (h *Handler) GetSingleAssetPropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	assetPropertys := new(models.RefAssetPropertys)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	assetPropertys, err := h.Service.GetSingleAssetPropertys(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPropertys"] = assetPropertys
	response.With200V2(w, "Success", m, platform)
}

//UpdateAssetPropertys :""
func (h *Handler) UpdateAssetPropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	assetPropertys := new(models.AssetPropertys)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&assetPropertys)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if assetPropertys.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAssetPropertys(ctx, assetPropertys)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPropertys"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableAssetPropertys : ""
func (h *Handler) EnableAssetPropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableAssetPropertys(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPropertys"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableAssetPropertys : ""
func (h *Handler) DisableAssetPropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableAssetPropertys(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPropertys"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteAssetPropertys : ""
func (h *Handler) DeleteAssetPropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteAssetPropertys(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPropertys"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterAssetPropertys : ""
func (h *Handler) FilterAssetPropertys(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var assetPropertys *models.FilterAssetPropertys
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
	err := json.NewDecoder(r.Body).Decode(&assetPropertys)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var assetPropertysFilter []models.RefAssetPropertys
	log.Println(pagination)
	assetPropertysFilter, err = h.Service.FilterAssetPropertys(ctx, assetPropertys, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(assetPropertysFilter) > 0 {
		m["AssetPropertys"] = assetPropertysFilter
	} else {
		res := make([]models.AssetPropertys, 0)
		m["AssetPropertys"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
