package handlers

import (
	"encoding/json"

	"haritv2-service/app"
	"haritv2-service/models"
	"haritv2-service/response"
	"net/http"
)

// SaveFPOInventory : ""
func (h *Handler) SaveFPOInventory(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	fpoinventory := new(models.FPOInventory)
	err := json.NewDecoder(r.Body).Decode(&fpoinventory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveFPOInventory(ctx, fpoinventory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpoinventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GetSingleFPOInventory : ""
func (h *Handler) GetSingleFPOInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	fpoinventory := new(models.RefFPOINVENTORY)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	fpoinventory, err := h.Service.GetSingleFPOInventoryWithCompalyID(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpoinventory"] = fpoinventory
	response.With200V2(w, "Success", m, platform)
}

// UpdateFPOInventory : ""
func (h *Handler) UpdateFPOInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	fpoinventory := new(models.FPOInventory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&fpoinventory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if fpoinventory.CompanyID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateFPOInventory(ctx, fpoinventory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpoinventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableFPOInventory : ""
func (h *Handler) EnableFPOInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableFPOInventory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpoinventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableFPOInventory : ""
func (h *Handler) DisableFPOInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableFPOInventory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpoinventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteFPOInventory : ""
func (h *Handler) DeleteFPOInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteFPOInventory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpoinventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// // FilterFPOInventory : ""
// func (h *Handler) FilterFPOInventory(w http.ResponseWriter, r *http.Request) {

// 	platform := r.URL.Query().Get("platform")
// 	var filter *models.FPOInventoryFilter
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())
// 	pageNo := r.URL.Query().Get("pageno")
// 	Limit := r.URL.Query().Get("limit")

// 	var pagination *models.Pagination
// 	if pageNo != "no" {
// 		pagination = new(models.Pagination)
// 		if pagination.PageNum = 1; pageNo != "" {
// 			page, err := strconv.Atoi(pageNo)
// 			if pagination.PageNum = 1; err == nil {
// 				pagination.PageNum = page
// 			}
// 		}
// 		if pagination.Limit = 10; Limit != "" {
// 			limit, err := strconv.Atoi(Limit)
// 			if pagination.Limit = 10; err == nil {
// 				pagination.Limit = limit
// 			}
// 		}
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&filter)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}

// 	var FPOInventorys []models.RefFPOInventory
// 	log.Println(pagination)
// 	FPOInventorys, err = h.Service.FilterFPOInventory(ctx, filter, pagination)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})

// 	if len(FPOInventorys) > 0 {
// 		m["fpoinventory"] = FPOInventorys
// 	} else {
// 		res := make([]models.FPOInventory, 0)
// 		m["fpoinventory"] = res
// 	}
// 	if pagination != nil {
// 		if pagination.PageNum > 0 {
// 			m["pagination"] = pagination
// 		}
// 	}

// 	response.With200V2(w, "Success", m, platform)
// }

// UpdateFPOInventory : ""
func (h *Handler) FPOInventoryQuantityUpdate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	fpoinventory := new(models.FPOInventory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&fpoinventory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if fpoinventory.CompanyID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.FPOInventoryQuantityUpdate(ctx, fpoinventory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpoinventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FPOInventoryPriceUpdate : ""
func (h *Handler) FPOInventoryPriceUpdate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	fpoinventory := new(models.FPOInventory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&fpoinventory)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if fpoinventory.CompanyID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.FPOInventoryQuantityUpdate(ctx, fpoinventory)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["fpoinventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}
