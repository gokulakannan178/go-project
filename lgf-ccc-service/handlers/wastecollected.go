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

//SaveWasteCollected : ""
func (h *Handler) SaveWasteCollected(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	WasteCollected := new(models.WasteCollected)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&WasteCollected)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.SaveWasteCollected(ctx, WasteCollected)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WasteCollected"] = WasteCollected
	response.With200V2(w, "Success", m, platform)
}

//GetSingleWasteCollected :""
func (h *Handler) GetSingleWasteCollected(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	WasteCollected := new(models.RefWasteCollected)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	WasteCollected, err := h.Service.GetSingleWasteCollected(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WasteCollected"] = WasteCollected
	response.With200V2(w, "Success", m, platform)
}

//UpdateWasteCollected :""
func (h *Handler) UpdateWasteCollected(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	WasteCollected := new(models.WasteCollected)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&WasteCollected)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if WasteCollected.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateWasteCollected(ctx, WasteCollected)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WasteCollected"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableWasteCollected : ""
func (h *Handler) EnableWasteCollected(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableWasteCollected(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WasteCollected"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableWasteCollected : ""
func (h *Handler) DisableWasteCollected(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableWasteCollected(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WasteCollected"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) WasteCollectedCompleted(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.WasteCollectedCompleted(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WasteCollected"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableWasteCollected : ""
func (h *Handler) WasteCollectedPending(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.WasteCollectedPending(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WasteCollected"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteWasteCollected : ""
func (h *Handler) DeleteWasteCollected(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteWasteCollected(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["WasteCollected"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//FilterWasteCollected : ""
func (h *Handler) FilterWasteCollected(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var WasteCollected *models.FilterWasteCollected
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
	err := json.NewDecoder(r.Body).Decode(&WasteCollected)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var WasteCollecteds []models.RefWasteCollected
	log.Println(pagination)
	WasteCollecteds, err = h.Service.FilterWasteCollected(ctx, WasteCollected, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(WasteCollecteds) > 0 {
		m["WasteCollected"] = WasteCollecteds
	} else {
		res := make([]models.WasteCollected, 0)
		m["WasteCollected"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// func (h *Handler) WasteCollectedAssign(w http.ResponseWriter, r *http.Request) {
// 	platform := r.URL.Query().Get("platform")
// 	WasteCollectedAssign := new(models.WasteCollectedAssign)
// 	ctx := app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())
// 	err := json.NewDecoder(r.Body).Decode(&WasteCollectedAssign)
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	defer r.Body.Close()

// 	err = h.Service.WasteCollectedAssign(ctx, WasteCollectedAssign)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["WasteCollected"] = WasteCollectedAssign
// 	response.With200V2(w, "Success", m, platform)
// }
// func (h *Handler) RevokeWasteCollected(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")

// 	WasteCollected := new(models.WasteCollected)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())

// 	err := json.NewDecoder(r.Body).Decode(&WasteCollected)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	if WasteCollected.UniqueID == "" {
// 		response.With400V2(w, "id is missing", platform)
// 	}
// 	err = h.Service.RevokeWasteCollected(ctx, WasteCollected)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["WasteCollected"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }
