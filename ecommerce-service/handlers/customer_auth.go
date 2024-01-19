package handlers

import (
	"ecommerce-service/app"
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"ecommerce-service/response"
	"encoding/json"
	"log"
	"net/http"
)

//LoginV2 : "Login Customer"
func (h *Handler) CustomerLogin(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	customer := new(models.Login)
	err := json.NewDecoder(r.Body).Decode(&customer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}

	token, stat, err := h.Service.CustomerLogin(ctx, customer)
	log.Println("stat ==>", stat)
	//	log.Println("err ==>", err.Error())
	log.Println("TOKEN==>", token)
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
	respCustomer, err := h.Service.GetSingleCustomer(ctx, customer.UserName)
	if err != nil {
		log.Println("err=>", err.Error())
	}
	m := make(map[string]interface{})
	m["token"] = token
	m["customer"] = respCustomer
	// m["role"] = role
	response.With200V2(w, "Success", m, platform)
}

//OTPLoginGenerateOTP : "Login customer"
func (h *Handler) CustomerOTPLoginGenerateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	customer := new(models.Login)
	err := json.NewDecoder(r.Body).Decode(&customer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}

	err = h.Service.CustomerOTPLoginGenerateOTP(ctx, customer)

	if err != nil {

		response.With500mV2(w, err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["otp"] = "Otp Sent Succesfully"

	response.With200V2(w, "Success", m, platform)
}

//OTPLoginValidateOTP : "Login customer using OTP"
func (h *Handler) CustomerOTPLoginValidateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	customer := new(models.OTPLogin)
	err := json.NewDecoder(r.Body).Decode(&customer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}

	customerdata, stat, err := h.Service.CustomerOTPLoginValidateOTP(ctx, customer)

	if err != nil {
		if err.Error() == constants.NOTFOUND {
			response.With403mV2(w, "Invalid Customer", platform)
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
	m["customer"] = customerdata
	//fmt.Println("TOKENNN", customerdata.Token)

	response.With200V2(w, "Success", m, platform)
}
