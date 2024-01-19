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
func (h *Handler) DealerregistrationGenerateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	dealer := new(models.RegistrationDealer)
	err := json.NewDecoder(r.Body).Decode(&dealer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}
	if dealer.Mobile == "" {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" -invalid mobile number", platform)
		return
	}
	err = h.Service.DealerregistrationGenerateOTP(ctx, dealer)

	if err != nil {

		response.With500mV2(w, err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["otp"] = "Otp Sent Succesfully"

	response.With200V2(w, "Success", m, platform)
}

//OTPLoginValidateOTP : "Login user using OTP"
func (h *Handler) DealerregistrationValidateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	dealer := new(models.RegistrationDealer)
	err := json.NewDecoder(r.Body).Decode(&dealer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}
	if dealer.Mobile == "" {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" -invalid mobile number", platform)
		return
	}
	if dealer.OTP == "" {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" -invalid  otp", platform)
		return
	}

	Dealerdata, stat, err := h.Service.DealerregistrationValidateOTP(ctx, dealer)

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
	m["dealer"] = Dealerdata

	response.With200V2(w, "Success", m, platform)
}
