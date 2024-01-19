package handlers

import (
	"encoding/json"
	"lgf-ccc-service/app"
	"lgf-ccc-service/models"
	"lgf-ccc-service/response"
	"log"
	"net/http"
	"strconv"
)

//SaveUser : ""
func (h *Handler) SaveUser(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	user := new(models.User)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveUser(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = user
	response.With200V2(w, "Success", m, platform)
}

//UpdateUser :""
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	user := new(models.User)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if user.UniqueID == "" {
		response.With400V2(w, "userName is missing", platform)
	}
	err = h.Service.UpdateUser(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableUser : ""
func (h *Handler) EnableUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UserName := r.URL.Query().Get("id")

	if UserName == "" {
		response.With400V2(w, "userName is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.EnableUser(ctx, UserName)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableUser : ""
func (h *Handler) DisableUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UserName := r.URL.Query().Get("id")

	if UserName == "" {
		response.With400V2(w, "userName is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DisableUser(ctx, UserName)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteUser : ""
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UserName := r.URL.Query().Get("id")

	if UserName == "" {
		response.With400V2(w, "userName is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.DeleteUser(ctx, UserName)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleUser :""
func (h *Handler) GetSingleUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UserName := r.URL.Query().Get("id")

	if UserName == "" {
		response.With400V2(w, "userName is missing", platform)
	}

	user := new(models.RefUser)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	user, err := h.Service.GetSingleUser(ctx, UserName)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = user
	response.With200V2(w, "Success", m, platform)
}

//FilterUser : ""
func (h *Handler) FilterUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var user *models.UserFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
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
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var users []models.RefUser
	log.Println(pagination)
	users, err = h.Service.FilterUser(ctx, user, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(users) > 0 {
		m["user"] = users
	} else {
		res := make([]models.User, 0)
		m["user"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//ResetUserPassword : ""
func (h *Handler) ResetUserPassword(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	code := r.URL.Query().Get("userName")
	password := r.URL.Query().Get("password")
	if code == "" {
		response.With400V2(w, "userName is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := h.Service.ResetUserPassword(ctx, code, password)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//ChangePassword : ""
func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	cp := new(models.UserChangePassword)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&cp)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if cp.UserName == "" {
		response.With400V2(w, "userName is missing", platform)
	}
	ok, msg, err := h.Service.ChangePassword(ctx, cp)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if !ok {
		response.With500mV2(w, msg, platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = msg
	response.With200V2(w, "Success", m, platform)
}

//ForgetPasswordNewPassword : ""
func (h *Handler) ForgetPasswordNewPassword(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	np := new(models.UserNewPassword)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	err := json.NewDecoder(r.Body).Decode(&np)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if np.UserName == "" {
		response.With400V2(w, "UserName is missing", platform)
	}
	if np.Token == "" {
		response.With400V2(w, "token is missing", platform)
	}
	ok, msg, err := h.Service.ForgetPasswordNewPassword(ctx, np)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if !ok {
		response.With403mV2(w, msg, platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//ForgetPasswordGenerateOTP : ""
func (h *Handler) ForgetPasswordGenerateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	username := r.URL.Query().Get("username")
	if username == "" {
		response.With400V2(w, "username is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.ForgetPasswordGenerateOTP(ctx, username)
	if err != nil {
		if err.Error() == "user not fount" {
			response.With403mV2(w, "Check username", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return

	}
	refUser, err := h.Service.GetSingleUserWithUserName(ctx, username)
	if err != nil {
		if err.Error() == "user not fount" {
			response.With403mV2(w, "Check username", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return

	}
	m := make(map[string]interface{})
	m["user"] = "success"
	m["refuser"] = refUser
	response.With200V2(w, "Success", m, platform)
}

//ForgetPasswordValidateOTP : ""
func (h *Handler) ForgetPasswordValidateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UserName := r.URL.Query().Get("username")
	otp := r.URL.Query().Get("otp")
	if otp == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	pass, err := h.Service.ForgetPasswordValidateOTP(ctx, UserName, otp)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["token"] = pass
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UserUniquenessCheckRegistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	OrgID := r.URL.Query().Get("Org")
	Param := r.URL.Query().Get("param")
	Value := r.URL.Query().Get("value")
	if OrgID == "" || Param == "" || Value == "" {
		response.With400V2(w, "orgId/Param/Value is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	pass, err := h.Service.UserUniquenessCheckRegistration(ctx, OrgID, Param, Value)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["token"] = pass
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) AreaAssignForPM(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	VechileAssign := new(models.AreaAssign)
	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&VechileAssign)
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	defer r.Body.Close()

	err = h.Service.AreaAssignForPM(ctx, VechileAssign)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Vechile"] = VechileAssign
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GetSingleUserWithMobileNumber(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UserName := r.URL.Query().Get("id")

	if UserName == "" {
		response.With400V2(w, "Mobileno is missing", platform)
	}

	user := new(models.RefUser)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)

	user, err := h.Service.GetSingleUserWithMobileNumber(ctx, UserName)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = user
	response.With200V2(w, "Success", m, platform)
}
