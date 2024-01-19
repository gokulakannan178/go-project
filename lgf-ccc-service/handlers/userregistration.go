package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"net/http"
)

//OTPLoginGenerateOTP : "Login user"
func (h *Handler) UserRegistrationGenerateOTP(w http.ResponseWriter, r *http.Request) {
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
	err = h.Service.UserRegistrationGenerateOTP(ctx, consumer)

	if err != nil {

		response.With500mV2(w, err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["otp"] = "Otp Sent Succesfully"

	response.With200V2(w, "Success", m, platform)
}

//OTPLoginValidateOTP : "Login user using OTP"
func (h *Handler) UserRegistrationValidateOTP(w http.ResponseWriter, r *http.Request) {
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

	consumerdata, stat, err := h.Service.UserRegistrationValidateOTP(ctx, consumer)

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
