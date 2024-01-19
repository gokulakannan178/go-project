package handlers

import (
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

//GetContentDisseminationUserAndFarmer :""
func (h *Handler) GetContentDisseminationUserAndFarmer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	contentDissiminateUserAndFarmer := new(models.ContentDissiminateUserAndFarmer)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	contentDissiminateUserAndFarmer, err := h.Service.GetContentDisseminationUserAndFarmer(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["contentDissiminateUserAndFarmer"] = contentDissiminateUserAndFarmer
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetContentDisseminationUserAndFarmerCount(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var content *models.ContentDataAccess

	contentDissiminateUserAndFarmer := new(models.ContentDissiminateUserAndFarmer)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	contentDissiminateUserAndFarmer, err := h.Service.GetContentDisseminationUserAndFarmerCount(ctx, content)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["contentDissiminateUserAndFarmer"] = contentDissiminateUserAndFarmer
	response.With200V2(w, "Success", m, platform)
}
