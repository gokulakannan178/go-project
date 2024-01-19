package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//SendOTPConsumerLogin : ""
func (h *Handler) SendOTPConsumerLogin(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	mobileNo := r.URL.Query().Get("mobileNo")

	if mobileNo == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())

	err := h.Service.SendOTPConsumerLogin(ctx, mobileNo)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["otp"] = "sent successfully"
	response.With200V2(w, "Success", m, platform)
}

//ConsumerLoginValidateOTP : ""
func (h *Handler) ConsumerLoginValidateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
defer ctx.Client.Disconnect(r.Context())
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
	properties, err := h.Service.ConsumerLoginValidateOTP(ctx, otpLogin.Mobile, otpLogin.OTP)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["login"] = "success"
	m["properties"] = properties
	response.With200V2(w, "Success", m, platform)
}
