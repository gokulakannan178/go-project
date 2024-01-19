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

//SaveAssetPolicy : ""
func (h *Handler) SaveAssetPolicy(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	assetPolicy := new(models.AssetPolicy)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&assetPolicy)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveAssetPolicy(ctx, assetPolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicy"] = assetPolicy
	response.With200V2(w, "Success", m, platform)
}

//GetSingleAssetPolicy :""
func (h *Handler) GetSingleAssetPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	assetPolicy := new(models.RefAssetPolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	assetPolicy, err := h.Service.GetSingleAssetPolicy(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicy"] = assetPolicy
	response.With200V2(w, "Success", m, platform)
}

//UpdateAssetPolicy :""
func (h *Handler) UpdateAssetPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	assetPolicy := new(models.AssetPolicy)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&assetPolicy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if assetPolicy.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAssetPolicy(ctx, assetPolicy)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableAssetPolicy : ""
func (h *Handler) EnableAssetPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableAssetPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableAssetPolicy : ""
func (h *Handler) DisableAssetPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableAssetPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteAssetPolicy : ""
func (h *Handler) DeleteAssetPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteAssetPolicy(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetPolicy"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterAssetPolicy : ""
func (h *Handler) FilterAssetPolicy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var assetPolicy *models.FilterAssetPolicy
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
	err := json.NewDecoder(r.Body).Decode(&assetPolicy)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var assetPolicys []models.RefAssetPolicy
	log.Println(pagination)
	assetPolicys, err = h.Service.FilterAssetPolicy(ctx, assetPolicy, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(assetPolicys) > 0 {
		m["AssetPolicy"] = assetPolicys
	} else {
		res := make([]models.AssetPolicy, 0)
		m["AssetPolicy"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
