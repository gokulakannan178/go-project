package handlers

import (
	"encoding/json"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
)

//SavePunchIn : ""
func (h *Handler) SavePunchIn(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	user := new(models.UserAttendanceAction)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePunchIn(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userAttendance"] = user
	response.With200V2(w, "Success", m, platform)
}

//SavePunchOut : ""
func (h *Handler) SavePunchOut(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	// user.UniqueID := r.URL.Query().Get("id")
	user := new(models.UserAttendance)
	if user.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.SavePunchOut(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userAttendance"] = "success"
	response.With200V2(w, "Success", m, platform)
}
