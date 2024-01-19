package handlers

import (
	"bpms-service/app"
	"bpms-service/models"
	"bpms-service/response"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//SavePreregistration : ""
func (h *Handler) SavePreregistration(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	preregistration := new(models.Preregistration)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&preregistration)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if preregistration.MobileNumber == "" {
		response.With400V2(w, "please provide valid mobile number", platform)
		return
	}
	if !isEmailValid(preregistration.Email) {
		response.With400V2(w, "please provide valid email", platform)
		return
	}
	err = h.Service.SavePreregistration(ctx, preregistration)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["preregistration"] = preregistration
	response.With200V2(w, "Success", m, platform)
}

//SubmitPreregistration : ""
func (h *Handler) SubmitPreregistration(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	preregistration := new(models.Preregistration)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&preregistration)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if preregistration.MobileNumber == "" {
		response.With400V2(w, "please provide valid mobile number", platform)
		return
	}
	if !isEmailValid(preregistration.Email) {
		response.With400V2(w, "please provide valid email", platform)
		return
	}
	err = h.Service.SubmitPreregistration(ctx, preregistration)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["preregistration"] = preregistration
	response.With200V2(w, "Success", m, platform)
}

//ReapplyPreregistration : ""
func (h *Handler) ReapplyPreregistration(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	preregistration := new(models.Preregistration)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&preregistration)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if preregistration.MobileNumber == "" {
		response.With400V2(w, "please provide valid mobile number", platform)
		return
	}
	if !isEmailValid(preregistration.Email) {
		response.With400V2(w, "please provide valid email", platform)
		return
	}
	err = h.Service.ReapplyPreregistration(ctx, preregistration)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["preregistration"] = preregistration
	response.With200V2(w, "Success", m, platform)
}

//UpdatePreregistration :""
func (h *Handler) UpdatePreregistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	preregistration := new(models.Preregistration)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&preregistration)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	id := preregistration.UniqueID
	email := preregistration.Email
	mobile := preregistration.MobileNumber

	if id == "" {
		response.With400V2(w, "unique id is missing", platform)
		return
	}
	if preregistration.Email != "" {
		err := h.validateEmailAtUpdate(ctx, id, email)
		if err != nil {
			response.With400V2(w, err.Error(), platform)
			return
		}
	}

	if preregistration.MobileNumber != "" {
		err := h.validateMobileAtUpdate(ctx, id, mobile)
		if err != nil {
			response.With400V2(w, err.Error(), platform)
			return
		}
	}

	err = h.Service.UpdatePreregistration(ctx, preregistration)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnablePreregistration : ""
func (h *Handler) EnablePreregistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnablePreregistration(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisablePreregistration : ""
func (h *Handler) DisablePreregistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	uniqueID := r.URL.Query().Get("id")

	if uniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisablePreregistration(ctx, uniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeletePreregistration : ""
func (h *Handler) DeletePreregistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeletePreregistration(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePreregistration :""
func (h *Handler) GetSinglePreregistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	mobileNumber := r.URL.Query().Get("mobileNumber")

	if mobileNumber == "" {
		response.With400V2(w, "mobile number is missing", platform)
		return
	}

	if len(mobileNumber) != 10 {
		response.With400V2(w, "please provide valid mobile number", platform)
		return
	}

	preregistration := new(models.RefPreregistration)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	preregistration, err := h.Service.GetSinglePreregistration(ctx, mobileNumber)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = preregistration
	response.With200V2(w, "Success", m, platform)
}

//FilterPreregistration : ""
func (h *Handler) FilterPreregistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var preregistration *models.PreregistrationFilter
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
	err := json.NewDecoder(r.Body).Decode(&preregistration)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var preregistrations []models.RefPreregistration
	log.Println(pagination)
	preregistrations, err = h.Service.FilterPreregistration(ctx, preregistration, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(preregistrations) > 0 {
		m["data"] = preregistrations
	} else {
		res := make([]models.Preregistration, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//ValidateMobileNumber :""
func (h *Handler) ValidateMobileNumber(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	mobileNumber := r.URL.Query().Get("mobileNumber")
	uniqueID := r.URL.Query().Get("id")

	if mobileNumber == "" {
		response.With400V2(w, "mobile number is missing", platform)
		return
	}
	if len(mobileNumber) != 10 {
		response.With400V2(w, "please provide valid mobile number", platform)
		return
	}

	if uniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	preregistration := new(models.RefPreregistration)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	isValid, preregistration, err := h.Service.ValidateMobileNumber(ctx, mobileNumber, uniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = preregistration
	m["isValid"] = isValid
	response.With200V2(w, "Success", m, platform)
}

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

// validateEmailAtUpdate = ""
func (h *Handler) validateEmailAtUpdate(ctx *models.Context, uniqueID, email string) error {
	if !isEmailValid(email) {
		return errors.New("provide valid emailId")
	}
	err := h.Service.ValidateEmailAtUpdate(ctx, email, uniqueID)
	if err != nil {
		return err
	}
	return nil
}

// validateMobileAtUpdate = ""
func (h *Handler) validateMobileAtUpdate(ctx *models.Context, uniqueID, mobile string) error {
	if len(mobile) != 10 {
		return errors.New("provide valid mobile number")
	}
	err := h.Service.ValidateMobileAtUpdate(ctx, mobile, uniqueID)
	if err != nil {
		return err
	}

	return nil
}

//GetSinglePreregistration :""
func (h *Handler) GetSinglePreregistrationWithUniqueID(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	uniqueID := r.URL.Query().Get("id")

	if uniqueID == "" {
		response.With400V2(w, "id  is missing", platform)
		return
	}

	preregistration := new(models.RefPreregistration)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	preregistration, err := h.Service.GetSinglePreregistrationWithUniqueID(ctx, uniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = preregistration
	response.With200V2(w, "Success", m, platform)
}

//PreregistrationStatusChange : ""
func (h *Handler) PreregistrationStatusChange(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	psc := new(models.PreregistrationStatusChange)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&psc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.PreregistrationStatusChange(ctx, psc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["preregistration"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//PaymentPreregistration : ""
func (h *Handler) PaymentPreregistration(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	prp := new(models.PreregistrationPayment)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&prp)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.PaymentPreregistration(ctx, prp)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//PaymentPendingNoticeForPreRegistration : ""
func (h *Handler) PaymentPendingNoticeForPreRegistration(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	mobile := r.URL.Query().Get("id")
	if mobile == "" {
		response.With400V2(w, "id  is missing", platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := h.Service.PaymentPendingNoticeForPreRegistration(ctx, mobile)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["sendNotice"] = "success"
	response.With200V2(w, "Success", m, platform)
}
