package handlers

import (
	"encoding/json"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

//OTPLoginGenerateOTP : "Login user"
func (h *Handler) LoginGenerateotpFarmer(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	farmer := new(models.FarmerLogin)
	err := json.NewDecoder(r.Body).Decode(&farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}

	err = h.Service.LoginGenerateotpFarmer(ctx, farmer)

	if err != nil {

		response.With500mV2(w, err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["otp"] = "Otp Sent Succesfully"

	response.With200V2(w, "Success", m, platform)
}

//OTPLoginValidateOTP : "Login user using OTP"
func (h *Handler) LoginValidateOTPFarmer(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	farmer := new(models.FarmerOTPLogin)
	err := json.NewDecoder(r.Body).Decode(&farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}

	farmerdata, stat, err := h.Service.LoginValidateOTPFarmer(ctx, farmer)

	if err != nil {
		if err.Error() == constants.NOTFOUND {
			response.With403mV2(w, "Invalid Farmer", platform)
			return
		}
		response.With500mV2(w, err.Error(), platform)
		return
	}
	if !stat {
		response.With403mV2(w, "Invalid farmername or Password", platform)
		return
	}

	m := make(map[string]interface{})
	m["farmer"] = farmerdata
	response.With200V2(w, "Success", m, platform)
}
