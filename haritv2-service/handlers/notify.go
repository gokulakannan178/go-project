package handlers

import (
	"encoding/json"
	"haritv2-service/app"
	"haritv2-service/models"
	"haritv2-service/response"
	"net/http"
)

//NotifyForUpdateLocation : ""
func (h *Handler) NotifyForUpdateLocation(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	notificationType := r.URL.Query().Get("type")

	var ctx *models.Context

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	// err := json.NewDecoder(r.Body).Decode(&notificationType)
	defer r.Body.Close()
	// if err != nil {
	// 	response.With400V2(w, err.Error(), platform)
	// 	return
	// }
	h.Service.NotifyForUpdateLocation(ctx, notificationType)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["notifyforupdatelocation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//NotifyForUpdateLocation : ""
func (h *Handler) NotifyForUpdateProfile(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	notificationType := r.URL.Query().Get("type")

	var ctx *models.Context

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	// err := json.NewDecoder(r.Body).Decode(&notificationType)
	defer r.Body.Close()
	// if err != nil {
	// 	response.With400V2(w, err.Error(), platform)
	// 	return
	// }
	h.Service.NotifyForUpdateProfile(ctx, notificationType)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["notifyforupdatelocation"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//NotifyForULBInventoryUpdate : ""
func (h *Handler) NotifyForULBInventoryUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	notifyType := r.URL.Query().Get("type")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	// err := json.NewDecoder(r.Body).Decode(&notificationType)
	defer r.Body.Close()
	// if err != nil {
	// 	response.With400V2(w, err.Error(), platform)
	// 	return
	// }
	h.Service.NotifyForULBInventoryUpdate(ctx, notifyType)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["notifyforulbinventoryupdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//NotifyForULBInventoryUpdatev2 : "For Custom months"
func (h *Handler) NotifyForULBInventoryUpdateV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	filter := new(models.ULBInventoryUpdateMessageFilterV2)
	err := json.NewDecoder(r.Body).Decode(&filter)

	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer ctx.Client.Disconnect(r.Context())
	// err := json.NewDecoder(r.Body).Decode(&notificationType)
	defer r.Body.Close()
	// if err != nil {
	// 	response.With400V2(w, err.Error(), platform)
	// 	return
	// }
	h.Service.NotifyForULBInventoryUpdateV2(ctx, filter)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})
	m["notifyforulbinventoryupdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}
