package handlers

import (
	"encoding/json"
	"log"
	"logikoof-echalan-service/app"
	"logikoof-echalan-service/models"
	"logikoof-echalan-service/response"
	"net/http"
	"strconv"
)

//SaveUser : ""
func (h *Handler) SaveUser(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	user := new(models.User)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveUser(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = user
	response.With200V2(w, "Success", m, platform)
}

//UpdateUser :""
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	user := new(models.User)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if user.UserName == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateUser(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableUser : ""
func (h *Handler) EnableUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableUser(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableUser : ""
func (h *Handler) DisableUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableUser(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteUser : ""
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteUser(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleUser :""
func (h *Handler) GetSingleUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	user := new(models.RefUser)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	user, err := h.Service.GetSingleUser(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = user
	response.With200V2(w, "Success", m, platform)
}

//FilterUser : ""
func (h *Handler) FilterUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var user *models.UserFilter
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
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var users []models.RefUser
	log.Println(pagination)
	users, err = h.Service.FilterUser(ctx, user, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(users) > 0 {
		m["user"] = users
	} else {
		res := make([]models.User, 0)
		m["user"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//ResetUserPassword : ""
func (h *Handler) ResetUserPassword(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	code := r.URL.Query().Get("id")
	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.ResetUserPassword(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//ChangePassword : ""
func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	cp := new(models.UserChangePassword)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&cp)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if cp.UserName == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}
	ok, msg, err := h.Service.ChangePassword(ctx, cp)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if !ok {
		response.With403mV2(w, msg, platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//ConsumerLoginSendOTP : ""
func (h *Handler) ConsumerLoginSendOTP(w http.ResponseWriter, r *http.Request) {
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	platform := r.URL.Query().Get("platform")
	regNo := r.URL.Query().Get("regNo")
	if regNo == "" {
		response.With400V2(w, "regNo no is missing", platform)
		return
	}
	err := h.Service.ConsumerLoginSendOTP(ctx, regNo)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//ConsumerLoginValidateOTP : ""
func (h *Handler) ConsumerLoginValidateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	var otpLogin *models.OTPLogin
	err := json.NewDecoder(r.Body).Decode(&otpLogin)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if otpLogin.Mobile == "" || otpLogin.OTP == "" {
		response.With400V2(w, "mobile no or otp is missing", platform)
		return
	}
	err = h.Service.ConsumerLoginValidateOTP(ctx, otpLogin.Mobile, otpLogin.OTP)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["login"] = "success"
	response.With200V2(w, "Success", m, platform)
}
