package handlers

import (
	"ecommerce-service/app"
	"ecommerce-service/models"
	"ecommerce-service/response"
	"encoding/json"
	"log"

	"net/http"
	"strconv"
)

// SaveInventory : ""
func (h *Handler) SaveInventory(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	block := new(models.Inventory)
	err := json.NewDecoder(r.Body).Decode(&block)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SaveInventory(ctx, block)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["inventory"] = block
	response.With200V2(w, "Success", m, platform)
}

// GetSingleInventory : ""
func (h *Handler) GetSingleInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	crop := new(models.RefInventory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	crop, err := h.Service.GetSingleInventory(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["inventory"] = crop
	response.With200V2(w, "Success", m, platform)
}

// UpdateInventory : ""
func (h *Handler) UpdateInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	crop := new(models.Inventory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&crop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if crop.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateInventory(ctx, crop)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["inventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// UpdateInventoryQuantityDetails : ""
func (h *Handler) UpdateInventoryQuantityDetails(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	crop := new(models.Inventory)
	//var ctx *models.Context
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&crop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if crop.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateInventoryQuantityDetails(ctx, crop)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["inventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableInventory : ""
func (h *Handler) EnableInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableInventory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Inventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DisableInventory : ""
func (h *Handler) DisableInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableInventory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["inventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// DeleteInventory : ""
func (h *Handler) DeleteInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteInventory(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["inventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// FilterInventory : ""
func (h *Handler) FilterInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.InventoryFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var Inventorys []models.RefInventory
	log.Println(pagination)
	Inventorys, err = h.Service.FilterInventory(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Inventorys) > 0 {
		m["inventory"] = Inventorys
	} else {
		res := make([]models.Inventory, 0)
		m["inventory"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// CreateMesh : ""
// func (h *Handler) CreateMesh(w http.ResponseWriter, r *http.Request) {
// 	platform := r.URL.Query().Get("platform")
// 	Imc := new(models.InventoryMeshCreate)
// 	err := json.NewDecoder(r.Body).Decode(&Imc)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}

// 	ctx := app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())
// 	mesh, err := h.Service.CreateMesh(ctx, Imc)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["Inventory"] = mesh
// 	response.With200V2(w, "Success", m, platform)
// }

// ImageInventory : ""
func (h *Handler) ImageInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	crop := new(models.Inventory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&crop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if crop.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.ImageInventory(ctx, crop)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["inventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// ImagesInventory : ""
func (h *Handler) ImagesInventory(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	crop := new(models.Inventory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&crop)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if crop.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.ImagesInventory(ctx, crop)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["inventory"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetbyBarcodeAndVendor(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	Barcode := r.URL.Query().Get("barcode")
	Vendor := r.URL.Query().Get("vendor")

	if Barcode == "" {
		response.With400V2(w, "barcode is missing", platform)
		return
	}
	if Vendor == "" {
		response.With400V2(w, "vendor is missing", platform)
		return
	}

	inventory := new(models.RefInventory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	inventory, err := h.Service.GetbyBarcodeAndVendor(ctx, Barcode, Vendor)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["inventory"] = inventory
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) ChkUniqueness(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	Barcode := r.URL.Query().Get("barcode")
	Vendor := r.URL.Query().Get("vendor")

	if Barcode == "" {
		response.With400V2(w, "barcode is missing", platform)
		return
	}
	if Vendor == "" {
		response.With400V2(w, "vendor is missing", platform)
		return
	}

	inventory := new(models.RefInventory)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	inventory, err := h.Service.ChkUniqueness(ctx, Barcode, Vendor)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	if inventory != nil {
		m["inventory"] = inventory
	} else {
		res := make([]models.Inventory, 0)
		m["inventory"] = res
	}
	m["inventory"] = inventory
	response.With200V2(w, "Success", m, platform)
}
