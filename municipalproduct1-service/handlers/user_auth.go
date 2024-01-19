package handlers

import (
	"encoding/json"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//LoginV2 : "Login user"
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	user := new(models.Login)
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}

	token, stat, err := h.Service.Login(ctx, user)
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
	respUser, err := h.Service.GetSingleUser(ctx, user.UserName)
	if err != nil {
		log.Println("err=>", err.Error())
	}
	if respUser.AccessPrivilege == nil {
		respUser.AccessPrivilege = new(models.AccessPrivilege)
		respUser.AccessPrivilege.Wards = make([]string, 0)
	}
	m := make(map[string]interface{})
	m["token"] = token
	m["user"] = respUser
	// m["role"] = role
	response.With200V2(w, "Success", m, platform)
}

//LoginGenerateOTP : "Login user"
func (h *Handler) LoginGenerateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	mobileNo := r.URL.Query().Get("mobileNo")
	if mobileNo == "" {
		response.With400V2(w, "Mobile No is missing", platform)
	}
	err := h.Service.LoginGenerateOTP(ctx, mobileNo)
	if err != nil {
		response.With500mV2(w, err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["otp"] = "sent"
	response.With200V2(w, "Success", m, platform)
}

//LoginValidateOTP : "Login user"
func (h *Handler) LoginValidateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	mobileNo := r.URL.Query().Get("mobileNo")
	if mobileNo == "" {
		response.With400V2(w, "Mobile No is missing", platform)
	}
	otp := r.URL.Query().Get("otp")
	if mobileNo == "" {
		response.With400V2(w, "otp is missing", platform)
	}
	user, err := h.Service.LoginValidateOTP(ctx, mobileNo, otp)
	if err != nil {
		response.With500mV2(w, err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["otp"] = "sent"
	m["user"] = user
	response.With200V2(w, "Success", m, platform)
}

//LoginValidateOTPV2 : "Login user"
func (h *Handler) LoginValidateOTPV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	mobileNo := r.URL.Query().Get("mobileNo")
	if mobileNo == "" {
		response.With400V2(w, "Mobile No is missing", platform)
	}
	otp := r.URL.Query().Get("otp")
	if mobileNo == "" {
		response.With400V2(w, "otp is missing", platform)
	}
	user, err := h.Service.LoginValidateOTP(ctx, mobileNo, otp)
	if err != nil {
		response.With500mV2(w, err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["otp"] = "sent"
	m["user"] = user
	response.With200V2(w, "Success", m, platform)
}

//LoginGenerateOTP : "Login user"
func (h *Handler) LoginGenerateOTPV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	mobileNo := r.URL.Query().Get("mobileNo")
	versionKey := r.URL.Query().Get("versionKey")
	if mobileNo == "" {
		response.With400V2(w, "Mobile No is missing", platform)
		return
	}
	if versionKey == "" {
		response.With500mV2(w, "please update app version", platform)
		return
	}

	err := h.Service.LoginGenerateOTPV2(ctx, mobileNo, versionKey)
	if err != nil {
		response.With500mV2(w, err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["otp"] = "sent"
	response.With200V2(w, "Success", m, platform)
}
