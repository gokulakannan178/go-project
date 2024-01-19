package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

//CreatedBankDeposit : ""
func (h *Handler) CreatedBankDeposit(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	bankDeposit := new(models.BankDeposit)
	err := json.NewDecoder(r.Body).Decode(&bankDeposit)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err = h.Service.CreatedBankDeposit(ctx, bankDeposit)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankDeposit"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//VerifyBankDeposit
func (h *Handler) VerifyBankDeposit(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	uniqueID := r.URL.Query().Get("id")
	bdv := new(models.BankDepositVerifier)
	err := json.NewDecoder(r.Body).Decode(&bdv)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err = h.Service.VerifierBankDeposit(ctx, uniqueID, bdv)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Verify"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleBankDeposit
func (h *Handler) GetSingleBankDeposit(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	property := new(models.BankDeposit)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	property, err := h.Service.GetSingleBankDeposit(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = property
	response.With200V2(w, "Success", m, platform)
}

//NotVerifyBankDeposit
func (h *Handler) NotVerifyBankDeposit(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	uniqueID := r.URL.Query().Get("id")
	bdv := new(models.BankDepositVerifier)
	err := json.NewDecoder(r.Body).Decode(&bdv)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
	err = h.Service.NotVerifierBankDeposit(ctx, uniqueID, bdv)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Verify"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//BankDepositFilter
func (h *Handler) BankDepositFilter(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	//uniqueID := r.URL.Query().Get("id")
	//var user *models.UserFilter
	bdf := new(models.BankDepositFilter)
	err := json.NewDecoder(r.Body).Decode(&bdf)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	var pagination *models.Pagination
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
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

	data, err := h.Service.BankDepositFilter(ctx, bdf, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankdeposit"] = data
	m["pagination"] = pagination
	response.With200V2(w, "Success", m, platform)
}
