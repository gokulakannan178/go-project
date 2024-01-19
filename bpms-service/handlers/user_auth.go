package handlers

import (
	"bpms-service/app"
	"bpms-service/constants"
	"bpms-service/models"
	"bpms-service/response"
	"encoding/json"
	"net/http"
)

//Login : ""
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	login := new(models.Login)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&login)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	token, ok, user, err := h.Service.Login(ctx, login)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if err != nil {
		if err.Error() == constants.NOTFOUND {
			response.With403mV2(w, "Invalid User", platform)
			return
		}
		response.With500mV2(w, err.Error(), platform)
		return
	}
	if !ok {
		response.With403mV2(w, "Invalid Username or Password", platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = user
	m["token"] = token
	response.With200V2(w, "Success", m, platform)
}
