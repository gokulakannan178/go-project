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

//SaveFarmerLiveStock : ""
func (h *Handler) SaveFarmerLiveStock(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	farmerLiveStock := new(models.FarmerLiveStock)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&farmerLiveStock)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveFarmerLiveStock(ctx, farmerLiveStock)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerLiveStock"] = farmerLiveStock
	response.With200V2(w, "Success", m, platform)
}

//UpdateFarmerLiveStock :""
func (h *Handler) UpdateFarmerLiveStock(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	farmerLiveStock := new(models.FarmerLiveStock)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&farmerLiveStock)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if farmerLiveStock.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateFarmerLiveStock(ctx, farmerLiveStock)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["FarmerLiveStock"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableFarmerLiveStock : ""
func (h *Handler) EnableFarmerLiveStock(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableFarmerLiveStock(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerLiveStock"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableFarmerLiveStock : ""
func (h *Handler) DisableFarmerLiveStock(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableFarmerLiveStock(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerLiveStock"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteFarmerLiveStock : ""
func (h *Handler) DeleteFarmerLiveStock(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteFarmerLiveStock(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerLiveStock"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleFarmerLiveStock :""
func (h *Handler) GetSingleFarmerLiveStock(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	farmerLiveStock := new(models.RefFarmerLiveStock)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	farmerLiveStock, err := h.Service.GetSingleFarmerLiveStock(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerLiveStock"] = farmerLiveStock
	response.With200V2(w, "Success", m, platform)
}

//FilterFarmerLiveStock : ""
func (h *Handler) FilterFarmerLiveStock(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var farmerLiveStock *models.FarmerLiveStockFilter
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
	err := json.NewDecoder(r.Body).Decode(&farmerLiveStock)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var farmerLiveStocks []models.RefFarmerLiveStock
	log.Println(pagination)
	farmerLiveStocks, err = h.Service.FilterFarmerLiveStock(ctx, farmerLiveStock, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(farmerLiveStocks) > 0 {
		m["farmerLiveStock"] = farmerLiveStocks
	} else {
		res := make([]models.FarmerLiveStock, 0)
		m["farmerLiveStock"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
