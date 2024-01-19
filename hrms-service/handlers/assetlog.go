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

// SaveAssetLog : ""
func (h *Handler) SaveAssetLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Elog := new(models.AssetLog)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Elog)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()
	err = h.Service.SaveAssetLog(ctx, Elog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetLog"] = Elog
	response.With200V2(w, "Success", m, platform)
}

// GetSingleAssetLog : ""
func (h *Handler) GetSingleAssetLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	task := new(models.RefAssetLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	task, err := h.Service.GetSingleAssetLog(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetLog"] = task
	response.With200V2(w, "Success", m, platform)
}

//UpdateAssetLog : ""
func (h *Handler) UpdateAssetLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	assetlog := new(models.AssetLog)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&assetlog)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if assetlog.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateAssetLog(ctx, assetlog)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// EnableAssetLog : ""
func (h *Handler) EnableAssetLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.EnableAssetLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DisableAssetLog : ""
func (h *Handler) DisableAssetLog(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
		return
	}
	err := h.Service.DisableAssetLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetLog"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteAssetLog : ""
func (h *Handler) DeleteAssetLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "ID is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteAssetLog(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["AssetLog"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterAssetLog : ""
func (h *Handler) FilterAssetLog(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var ft *models.FilterAssetLog
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
	err := json.NewDecoder(r.Body).Decode(&ft)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var fts []models.RefAssetLog
	log.Println(pagination)
	fts, err = h.Service.FilterAssetLog(ctx, ft, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(fts) > 0 {
		m["AssetLog"] = fts
	} else {
		res := make([]models.AssetLog, 0)
		m["AssetLog"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
