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

//SaveBankInformation : ""
func (h *Handler) SaveBankInformation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	bankInformation := new(models.BankInformation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&bankInformation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveBankInformation(ctx, bankInformation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankInformation"] = bankInformation
	response.With200V2(w, "Success", m, platform)
}

//UpdateBankInformation :""
func (h *Handler) UpdateBankInformation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	bankInformation := new(models.BankInformation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&bankInformation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if bankInformation.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateBankInformation(ctx, bankInformation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankInformation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableBankInformation : ""
func (h *Handler) EnableBankInformation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableBankInformation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankInformation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableBankInformation : ""
func (h *Handler) DisableBankInformation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableBankInformation(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankInformation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteBankInformation : ""
func (h *Handler) DeleteBankInformation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteBankInformation(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankInformation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleBankInformation :""
func (h *Handler) GetSingleBankInformation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	BankInformation := new(models.RefBankInformation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	BankInformation, err := h.Service.GetSingleBankInformation(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankInformation"] = BankInformation
	response.With200V2(w, "Success", m, platform)
}

//FilterBankInformation : ""
func (h *Handler) FilterBankInformation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var filter *models.BankInformationFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var bankInformations []models.RefBankInformation
	log.Println(pagination)
	bankInformations, err = h.Service.FilterBankInformation(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(bankInformations) > 0 {
		m["bankInformation"] = bankInformations
	} else {
		res := make([]models.BankInformation, 0)
		m["bankInformation"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UpdateEmployeeBankInformation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	bankInformation := new(models.BankInformation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&bankInformation)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if bankInformation.EmployeeID == "" {
		response.With400V2(w, "employee id is missing", platform)
	}
	err = h.Service.UpdateEmployeeBankInformation(ctx, bankInformation)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankInformation"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleBankInformationWithEmployee(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("employeeId")

	if UniqueID == "" {
		response.With400V2(w, "employeeId is missing", platform)
	}

	BankInformation := new(models.RefBankInformation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	BankInformation, err := h.Service.GetSingleBankInformationWithEmployee(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["bankInformation"] = BankInformation
	response.With200V2(w, "Success", m, platform)
}
