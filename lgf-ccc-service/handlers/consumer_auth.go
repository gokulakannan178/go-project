package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
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

func (h *Handler) CitizenregistrationGenerateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	consumer := new(models.RegistrationUser)
	err := json.NewDecoder(r.Body).Decode(&consumer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}
	if consumer.Mobile == "" {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" -invalid mobile number", platform)
		return
	}
	err = h.Service.CitizenregistrationGenerateOTP(ctx, consumer)

	if err != nil {

		response.With500mV2(w, err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["otp"] = "Otp Sent Succesfully"

	response.With200V2(w, "Success", m, platform)
}

//OTPLoginValidateOTP : "Login user using OTP"
func (h *Handler) CitizenregistrationValidateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	consumer := new(models.RegistrationUser)
	err := json.NewDecoder(r.Body).Decode(&consumer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}
	if consumer.Mobile == "" {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" -invalid mobile number", platform)
		return
	}
	if consumer.OTP == "" {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" -invalid  otp", platform)
		return
	}

	consumerdata, stat, err := h.Service.CitizenregistrationValidateOTP(ctx, consumer)

	if err != nil {
		if err.Error() == constants.NOTFOUND {
			response.With403mV2(w, "Invalid User", platform)
			return
		}
		response.With500mV2(w, err.Error(), platform)
		return
	}
	if !stat {
		response.With403mV2(w, "Invalid Username or Password", platform)
		return
	}

	m := make(map[string]interface{})
	m["consumer"] = consumerdata

	response.With200V2(w, "Success", m, platform)
}
