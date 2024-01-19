package handlers

import (
	"encoding/json"
	"haritv2-service/app"
	"haritv2-service/constants"
	"haritv2-service/models"
	"haritv2-service/response"
	"log"
	"net/http"
	"strconv"
)

//SaveULB : ""
func (h *Handler) SaveULB(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ulb := new(models.ULB)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ulb)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveULB(ctx, ulb)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulb"] = ulb
	response.With200V2(w, "Success", m, platform)
}

//UpdateULB :""
func (h *Handler) UpdateULB(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	ulb := new(models.ULB)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&ulb)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if ulb.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	if ulb.NodalOfficer.MobileNo == "" {
		response.With400V2(w, "mobileno is missing", platform)
	}
	err = h.Service.UpdateULB(ctx, ulb)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableULB : ""
func (h *Handler) EnableULB(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableULB(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableULB : ""
func (h *Handler) DisableULB(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableULB(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteULB : ""
func (h *Handler) DeleteULB(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteULB(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleULB :""
func (h *Handler) GetSingleULB(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	ulb := new(models.RefULB)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	ulb, err := h.Service.GetSingleULB(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["data"] = ulb
	response.With200V2(w, "Success", m, platform)
}

//FilterULB : ""
func (h *Handler) FilterULB(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var ulb *models.ULBFilter
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
	err := json.NewDecoder(r.Body).Decode(&ulb)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ulbs []models.RefULB
	log.Println(pagination)
	if resType == "excel" {
		if ulb.ExcelType == constants.ULBEXCELTESTCERTPENDING {
			file, err := h.Service.UlbTestcertExcelForPending(ctx, ulb, nil)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Disposition", "attachment; filename=UlbtestcertPending.xlsx")
			w.Header().Set("ocntent-Transfer-Encoding", "binary")
			file.Write(w)
			return
		}
		if ulb.ExcelType == constants.ULBEXCELTESTCERTAPPROVED {
			file, err := h.Service.UlbTestcertExcelForApproved(ctx, ulb, nil)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Disposition", "attachment; filename=UlbtestcertApproved.xlsx")
			w.Header().Set("ocntent-Transfer-Encoding", "binary")
			file.Write(w)
			return
		}
		if ulb.ExcelType == constants.ULBEXCELTESTCERTREJECTED {
			file, err := h.Service.UlbTestcertExcelForRejected(ctx, ulb, nil)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Disposition", "attachment; filename=UlbtestcertRejected.xlsx")
			w.Header().Set("ocntent-Transfer-Encoding", "binary")
			file.Write(w)
			return
		}
		if ulb.ExcelType == constants.ULBEXCELTESTCERTREAPPLIED {
			file, err := h.Service.UlbTestcertExcelForReApplied(ctx, ulb, nil)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Disposition", "attachment; filename=UlbtestcertReApplied.xlsx")
			w.Header().Set("ocntent-Transfer-Encoding", "binary")
			file.Write(w)
			return
		}
		if ulb.ExcelType == constants.ULBEXCELTESTCERTEXPIRY {
			file, err := h.Service.UlbTestcertExcelForExpiry(ctx, ulb, nil)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Disposition", "attachment; filename=UlbtestcertExpiry.xlsx")
			w.Header().Set("ocntent-Transfer-Encoding", "binary")
			file.Write(w)
			return
		}

	}

	//var ulbs []models.RefULB
	log.Println(pagination)
	ulbs, err = h.Service.FilterULB(ctx, ulb, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(ulbs) > 0 {
		m["data"] = ulbs
	} else {
		res := make([]models.ULB, 0)
		m["data"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//AddULBTestCert : ""
func (h *Handler) AddULBTestCert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ulbestCert := new(models.ULBTestCert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ulbestCert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.AddULBTestCert(ctx, UniqueID, ulbestCert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbTestCert"] = ulbestCert
	response.With200V2(w, "Success", m, platform)

}

//ApplyForTestCert : ""
func (h *Handler) ApplyForTestCert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	ulbestCert := new(models.ULBTestCert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ulbestCert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.ApplyForTestCert(ctx, UniqueID, ulbestCert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbTestCert"] = ulbestCert
	response.With200V2(w, "Success", m, platform)

}
func (h *Handler) ReApplyForTestCert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	ulbestCert := new(models.ULBTestCert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ulbestCert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.ReApplyForTestCert(ctx, UniqueID, ulbestCert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbTestCert"] = ulbestCert
	response.With200V2(w, "Success", m, platform)

}

//AcceptTestCert : ""
func (h *Handler) AcceptTestCert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ulbTestCert := new(models.ULBTestCert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ulbTestCert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	// if ulbTestCert.UniqueID == "" {
	// 	response.With400V2(w, "No Unique Id", platform)
	// 	return
	// }
	err = h.Service.AcceptTestCert(ctx, UniqueID, ulbTestCert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbTestCert"] = ulbTestCert
	response.With200V2(w, "Success", m, platform)

}
func (h *Handler) RejectTestCert(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ulbTestCert := new(models.ULBTestCert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ulbTestCert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	// if ulbTestCert.UniqueID == "" {
	// 	response.With400V2(w, "No Unique Id", platform)
	// 	return
	// }
	err = h.Service.RejectTestCert(ctx, UniqueID, ulbTestCert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbTestCert"] = ulbTestCert
	response.With200V2(w, "Success", m, platform)
}

//ULBTestCertStatus : ""
func (h *Handler) ULBTestCertStatus(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ulbTestCert := new(models.ULBTestCert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ulbTestCert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	// if ulbTestCert.UniqueID == "" {
	// 	response.With400V2(w, "No Unique Id", platform)
	// 	return
	// }
	err = h.Service.AddULBTestCert(ctx, UniqueID, ulbTestCert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbTestCert"] = ulbTestCert
	response.With200V2(w, "Success", m, platform)

}

//GetSingleULB :""
func (h *Handler) GetULBTestCertStatus(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	ulbTestCert := new(models.ULBTestCert)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&ulbTestCert)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	// if ulbTestCert.UniqueID == "" {
	// 	response.With400V2(w, "No Unique Id", platform)
	// 	return
	// }
	err = h.Service.AddULBTestCert(ctx, UniqueID, ulbTestCert)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ulbTestCert"] = ulbTestCert
	response.With200V2(w, "Success", m, platform)

}
func (h *Handler) ULBMobileUniqueness(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var filter *models.ULB
	var mobile string
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&mobile)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.ULBMobileUniqueness(ctx, filter, mobile)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["mobile"] = mobile
	response.With200V2(w, "Success", m, platform)
}
