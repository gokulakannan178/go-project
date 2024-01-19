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

//SaveFarmer : ""
func (h *Handler) SaveFarmer(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Farmer := new(models.Farmer)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveFarmer(ctx, Farmer)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["Farmer"] = Farmer
	response.With200V2(w, "Success", m, platform)
}

//UpdateFarmer :""
func (h *Handler) UpdateFarmer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	Farmer := new(models.Farmer)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&Farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if Farmer.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateFarmer(ctx, Farmer)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableFarmer : ""
func (h *Handler) EnableFarmer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableFarmer(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableFarmer : ""
func (h *Handler) DisableFarmer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableFarmer(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteFarmer : ""
func (h *Handler) DeleteFarmer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteFarmer(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleFarmer :""
func (h *Handler) GetSingleFarmer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	Farmer := new(models.RefFarmer)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	Farmer, err := h.Service.GetSingleFarmer(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = Farmer
	response.With200V2(w, "Success", m, platform)
}

//FilterFarmer : ""
func (h *Handler) FilterFarmer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	var Farmer *models.FarmerFilter
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
	err := json.NewDecoder(r.Body).Decode(&Farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var Farmers []models.RefFarmer
	log.Println(pagination)
	if resType == "excel" {
		file, err := h.Service.FarmerExcel(ctx, Farmer, nil)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=FarmerReport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	if resType == "reportexcel" {
		file, err := h.Service.FarmerReportExcel(ctx, Farmer, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=FarmerReport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	Farmers, err = h.Service.FilterFarmer(ctx, Farmer, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Farmers) > 0 {
		m["data"] = Farmers
	} else {
		res := make([]models.Farmer, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//FilterFarmerBasic : ""
func (h *Handler) FilterFarmerBasic(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var Farmer *models.FarmerFilter
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
	err := json.NewDecoder(r.Body).Decode(&Farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	log.Println(pagination)

	Farmers, err := h.Service.FilterFarmerBasic(ctx, Farmer, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Farmers) > 0 {
		m["data"] = Farmers
	} else {
		res := make([]models.Farmer, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) FarmerUniquenessCheckRegistration(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	OrgID := r.URL.Query().Get("farmerOrg")
	Param := r.URL.Query().Get("param")
	Value := r.URL.Query().Get("value")
	if OrgID == "" || Param == "" || Value == "" {
		response.With400V2(w, "orgId/Param/Value is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	pass, err := h.Service.FarmerUniquenessCheckRegistration(ctx, OrgID, Param, Value)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["token"] = pass
	response.With200V2(w, "Success", m, platform)
}

//GetSingleFarmerWithMobilenoAndOrg :""
func (h *Handler) GetSingleFarmerWithMobilenoAndOrg(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("mobileNumber")
	org := r.URL.Query().Get("farmerOrg")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	if org == "" {
		response.With400V2(w, "organisation is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	_, err := h.Service.GetSingleFarmerWithMobilenoAndOrg(ctx, org, UniqueID)
	if err != nil {
		if err.Error() == "farmer not found" {
			response.With201mV2(w, "Success", platform)
			return
		}
		//	response.With201mV2(w, "Success", platform)
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "duplicate user"
	response.With500dV2(w, "duplicate user", m, platform)
}
func (h *Handler) GenerateotpFarmerRegistration(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	Farmer := new(models.Farmer)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&Farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.GenerateotpFarmerRegistration(ctx, Farmer)
	if err != nil {
		if err.Error() == "Farmer Already Registered" {
			m := make(map[string]interface{})
			m["Farmer"] = "Farmer Already Registered"
			response.With200V2(w, "Success", m, platform)
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
func (h *Handler) FarmerNearBy(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	farmernb := new(models.NearBy)

	ctx := app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&farmernb)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	//km!=0
	if farmernb.KM == 0 {
		response.With400V2(w, "id is missing", platform)
	}
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
	farmers, err := h.Service.FarmerNearBy(ctx, farmernb, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["farmerNearby"] = farmers
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//OTPLoginValidateOTP : "Login user using OTP"
func (h *Handler) RegistrationValidateOTPFarmer(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	farmer := new(models.FarmerOTPLogin)
	err := json.NewDecoder(r.Body).Decode(&farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, constants.RESPONSEINVALIDDATA+" "+err.Error(), platform)
		return
	}

	err = h.Service.RegistrationValidateOTPFarmer(ctx, farmer)

	if err != nil {
		if err.Error() == "Farmer Already Registered" {
			response.With403mV2(w, "Farmer already Registered", platform)
			return
		}
		response.With500mV2(w, err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["farmer"] = farmer
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) AddProjectFarmer(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var Farmer *models.FarmerFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&Farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var Farmers []models.AddProjectFarmer

	Farmers, err = h.Service.AddProjectFarmer(ctx, Farmer)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Farmers) > 0 {
		m["farmer"] = Farmers
	} else {
		res := make([]models.Farmer, 0)
		m["farmer"] = res
	}
	response.With200V2(w, "Success", m, platform)
}

//UpdateFarmerProfileImage :""
func (h *Handler) UpdateFarmerProfileImage(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	Farmer := new(models.Farmer)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&Farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if Farmer.ID.IsZero() {
		response.With400V2(w, "id is missing", platform)
		return
	}
	err = h.Service.UpdateFarmerProfileImage(ctx, Farmer)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}
func (h *Handler) FilterFarmerWithLocation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	//resType := r.URL.Query().Get("resType")

	var Farmer *models.FarmerFilter
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
	err := json.NewDecoder(r.Body).Decode(&Farmer)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var Farmers []models.FarmerLocation
	log.Println(pagination)
	Farmers, err = h.Service.FilterFarmerWithLocation(ctx, Farmer, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(Farmers) > 0 {
		m["data"] = Farmers
	} else {
		res := make([]models.Farmer, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}
