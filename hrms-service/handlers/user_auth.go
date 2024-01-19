package handlers

import (
	"encoding/json"
	"hrms-services/app"
	"hrms-services/constants"
	"hrms-services/models"
	"hrms-services/response"
	"log"
	"net/http"
)

//LoginV2 : "Login user"
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
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
	respUser, err := h.Service.GetSingleUserWithLogin(ctx, user.UserName)
	if err != nil {
		log.Println("err=>", err.Error())
	}
	m := make(map[string]interface{})
	m["token"] = token
	m["user"] = respUser
	// m["role"] = role
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleProfile(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	//	billClaim := new(models.RefBillClaim)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	user, err := h.Service.GetSingleProfile(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = user
	response.With200V2(w, "Success", m, platform)
}
