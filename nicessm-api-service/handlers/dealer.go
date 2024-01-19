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

//SaveDealer : ""
func (h *Handler) SaveDealer(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	dealer := new(models.Dealer)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&dealer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveDealer(ctx, dealer)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dealer"] = dealer
	response.With200V2(w, "Success", m, platform)
}

//UpdateDealer :""
func (h *Handler) UpdateDealer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	dealer := new(models.Dealer)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&dealer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if dealer.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateDealer(ctx, dealer)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableDealer : ""
func (h *Handler) EnableDealer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableDealer(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableDealer : ""
func (h *Handler) DisableDealer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableDealer(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteDealer : ""
func (h *Handler) DeleteDealer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteDealer(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleDealer :""
func (h *Handler) GetSingleDealer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	dealer := new(models.RefDealer)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	dealer, err := h.Service.GetSingleDealer(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = dealer
	response.With200V2(w, "Success", m, platform)
}

//FilterDealer : ""
func (h *Handler) FilterDealer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var dealer *models.DealerFilter
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
	err := json.NewDecoder(r.Body).Decode(&dealer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var dealers []models.RefDealer
	log.Println(pagination)
	dealers, err = h.Service.FilterDealer(ctx, dealer, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(dealers) > 0 {
		m["data"] = dealers
	} else {
		res := make([]models.Dealer, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DealerUniquenessCheckRegistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	Param := r.URL.Query().Get("param")
	Value := r.URL.Query().Get("value")
	if Param == "" || Value == "" {
		response.With400V2(w, "Param/Value is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	pass, err := h.Service.DealerUniquenessCheckRegistration(ctx, Param, Value)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["status"] = pass
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DealerNearBy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	dealernb := new(models.NearBy)

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&dealernb)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	//km!=0
	if dealernb.KM == 0 {
		response.With400V2(w, "id is missing", platform)
	}
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
	dealers, err := h.Service.DealerNearBy(ctx, dealernb, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dealerNearby"] = dealers
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DealerCertificationApply(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	DealerID := r.URL.Query().Get("id")
	certification := new(models.DealerCertification)

	if DealerID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DealerCertificationApply(ctx, DealerID, certification)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DealerCertificationApprove(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	dealerID := r.URL.Query().Get("id")
	certification := new(models.DealerCertification)
	if dealerID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DealerCertificationApprove(ctx, dealerID, certification)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) DealerCertificationReject(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	dealerID := r.URL.Query().Get("id")
	certification := new(models.DealerCertification)

	if dealerID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DealerCertificationReject(ctx, dealerID, certification)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
