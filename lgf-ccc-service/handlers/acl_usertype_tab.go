package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"net/http"
)

//SaveACLUserTypeTabMultiple : ""
func (h *Handler) SaveACLUserTypeTabMultiple(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	modules := []models.ACLUserTypeTab{}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&modules)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveACLUserTypeTabMultiple(ctx, modules)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["update"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleUserTypeTabAccess :""
func (h *Handler) GetSingleUserTypeTabAccess(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	userTypeID := r.URL.Query().Get("userTypeId")
	moduleID := r.URL.Query().Get("moduleId")

	if userTypeID == "" {
		response.With400V2(w, "user type id is missing", platform)
	}

	if moduleID == "" {
		response.With400V2(w, "module id is missing", platform)
	}
	module := new(models.UserTypeTabAccess)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	module, err := h.Service.GetSingleUserTypeTabAccess(ctx, userTypeID, moduleID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["userType"] = module
	response.With200V2(w, "Success", m, platform)
}
