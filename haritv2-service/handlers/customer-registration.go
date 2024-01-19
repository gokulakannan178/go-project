package handlers

import (
	"encoding/json"
	"haritv2-service/app"
	"haritv2-service/constants"
	"haritv2-service/models"
	"haritv2-service/response"
	"net/http"
)

//OTPLoginGenerateOTP : "Login user"
func (h *Handler) CustomerregistrationGenerateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	customer := new(models.RegistrationCustomer)
	err := json.NewDecoder(r.Body).Decode(&customer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}
	if customer.Mobile == "" {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" -invalid mobile number", platform)
		return
	}
	err = h.Service.CustomerregistrationGenerateOTP(ctx, customer)

	if err != nil {

		response.With500mV2(w, err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["otp"] = "Otp Sent Succesfully"

	response.With200V2(w, "Success", m, platform)
}

//OTPLoginValidateOTP : "Login user using OTP"
func (h *Handler) CustomerregistrationValidateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	customer := new(models.RegistrationCustomer)
	err := json.NewDecoder(r.Body).Decode(&customer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}
	if customer.Mobile == "" {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" -invalid mobile number", platform)
		return
	}
	if customer.OTP == "" {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" -invalid  otp", platform)
		return
	}

	consumerdata, stat, err := h.Service.CustomerregistrationValidateOTP(ctx, customer)

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
	m["customer"] = consumerdata

	response.With200V2(w, "Success", m, platform)
}
