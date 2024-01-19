package handlers

import (
	"bpms-service/app"
	"bpms-service/models"
	"bpms-service/response"
	"encoding/json"
	"net/http"
)

//ApplicantLoginSendOTP : ""
func (h *Handler) ApplicantLoginSendOTP(w http.ResponseWriter, r *http.Request) {
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	platform := r.URL.Query().Get("platform")
	mobile := r.URL.Query().Get("mobile")
	if mobile == "" {
		response.With400V2(w, "mobile number is missing", platform)
		return
	}
	err := h.Service.ApplicantLoginSendOTP(ctx, mobile)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//ApplicantLoginValidateOTP : ""
func (h *Handler) ApplicantLoginValidateOTP(w http.ResponseWriter, r *http.Request) {
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
	err = h.Service.ApplicantLoginValidateOTP(ctx, otpLogin.Mobile, otpLogin.OTP)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["login"] = "success"
	response.With200V2(w, "Success", m, platform)
}
