package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//InitiatePropertyPayment : ""
func (h *Handler) InitiatePropertyPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	filter := new(models.PropertyDemandFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	txnID, err := h.Service.InitiatePropertyPayment(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["txtnId"] = txnID
	response.With200V2(w, "Success", m, platform)
}

//GetSinglePropertyPaymentTxtID : ""
func (h *Handler) GetSinglePropertyPaymentTxtID(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.GetSinglePropertyPaymentTxtID(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyPayment"] = data
	response.With200V2(w, "Success", m, platform)
}

//PropertyMakePayment : ""
func (h *Handler) PropertyMakePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	payment := new(models.PropertyMakePayment)
	err := json.NewDecoder(r.Body).Decode(&payment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	propertyId, err := h.Service.PropertyMakePayment(ctx, payment)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if propertyId == "" {
		response.With500mV2(w, "failed to get property id- ", platform)
		return
	}
	//if payment.Details.MOP.Mode != constants.MOPCHEQUE {
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyDemand(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.PropertyUpdateCollection(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	//}
	m := make(map[string]interface{})
	m["payment"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetAllPaymentsForProperty : ""
func (h *Handler) GetAllPaymentsForProperty(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.GetAllPaymentsForProperty(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["payments"] = data
	response.With200V2(w, "Success", m, platform)
}

//DashboardTotalCollection :""
func (h *Handler) DashboardTotalCollection(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	filter := new(models.DashboardTotalCollectionFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	data, err := h.Service.DashboardTotalCollection(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dashboard"] = data
	response.With200V2(w, "Success", m, platform)
}

//FilterPropertyPayment : ""
func (h *Handler) FilterPropertyPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyPaymentFilter
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
				if limit == 100 {
					pagination.Limit = 10
				}
			}

		}
	}
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterPropertyPaymentExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertypayment.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	// header for pdf file
	if resType == "pdf" {
		data, err := h.Service.FilterPropertyPaymentPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
		w.Write(data)
		return

	}
	// header for pdfv2 file
	if resType == "pdfv2" {
		data, err := h.Service.FilterPropertyPaymentPDFV2(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=paymentreceiptv2.pdf")
		w.Write(data)
		return

	}
	// header for pdfv3 file
	if resType == "pdfv3" {
		data, err := h.Service.FilterPropertyPaymentPDFV3(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=paymentreceiptv3.pdf")
		w.Write(data)
		return

	}
	var payments []models.RefPropertyPayment
	log.Println(pagination)
	payments, err = h.Service.FilterPropertyPayment(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["payments"] = payments
	} else {
		res := make([]models.RefPropertyPayment, 0)
		m["payments"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// CheckReport : ""
func (h *Handler) ChequeReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	reportType := r.URL.Query().Get("reportType")
	var filter *models.PropertyPaymentFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		var file *excelize.File
		var fileName string
		var err error
		switch reportType {
		case "bounce":
			fileName = "checkbouncereport"
			file, err = h.Service.ChequeBounceReport(ctx, filter, pagination)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
		case "pending":
			fileName = "checkpendingreport"
			file, err = h.Service.PendingChequeReport(ctx, filter, pagination)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
		case "verified":
			fileName = "checkverifiedreport"
			file, err = h.Service.VerifyChequeReport(ctx, filter, pagination)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}

		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName+".xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	// header for Pdf file
	if resType == "pdf" {
		var file []byte
		var fileName string
		var err error
		switch reportType {
		case "bounce":
			fileName = "checkpendingreport"
			file, err = h.Service.ChequeBounceReportPdf(ctx, filter, pagination)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
		case "pending":
			fileName = "checkpendingreport"
			file, err = h.Service.PendingChequeReportPdf(ctx, filter, pagination)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
		case "verified":
			fileName = "checkverifiedreport"
			file, err = h.Service.VerifyChequeReportPdf(ctx, filter, pagination)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
		}

		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName+".pdf")
		w.Write(file)

	}

}

// CounterReport : ""
func (h *Handler) CounterReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyPaymentFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.CounterReport(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=counterreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var payments []models.RefPropertyPayment
	log.Println(pagination)
	payments, err = h.Service.FilterPropertyPayment(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["payments"] = payments
	} else {
		res := make([]models.RefPropertyPayment, 0)
		m["payments"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//GenerateReciptForAPayment : ""
func (h *Handler) GenerateReciptForAPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	data, err := h.Service.GenerateReciptForAPayment(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyPayment"] = data
	response.With200V2(w, "Success", m, platform)
}

// VerifyPayment : ""
func (h *Handler) VerifyPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var vp models.VerifyPayment
	if err := json.NewDecoder(r.Body).Decode(&vp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	propertyId, err := h.Service.VerifyPayment(ctx, &vp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyDemand(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.PropertyUpdateCollection(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["verifiedProperty"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// NotVerifiedPayment : ""
func (h *Handler) NotVerifiedPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var vp models.NotVerifiedPayment
	if err := json.NewDecoder(r.Body).Decode(&vp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))

		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := h.Service.NotVerifiedPayment(ctx, &vp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}
	m := make(map[string]interface{})
	m["notVerifiedProperty"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// BouncePayment : ""
func (h *Handler) BouncePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var bp models.BouncePayment
	if err := json.NewDecoder(r.Body).Decode(&bp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	fmt.Println("tnx id - ", bp.TnxID)
	propertyId, err := h.Service.BouncePayment(ctx, &bp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyDemand(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.PropertyUpdateCollection(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["bouncedPayment"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// RejectPayment : ""
func (h *Handler) RejectPayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var rp models.RejectPayment
	if err := json.NewDecoder(r.Body).Decode(&rp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	fmt.Println("tnx id - ", rp.TnxID)
	propertyId, err := h.Service.RejectPayment(ctx, &rp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyDemand(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.PropertyUpdateCollection(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["verifiedProperty"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

//RejectPaymentByReceiptNo
func (h *Handler) RejectPaymentByReceiptNo(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	var rp models.RejectPayment
	if err := json.NewDecoder(r.Body).Decode(&rp); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Data:" + err.Error()))
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	fmt.Println("tnx id - ", rp.TnxID)
	propertyId, err := h.Service.RejectPaymentByReceiptNo(ctx, &rp)
	if err != nil {
		response.With500mV2(w, "failed no data for this id", platform)
		return
	}

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyDemand(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.PropertyUpdateCollection(ctx, propertyId)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["verifiedProperty"] = "Success"
	response.With200V2(w, "Success", m, platform)
}

// DateRangeWisePropertyPaymentReport : ""
func (h *Handler) DateRangeWisePropertyPaymentReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.DateWisePropertyPaymentReportFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if resType == "excel" {
		file, err := h.Service.DayWisePropertyPaymentExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=mobiletoweroveralldemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	report := new(models.RefDateWisePropertyPaymentReport)
	report, err = h.Service.DateRangeWisePropertyPaymentReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dateRangeWise"] = report
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) DateWisePropertyPaymentReport(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	filter := new(models.DateWisePropertyPaymentReportFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if resType == "excel" {
		file, err := h.Service.DayWisePropertyPaymentExcel(ctx, filter)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=mobiletoweroveralldemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	report := make([]models.DateWisePropertyPaymentReport, 0)
	report, err = h.Service.DateWisePropertyPaymentReport(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dateRangeWise"] = report
	response.With200V2(w, "Success", m, platform)
}

func decrypt(cipherText string, workingKey string) string {
	iv := []byte("\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f")
	decDigest := md5.Sum([]byte(workingKey))
	encryptedText, err := hex.DecodeString(cipherText)
	if err != nil {
		log.Println(err)

	}
	dec_cipher, err := aes.NewCipher(decDigest[:])
	if err != nil {
		log.Println(err)
	}
	decryptedText := make([]byte, len(encryptedText))
	dec := cipher.NewCBCDecrypter(dec_cipher, iv)
	dec.CryptBlocks(decryptedText, encryptedText)
	return string(decryptedText)
}

// GetFailedOnlinePayment : ""
func (h *Handler) GetFailedOnlinePayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	Params := r.URL.Query()
	fmt.Println(platform)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := make(map[string]interface{})
	for i, v := range Params {
		filter[i] = v
		fmt.Println(i, "value is", v)

	}
	fmt.Println(filter)
	encResp := r.FormValue("encResp")
	workingKey := "2775EA71B1D1F115FF568EC1BF6D4653"
	decryptedData := decrypt(encResp, workingKey)
	fmt.Println(decryptedData)
	// Split the string into key-value pairs
	pairs := strings.Split(decryptedData, "&")
	// Create a new map
	data := make(map[string]string)

	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		if len(parts) >= 2 {
			key := parts[0]
			value := parts[1]
			data[key] = value
			fmt.Println("key ====>", key)
			fmt.Println("value ====>", value)
		}
	}

	resPD, err := h.Service.GetSingleDefaultProductConfiguration(ctx)
	if err != nil {
		response.With500mV2(w, "error in getting product configuration - ", platform)
		return
	}
	fmt.Println("resPD ======>", resPD)

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	res := `<html>
	<body>
	

	<div>
    <table style="margin: auto;">
      <tr>
        <td><img src="` + resPD.APIURL + resPD.Logo + `"></td>
      </tr>
      <tr style="text-align: center; font-weight: bold; font-size: 24px;">
        <td>Payment Failed</td>
      </tr>
      <tr style="text-align: center; font-size: 20px;">
        <td><a href="` + resPD.UIURL + `#/consumerpropertyList">Go to Home Page</a></td>
      </tr>
    </table>
</div>

	</body>
	</html>`
	fmt.Fprintf(w, res)

}

// PutOnlinePayment : ""
func (h *Handler) PutOnlinePayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	Params := r.URL.Query()
	fmt.Println(platform)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := make(map[string]interface{})
	for i, v := range Params {
		filter[i] = v
		fmt.Println(i, "value is", v)

	}
	fmt.Println(filter)

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	res := `<html>
	<body>
	<h1>Payment Completed</h1>
	</body>
	</html>`
	fmt.Fprintf(w, res)
}

// PostOnlinePayment :""
func (h *Handler) PostOnlinePayment(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	onlinePayment := new(models.OnlinePayment)
	err := json.NewDecoder(r.Body).Decode(&onlinePayment)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	// filter := make(map[string]interface{})
	// for i, v := range onlinePayment {
	// 	filter[i] = v
	// 	fmt.Println(i, "value is", v)

	// }
	// fmt.Println(filter)

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	res := `<html>
	<body>
	<h1>Payment Completed</h1>
	</body>
	</html>`
	fmt.Fprintf(w, res)
}

func (h *Handler) CollectedPropertyPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var crr models.CollectionReceivedRequest
	err := json.NewDecoder(r.Body).Decode(&crr)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err = h.Service.CollectedPropertyPayment(ctx, &crr)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	response.With200V2(w, "Success", nil, platform)
}

func (h *Handler) RejectedPropertyPayment(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	var crr models.CollectionReceivedRequest
	err := json.NewDecoder(r.Body).Decode(&crr)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err = h.Service.RejectedPropertyPayment(ctx, &crr)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	response.With200V2(w, "Success", nil, platform)
}

// PropertyPaymentArrerAndCurrentCollection : ""
func (h *Handler) PropertyPaymentArrerAndCurrentCollection(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyPaymentFilter
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
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.PropertyPaymentArrerAndCurrentCollectionExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertypayment.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	// header for pdf file
	if resType == "pdf" {
		data, err := h.Service.FilterPropertyPaymentPDF(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
		w.Write(data)
		return

	}

	var payments []models.ArrerAndCurrentReport
	log.Println(pagination)
	payments, err = h.Service.PropertyPaymentArrerAndCurrentCollection(ctx, filter, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(payments) > 0 {
		m["payments"] = payments
	} else {
		res := make([]models.RefPropertyPayment, 0)
		m["payments"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyPaymentBasicPropertyID : ""
func (h *Handler) UpdatePropertyPaymentBasicPropertyID(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.UpdatePropertyUniqueID)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if len(property.UniqueIDs) == 0 {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyPaymentBasicPropertyID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["updatePropertyPaymentBasicPropertyID"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyPaymentFYPropertyID : ""
func (h *Handler) UpdatePropertyPaymentFYPropertyID(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.UpdatePropertyUniqueID)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if len(property.UniqueIDs) == 0 {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyPaymentFYPropertyID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["updatePropertyPaymentFYPropertyID"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyPaymentPropertyID : ""
func (h *Handler) UpdatePropertyPaymentPropertyID(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.UpdatePropertyUniqueID)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if len(property.UniqueIDs) == 0 {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyPaymentPropertyID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["updatePropertyPaymentPropertyID"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) PropertyPaymentSummaryUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("tnxId")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := h.Service.PropertyPaymentSummaryUpdate(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyPayment"] = "Success"
	response.With200V2(w, "Success", m, platform)
}
