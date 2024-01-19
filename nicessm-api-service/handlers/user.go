package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
	"strconv"
)

//SaveUser : ""
func (h *Handler) SaveUser(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	user := new(models.User)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveUserwithtransaction(ctx, user)
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
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if user.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
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
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableUser(ctx, UniqueID)
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

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableUser(ctx, code)
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

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteUser(ctx, UniqueID)
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
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	user := new(models.RefUser)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	user, err := h.Service.GetSingleUser(ctx, UniqueID)
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

//GetMobileValidation :""
func (h *Handler) GetMobileValidation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	Mobile := r.URL.Query().Get("id")

	if Mobile == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.GetMobileValidation(ctx, Mobile)
	if err != nil {

		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = ""
	response.With200V2(w, "Success", m, platform)
}

//FilterUser : ""
func (h *Handler) FilterUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var user *models.UserFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
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
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var users []models.RefUser
	log.Println(pagination)
	resType := r.URL.Query().Get("resType")
	if resType == "excel" {
		file, err := h.Service.UserReportExcelV2(ctx, user, nil)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=UserReport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
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
	code := r.URL.Query().Get("id")
	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.ResetUserPassword(ctx, code)
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
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&cp)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if cp.UserName == "" {
		response.With400V2(w, "id is missing", platform)
	}
	ok, msg, err := h.Service.ChangePassword(ctx, cp)
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
	id := r.URL.Query().Get("id")
	if id == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.ForgetPasswordGenerateOTP(ctx, id)
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
	response.With200V2(w, "Success", m, platform)
}

//ForgetPasswordValidateOTP : ""
func (h *Handler) ForgetPasswordValidateOTP(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	otp := r.URL.Query().Get("otp")
	if otp == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	pass, err := h.Service.ForgetPasswordValidateOTP(ctx, UniqueID, otp)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["token"] = pass
	response.With200V2(w, "Success", m, platform)
}

//PasswordUpdate :""
func (h *Handler) PasswordUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	user := new(models.RefPassword)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if user.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.PasswordUpdate(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//UserCollection :""
func (h *Handler) UserCollectionLimit(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	username := r.URL.Query().Get("userName")
	collectionLimit := new(models.CollectionLimit)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&collectionLimit)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if username == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UserCollectionLimit(ctx, username, collectionLimit)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["collectionLimit"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//UpdateOrganisation :""
func (h *Handler) UpdateUserTypeV2(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	user := new(models.User)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if user.ID.String() == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateUserTypeV2(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//UpdateUserPassword :""
func (h *Handler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	user := new(models.User)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if user.ID.String() == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateUserPassword(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) UserUniquenessCheckRegistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	OrgID := r.URL.Query().Get("userOrg")
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

//ApprovedUser :""
func (h *Handler) ApprovedUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	user := new(models.User)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if user.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.ApprovedUser(ctx, user)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//RejectUser : ""
func (h *Handler) RejectUser(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.RejectUser(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) GenerateOtpRegistrationUser(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	user := new(models.User)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.GenerateotpUserRegistration(ctx, user)
	if err != nil {
		if err.Error() == "user Already Registered" {
			response.With403mV2(w, "user already Registered", platform)
			return
		}
		response.With500mV2(w, err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	//if err == nil {
	m["otp"] = "Otp Sent Succesfully"
	//	}
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) RegistrationValidateOTPUser(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	user := new(models.UserOTPLogin)
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}

	err = h.Service.RegistrationValidateOTPUser(ctx, user)

	if err != nil {
		if err.Error() == "user Already Registered" {
			response.With403mV2(w, "user already Registered", platform)
			return
		}
		response.With500mV2(w, err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["user"] = user
	response.With200V2(w, "Success", m, platform)
}

//GetSingleUser :""
func (h *Handler) CheckuniqcknessUserWithUserName(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("userName")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	user := new(models.RefUser)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	user, err := h.Service.GetSingleUserWithUserName(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if user != nil {
		response.With409mV2(w, "username is already registered", platform)
		return
	}
	m := make(map[string]interface{})
	m["user"] = "username not registered"
	response.With200V2(w, "Success", m, platform)
}
