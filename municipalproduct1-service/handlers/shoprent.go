package handlers

import (
	"encoding/json"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

//SaveShopRent : ""
func (h *Handler) SaveShopRent(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	shoprent := new(models.ShopRent)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&shoprent)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveShopRent(ctx, shoprent)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	/* Made by solomon
	 * Updating overall shoprent demand
	 * Previously UpdateShopRentDemandAndCollections was called to update demand and collection
	 * But new update is to call CalcShopRentOverallMonthlyDemand - which is written with shared transaction
	 * Given third parameter true to update in choprent collection
	 */
	/*var ctx1 *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	_, err = h.Service.CalcShopRentOverallMonthlyDemand(ctx1, shoprent.UniqueID, true)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	*/
	m := make(map[string]interface{})
	m["shoprent"] = shoprent
	response.With200V2(w, "Success", m, platform)
}

//UpdateShopRent :""
func (h *Handler) UpdateShopRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	shoprent := new(models.ShopRent)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&shoprent)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if shoprent.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateShopRent(ctx, shoprent)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	err = h.Service.UpdateShopRentDemandAndCollections(ctx, shoprent.UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableShopRent : ""
func (h *Handler) EnableShopRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableShopRent(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//RejectedShopRent : ""
func (h *Handler) RejectedShopRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.RejectedShopRent(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableShopRent : ""
func (h *Handler) DisableShopRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableShopRent(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteShopRent : ""
func (h *Handler) DeleteShopRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteShopRent(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleShopRent :""
func (h *Handler) GetSingleShopRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	shoprent := new(models.RefShopRent)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	shoprent, err := h.Service.GetSingleShopRent(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = shoprent
	response.With200V2(w, "Success", m, platform)
}

//FilterShopRent : ""
func (h *Handler) FilterShopRent(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.ShopRentFilter
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

	var shopRents []models.RefShopRent
	log.Println(pagination)
	shopRents, err = h.Service.FilterShopRent(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(shopRents) > 0 {
		m["shoprent"] = shopRents
	} else {
		res := make([]models.ShopRent, 0)
		m["shoprent"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//VerifyShopRentPayment : ""
func (h *Handler) VerifyShopRentPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeShopRentPaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	shoprentID, err := h.Service.VerifyShopRentPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	err = h.Service.UpdateShopRentDemandAndCollections(ctx, shoprentID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["verifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//NotVerifyShopRentPayment : ""
func (h *Handler) NotVerifyShopRentPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeShopRentPaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	shoprentID, err := h.Service.NotVerifyShopRentPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	err = h.Service.UpdateShopRentDemandAndCollections(ctx, shoprentID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["notVerifyPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//RejectShopRentPayment : ""
func (h *Handler) RejectShopRentPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	action := new(models.MakeShopRentPaymentsAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&action)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	shoprentID, err := h.Service.RejectShopRentPayment(ctx, action)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	err = h.Service.UpdateShopRentDemandAndCollections(ctx, shoprentID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["rejectPayment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//UpdateShopRentDemandAndCollections : ""
func (h *Handler) UpdateShopRentDemandAndCollections(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := h.Service.UpdateShopRentDemandAndCollections(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
