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

//SaveAssetType : ""
func (h *Handler) SaveAssetType(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	assetType := new(models.AssetType)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&assetType)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveAssetType(ctx, assetType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetType"] = assetType
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) SaveAssetTypeWithPropertys(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	assetType := new(models.AssetType)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&assetType)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveAssetTypeWithPropertys(ctx, assetType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetType"] = assetType
	response.With200V2(w, "Success", m, platform)
}

//GetSingleAssetType :""
func (h *Handler) GetSingleAssetType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	assetType := new(models.RefAssetType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	assetType, err := h.Service.GetSingleAssetType(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetType"] = assetType
	response.With200V2(w, "Success", m, platform)
}

//UpdateAssetType :""
func (h *Handler) UpdateAssetType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	assetType := new(models.AssetType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&assetType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if assetType.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAssetType(ctx, assetType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableAssetType : ""
func (h *Handler) EnableAssetType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableAssetType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableAssetType : ""
func (h *Handler) DisableAssetType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableAssetType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteAssetType : ""
func (h *Handler) DeleteAssetType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteAssetType(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetType"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterAssetType : ""
func (h *Handler) FilterAssetType(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var assetType *models.FilterAssetType
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
	err := json.NewDecoder(r.Body).Decode(&assetType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var assetTypes []models.RefAssetType
	log.Println(pagination)
	assetTypes, err = h.Service.FilterAssetType(ctx, assetType, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(assetTypes) > 0 {
		m["AssetType"] = assetTypes
	} else {
		res := make([]models.AssetType, 0)
		m["AssetType"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// func (h *Handler) GetAssetTypePropertys(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")
// 	UniqueID := r.URL.Query().Get("id")

// 	if UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 	}

// 	assetType := new(models.RefAssetType)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())

// 	assetType, err := h.Service.GetAssetTypePropertys(ctx, UniqueID)
// 	if err != nil {
// 		if err.Error() == "mongo: no documents in result" {
// 			response.With500mV2(w, "failed no data for this id", platform)
// 			return
// 		}
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["AssetType"] = assetType
// 	response.With200V2(w, "Success", m, platform)
// }
func (h *Handler) UpdateAssetTypeWithProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	assetType := new(models.UpdateAssetType)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&assetType)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if assetType.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateAssetTypeWithProperty(ctx, assetType)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetType"] = "success"
	response.With200V2(w, "Success", m, platform)
}
