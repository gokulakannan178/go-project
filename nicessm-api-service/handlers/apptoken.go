package handlers

import (
	"encoding/json"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"

	"log"
	"net/http"
	"strconv"
)

//SaveApptoken : ""
func (h *Handler) SaveApptoken(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	apptoken := new(models.Apptoken)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&apptoken)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveApptoken(ctx, apptoken)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["apptoken"] = apptoken
	response.With200V2(w, "Success", m, platform)
}

//UpdateApptoken :""
func (h *Handler) UpdateApptoken(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	apptoken := new(models.Apptoken)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&apptoken)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if apptoken.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateApptoken(ctx, apptoken)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableApptoken : ""
func (h *Handler) EnableApptoken(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableApptoken(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableApptoken : ""
func (h *Handler) DisableApptoken(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableApptoken(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteApptoken : ""
func (h *Handler) DeleteApptoken(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteApptoken(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleApptoken :""
func (h *Handler) GetSingleApptoken(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	apptoken, err := h.Service.GetSingleApptoken(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = apptoken
	response.With200V2(w, "Success", m, platform)
}

//FilterApptoken : ""
func (h *Handler) FilterApptoken(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var apptoken *models.ApptokenFilter

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	var pagination *models.Pagination
	if pageNo != "no" {
		pagination = new(models.Pagination)
		if pagination.PageNum = 1; pageNo != "" {
			page, err := strconv.Atoi(pageNo)
			if pagination.PageNum = 1; err == nil {
				pagination.PageNum = page
			}
		}
		if pagination.Limit = 10; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	err := json.NewDecoder(r.Body).Decode(&apptoken)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var apptokens []models.RefApptoken
	log.Println(pagination)
	apptokens, err = h.Service.FilterApptoken(ctx, apptoken, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(apptokens) > 0 {
		m["data"] = apptokens
	} else {
		res := make([]models.Apptoken, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// //SendAppTokenNotification : ""
// func (h *Handler) SendAppTokenNotification(w http.ResponseWriter, r *http.Request) {
// 	platform := r.URL.Query().Get("platform")
// 	apptoken := new(models.AppTokenNotification)
// 	var ctx *models.Context
// 	ctx = app.GetApp(r.Context(), h.Service.Daos)
// 	defer ctx.Client.Disconnect(r.Context())
// 	err := json.NewDecoder(r.Body).Decode(&apptoken)
// 	defer r.Body.Close()
// 	if err != nil {
// 		response.With400V2(w, err.Error(), platform)
// 		return
// 	}
// 	err = h.Service.SendAppTokenNotification(ctx, apptoken)
// 	if err != nil {
// 		response.With500mV2(w, "failed - "+err.Error(), platform)
// 		return
// 	}
// 	m := make(map[string]interface{})
// 	m["apptokenNotification"] = "success"
// 	response.With200V2(w, "Success", m, platform)
// }
func (h *Handler) LoginRegistration(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	apptoken := new(models.Apptoken)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&apptoken)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveApptoken(ctx, apptoken)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["apptoken"] = apptoken
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) LogoutApptoken(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.LogoutApptoken(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Apptoken"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) GetSingleApptokenWithUniqueCheck(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	apptoken, err := h.Service.GetSingleApptokenWithUniqueCheck(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]bool)

	if apptoken == nil {
		m["isApptoken"] = false
		response.With200V2bool(w, "Success", m, platform)

	} else {
		m["isApptoken"] = true
		response.With200V2bool(w, "Success", m, platform)

	}
}
