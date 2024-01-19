package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"municipalproduct1-service/app"
	"municipalproduct1-service/models"
	"municipalproduct1-service/response"
	"net/http"
	"strconv"
)

//SaveProperty : ""
func (h *Handler) SaveProperty(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	property := new(models.Property)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SaveProperty(ctx, property, "")
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyDemand(ctx, property.UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	err = h.Service.SaveOverAllPropertyDemandToProperty(ctx, property.UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	m["property"] = property
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) SavePropertyV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	property := new(models.Property)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.SavePropertyV2(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	// err = h.Service.SavePropertyDemand(ctx, property.UniqueID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	// err = h.Service.SaveOverAllPropertyDemandToProperty(ctx, property.UniqueID)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }

	m := make(map[string]interface{})
	m["property"] = property
	response.With200V2(w, "Success", m, platform)
}

//UpdateProperty :""
func (h *Handler) UpdateProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.Property)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if property.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdateProperty(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyDemand(ctx, property.UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	err = h.Service.SaveOverAllPropertyDemandToProperty(ctx, property.UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyGISTagging : ""
func (h *Handler) UpdatePropertyGISTagging(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	uniqueID := r.URL.Query().Get("id")

	pgt := new(models.PropertyGISTagging)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&pgt)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.UpdatePropertyGISTagging(ctx, uniqueID, pgt)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertygistagging"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// BasicUpdateProperty : ""
func (h *Handler) BasicUpdateProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	bpu := new(models.BasicPropertyUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&bpu)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	err = h.Service.BasicUpdateProperty(ctx, bpu)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basicpropertyupdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// AcceptBasicPropertyUpdate : ""
func (h *Handler) AcceptBasicPropertyUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.AcceptBasicPropertyUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.AcceptBasicPropertyUpdate(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basicpropertyupdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// RejectBasicPropertyUpdate : ""
func (h *Handler) RejectBasicPropertyUpdate(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.RejectBasicPropertyUpdate)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.RejectBasicPropertyUpdate(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["basicpropertyupdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//UpdateProperty :""
func (h *Handler) UpdatePropertyPreviousYrCollection(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	ppyc := new(models.PropertyPreviousYrCollection)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&ppyc)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if ppyc.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyPreviousYrCollection(ctx, ppyc)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err = h.Service.SavePropertyDemand(ctx, ppyc.UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	err = h.Service.SaveOverAllPropertyDemandToProperty(ctx, ppyc.UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["ppyc"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//EnableProperty : ""
func (h *Handler) EnableProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.EnableProperty(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DisableProperty : ""
func (h *Handler) DisableProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	code := r.URL.Query().Get("id")

	if code == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableProperty(ctx, code)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//DeleteProperty : ""
func (h *Handler) DeleteProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DeleteProperty(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetSingleProperty :""
func (h *Handler) GetSingleProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")
	resType := r.URL.Query().Get("resType")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	property := new(models.RefProperty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	property, err := h.Service.GetSingleProperty(ctx, UniqueID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	if property.PlsContactAdministrator == true {
		response.With500mV2(w, "failed - Please Contact Administrator", platform)
		return
	}
	if resType == "pdf" {
		data, err := h.Service.GetBasicPropertyDetailsPDF(ctx, UniqueID)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=property.pdf")
		return
	}
	m := make(map[string]interface{})
	m["property"] = property
	response.With200V2(w, "Success", m, platform)
}

//FilterProperty : ""
func (h *Handler) FilterProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var property *models.PropertyFilter
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
					pagination.Limit = 100
				}
			}

		}
	}
	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.FilterPropertyExcel(ctx, property, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	if resType == "excel2" {
		file, err := h.Service.FilterPropertyExcelV2(ctx, property, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	if resType == "excel3" {
		file, err := h.Service.FilterPropertyExcelV3(ctx, property, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	// header for pdf file
	if resType == "pdf" {
		data, err := h.Service.FilterPropertyPdf(ctx, property, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
		w.Write(data)

	}

	var propertys []models.RefProperty
	log.Println(pagination)
	propertys, err = h.Service.FilterProperty(ctx, property, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(propertys) > 0 {
		m["property"] = propertys
	} else {
		res := make([]models.Property, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

//FilterBasicPropertyUpdate : ""
func (h *Handler) FilterBasicPropertyUpdate(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	var property *models.FilterBasicPropertyUpdate
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
	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	var propertys []models.RefBasicPropertyUpdateLog
	log.Println(pagination)
	propertys, err = h.Service.FilterBasicPropertyUpdate(ctx, property, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(propertys) > 0 {
		m["property"] = propertys
	} else {
		res := make([]models.Property, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// PropertyWiseDemandandCollectionExcel : ""
func (h *Handler) PropertyWiseDemandandCollectionExcel(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var filter *models.PropertyFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	// setting pagination
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
		file, err := h.Service.PropertyWiseDemandandCollectionExcel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywisedemandandcollection.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	if resType == "excelv2" {
		file, err := h.Service.PropertyWiseDemandandCollectionV2Excel(ctx, filter, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertywisedemandandcollectionv2.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	// var wardwiseDemand []models.ResPropertyWiseDemandandCollectionV2Report
	// log.Println(pagination)
	// wardwiseDemand, err = h.Service.PropertyWiseDemandandCollectionV2JSON(ctx, filter, pagination)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	// m := make(map[string]interface{})

	// if len(wardwiseDemand) > 0 {
	// 	m["propertyDC"] = wardwiseDemand
	// } else {
	// 	res := make([]models.ResPropertyWiseDemandandCollectionV2Report, 0)
	// 	m["propertyDC"] = res
	// }
	// if pagination != nil {
	// 	if pagination.PageNum > 0 {
	// 		m["pagination"] = pagination
	// 	}
	// }

	// response.With200V2(w, "Success", m, platform)

}

// WardwiseDemand : ""
func (h *Handler) WardwiseDemand(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var pwdf *models.PropertyWardwiseDemandFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	// setting pagination
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

	err := json.NewDecoder(r.Body).Decode(&pwdf)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.WardwiseDemandExcel(ctx, pwdf, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=wardwisedemand.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var wardwiseDemand []models.WardwiseDemandandCollection
	log.Println(pagination)
	wardwiseDemand, err = h.Service.WardwiseDemand(ctx, pwdf, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(wardwiseDemand) > 0 {
		m["wardwiseDemand"] = wardwiseDemand
	} else {
		res := make([]models.Property, 0)
		m["wardwiseDemand"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// WardwiseCollection : ""
func (h *Handler) WardwiseCollection(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var pwdf *models.PropertyWardwiseDemandFilter
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")
	// setting pagination
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

	err := json.NewDecoder(r.Body).Decode(&pwdf)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.WardwiseCollectionExcel(ctx, pwdf, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

	var wardwiseCollection []models.WardwiseDemandandCollection
	log.Println(pagination)
	wardwiseCollection, err = h.Service.WardwiseCollection(ctx, pwdf, pagination)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})

	if len(wardwiseCollection) > 0 {
		m["wardwiseDemand"] = wardwiseCollection
	} else {
		res := make([]models.Property, 0)
		m["wardwiseDemand"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

// PropertyDemandExcel : ""
func (h *Handler) PropertyDemand(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var property *models.PropertyFilter
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
	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.PropertyDemandExcel(ctx, property, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}

}

//ActivateProperty : ""
func (h *Handler) ActivateProperty(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.ActivateProperty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.ActivateProperty(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//RejectProperty : ""
func (h *Handler) RejectProperty(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	req := new(models.RejectProperty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	err := json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	err = h.Service.RejectProperty(ctx, req)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetPropertyDemandCalc : ""
func (h *Handler) GetPropertyDemandCalc(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID
	const (
		CONSTANTSMS         = "SMS"
		CONSTANTEMAIL       = "EMAIL"
		CONSTANTSMSANDEMAIL = "SMSEMAIL"
		CONSTANTWHATSAPP    = "WHATSAPP"
	)
	fmt.Println("notifyType", resType)

	if resType == CONSTANTSMS || resType == CONSTANTEMAIL || resType == CONSTANTSMSANDEMAIL {
		err := h.Service.GetPropertyDemandCalcNotify(ctx, filter, resType)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		return
	}
	if resType == CONSTANTWHATSAPP {
		res, err := h.Service.GetPropertyDemandCalcNotifyV2(ctx, filter, resType)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)

			return
		}
		m := make(map[string]interface{})
		m["msg"] = res
		response.With200V2(w, "Success", m, platform)
		return
	}

	propertyDemand, err := h.Service.GetPropertyDemandCalc(ctx, filter, "")
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	updateDemand := new(models.UpdatePropertyTotalDemand)
	updateDemand.PropertyID = propertyDemand.Property.UniqueID
	updateDemand.TotalAmount = propertyDemand.TotalTax

	if err := h.Service.UpdatePropertyTotalDemand(ctx, updateDemand); err != nil {
		fmt.Println("ERR IN UPDATING PROPERTY DEMAND IN PROPERTY COLLECTION - " + err.Error())
	}

	propertyDemand.PropertyID = filter.PropertyID
	propertyDemand.OverallPropertyDemand.PropertyID = filter.PropertyID
	//demand.OverallPropertyDemand
	if err := h.Service.UpdateOverallPropertyDemand(ctx, &propertyDemand.OverallPropertyDemand); err != nil {
		fmt.Println("ERR IN UPDATING OVERALL PROPERTY DEMAND - " + err.Error())
	}
	m := make(map[string]interface{})
	m["propertyDemand"] = propertyDemand
	response.With200V2(w, "Success", m, platform)
}

//GetPropertyDemandCalcV2
func (h *Handler) GetPropertyDemandCalcV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID
	propertyDemand, err := h.Service.GetPropertyDemandCalcV2(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyDemand"] = propertyDemand
	response.With200V2(w, "Success", m, platform)
}

//GetPropertyDemandCalcForFYs : ""
func (h *Handler) GetPropertyDemandCalcForFYs(w http.ResponseWriter, r *http.Request) {
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
	propertyDemand, err := h.Service.GetPropertyDemandCalc(ctx, filter, "")
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyDemand"] = propertyDemand
	response.With200V2(w, "Success", m, platform)
}

//DemandCalc : ""
func (h *Handler) DemandCalc(w http.ResponseWriter, r *http.Request) {
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
	propertyDemand, err := h.Service.DemandCalc(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyDemand"] = propertyDemand
	response.With200V2(w, "Success", m, platform)
}

//DashboardPropertyStatus :""
func (h *Handler) DashboardPropertyStatus(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	filter := new(models.DashboardPropertyStatusFilter)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	data, err := h.Service.DashboardPropertyStatus(ctx, filter)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["dashboard"] = data
	response.With200V2(w, "Success", m, platform)
}

//GetMultiplePropertyDemandCalc : ""
func (h *Handler) GetMultiplePropertyDemandCalc(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

	pageNo := r.URL.Query().Get("pageno")
	Limit := r.URL.Query().Get("limit")

	filter := new(models.PropertyDemandFilter)
	err := json.NewDecoder(r.Body).Decode(&filter)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
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
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	if resType == "excel" {
		// f, err := h.Service.GetMultiplePropertyDemandCalcExcel(ctx, filter, pagination)
		// if err != nil {
		// 	response.With500mV2(w, err.Error(), platform)
		// 	return
		// }
		// w.Header().Set("Content-Type", "application/octet-stream")
		// w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		// w.Header().Set("Content-Transfer-Encoding", "binary")
		// f.Write(w)
		// return
	}
	fmt.Println("before h.Service.GetMultiplePropertyDemandCalc is working")
	propertyDemand, err := h.Service.GetMultiplePropertyDemandCalc(ctx, filter, pagination)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	if len(propertyDemand) > 0 {
		m["propertyDemand"] = propertyDemand
	} else {
		res := make([]models.PropertyDemand, 0)
		m["propertyDemand"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}
	response.With200V2(w, "Success", m, platform)
}

//GetPropertyDemandCalcPDF : ""
func (h *Handler) GetPropertyDemandCalcPDF(w http.ResponseWriter, r *http.Request) {
	// defer func() {
	// 	r.Header.Set("Connection", "close")
	// 	r.Close = true
	// }()
	r.Body.Close()
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID
	data, err := h.Service.GetPropertyDemandCalcPDF(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	w.Write(data)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=demandreceipt.pdf")

}

//GetPropertyDemandCalcPDF : ""
func (h *Handler) GetPropertyDemandCalcPDFV2(w http.ResponseWriter, r *http.Request) {
	// defer func() {
	// 	r.Header.Set("Connection", "close")
	// 	r.Close = true
	// }()
	r.Body.Close()
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	v := r.URL.Query().Get("v")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID
	resPD, err := h.Service.GetSingleDefaultProductConfiguration(ctx)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	if v == "legacy" {
		data, err := h.Service.GetPropertyDemandCalcPDF(ctx, filter)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				response.With500mV2(w, "failed no data for this id", platform)
				return
			}
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}

		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=demandreceipt.pdf")
		return
	}
	if resPD.OldDemandReceipt == "Yes" {
		data, err := h.Service.GetPropertyDemandCalcPDF(ctx, filter)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				response.With500mV2(w, "failed no data for this id", platform)
				return
			}
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}

		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=demandreceipt.pdf")
		return
	} else {
		data, err := h.Service.GetPropertyDemandCalcPDFV2(ctx, filter)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				response.With500mV2(w, "failed no data for this id", platform)
				return
			}
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}

		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=demandreceipt.pdf")
	}
}

//
func (h *Handler) GetPaymentReceiptsPDF(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID

	data, err := h.Service.GetPaymentReceiptsPDF(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
	w.Write(data)

}

// GetPaymentReceiptsPDFV2 : ""
func (h *Handler) GetPaymentReceiptsPDFV2(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ID := r.URL.Query().Get("id")
	v := r.URL.Query().Get("v")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID
	if v == "legacy" {
		data, err := h.Service.GetPaymentReceiptsPDF(ctx, ID)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				response.With500mV2(w, "failed no data for this id", platform)
				return
			}
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Write(data)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
		return
	}
	data, err := h.Service.GetPaymentReceiptsPDFV2(ctx, ID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	w.Write(data)
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
}

//PropertyParkPenaltyEnable : ""
func (h *Handler) PropertyParkPenaltyEnable(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.PropertyParkPenaltyEnable(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//PropertyParkPenaltyDisable : ""
func (h *Handler) PropertyParkPenaltyDisable(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.PropertyParkPenaltyDisable(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyLocation : ""
func (h *Handler) UpdatePropertyLocation(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.PropertyLocation)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if property.PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyLocation(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyLocationUpdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyPicture : ""
func (h *Handler) UpdatePropertyPicture(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.PropertyPicture)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if property.PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.UpdatePropertyPicture(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyPictureUpdate"] = "success"
	response.With200V2(w, "Success", m, platform)
}

// GETPaymentReceiptsPDFFilesaved : ""
func (h *Handler) GETPaymentReceiptsPDFFilesaved(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")

	var IDs []string
	err := json.NewDecoder(r.Body).Decode(&IDs)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err = h.Service.GetPaymentReceiptsPDfServiceLOOP(ctx, IDs)
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

// SavePaymentReceiptsPDFFilesaved : ""
func (h *Handler) SavePaymentReceiptsPDFFilesaved(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")

	var IDs []string
	err := json.NewDecoder(r.Body).Decode(&IDs)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err = h.Service.SavePaymentReceiptsPDfServiceLOOP(ctx, IDs)
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

//GetPropertyDemandCalc : ""
func (h *Handler) SaveStoredDemand(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	//resType := r.URL.Query().Get("resType")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID

	propertyDemand, err := h.Service.SaveStoredDemand(ctx, filter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	propertyDemand.PropertyID = filter.PropertyID
	propertyDemand.OverallPropertyDemand.PropertyID = filter.PropertyID
	//demand.OverallPropertyDemand
	if err := h.Service.UpdateOverallPropertyDemand(ctx, &propertyDemand.OverallPropertyDemand); err != nil {
		fmt.Println("ERR IN UPDATING OVERALL PROPERTY DEMAND - " + err.Error())
	}
	m := make(map[string]interface{})
	m["propertyDemand"] = propertyDemand
	response.With200V2(w, "Success", m, platform)
}

//GetSingleProperty :""
func (h *Handler) GetDemandV3(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	UniqueID := r.URL.Query().Get("id")

	if UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
		return
	}

	//demand := new(models.DemandV3)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	demand, err := h.Service.GetDemandV3(ctx, UniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["demand"] = demand
	response.With200V2(w, "Success", m, platform)
}

//PropertyUpdateLocationReport : ""
func (h *Handler) PropertyUpdateLocationReport(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	var property *models.PropertyUpdateLocationFilter
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
	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}

	// header for excel file
	if resType == "excel" {
		file, err := h.Service.PropertyUpdateLocationExcelReport(ctx, property, pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=propertyupdatelocationexcelreport.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	// // header for pdf file
	// if resType == "pdf" {
	// 	data, err := h.Service.FilterPropertyPdf(ctx, property, pagination)
	// 	if err != nil {
	// 		response.With500mV2(w, "failed - "+err.Error(), platform)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/pdf")
	// 	w.Header().Set("Content-Disposition", "attachment; filename=paymentreceipt.pdf")
	// 	w.Write(data)

	// }

	var propertys []models.RefProperty
	log.Println(pagination)
	// propertys, err = h.Service.FilterProperty(ctx, property, pagination)
	// if err != nil {
	// 	response.With500mV2(w, "failed - "+err.Error(), platform)
	// 	return
	// }
	m := make(map[string]interface{})

	if len(propertys) > 0 {
		m["property"] = propertys
	} else {
		res := make([]models.Property, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}

	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) EnableHoldingProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	property := new(models.Property)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&property)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if property.UniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.EnableHoldingProperty(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = property
	response.With200V2(w, "Success", m, platform)
}

//DisableProperty : ""
func (h *Handler) DisableHoldingProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	uniqueID := r.URL.Query().Get("id")

	if uniqueID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := h.Service.DisableHoldingProperty(ctx, uniqueID)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["property"] = "success"
	response.With200V2(w, "Success", m, platform)
}

//GetPropertyDemandCalcWithStoredCalc
func (h *Handler) GetPropertyDemandCalcWithStoredCalc(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")
	ID := r.URL.Query().Get("id")
	if ID == "" {
		response.With400V2(w, "id is missing", platform)
	}

	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	filter := new(models.PropertyDemandFilter)
	filter.PropertyID = ID

	fmt.Println("notifyType", resType)

	propertyDemand, err := h.Service.GetPropertyDemandCalcWithStoredCalc(ctx, filter, "")
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	// propertyDemand.PropertyID = filter.PropertyID
	// propertyDemand.OverallPropertyDemand.PropertyID = filter.PropertyID
	// //demand.OverallPropertyDemand
	// if err := h.Service.UpdateOverallPropertyDemand(ctx, &propertyDemand.OverallPropertyDemand); err != nil {
	// 	fmt.Println("ERR IN UPDATING OVERALL PROPERTY DEMAND - " + err.Error())
	// }
	m := make(map[string]interface{})
	m["propertyDemand"] = propertyDemand
	response.With200V2(w, "Success", m, platform)
}

//GetPropertyDemandCalc : ""
func (h *Handler) GetAllPropertyDemandCalc(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	resType := r.URL.Query().Get("resType")

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
		if pagination.Limit = 50; Limit != "" {
			limit, err := strconv.Atoi(Limit)
			if pagination.Limit = 10; err == nil {
				pagination.Limit = limit
			}
		}
	}
	//	var filter *models.PropertyDemandFilter
	filter := new(models.PropertyDemandFilter)
	const (
		CONSTANTSMS         = "SMS"
		CONSTANTEMAIL       = "EMAIL"
		CONSTANTSMSANDEMAIL = "SMSEMAIL"
		CONSTANTWHATSAPP    = "WHATSAPP"
	)
	filter.Status = append(filter.Status, "Active")
	fmt.Println("notifyType", resType)
	if resType == "excel" {
		file, err := h.Service.GetAllPropertyDemandCalcReportExcel(ctx, filter, "", pagination)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=report.xlsx")
		w.Header().Set("ocntent-Transfer-Encoding", "binary")
		file.Write(w)
		return
	}
	propertyDemand, err := h.Service.GetAllPropertyDemandCalc(ctx, filter, "", pagination)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}

	m := make(map[string]interface{})
	if len(propertyDemand) > 0 {
		m["property"] = propertyDemand
	} else {
		res := make([]models.Property, 0)
		m["property"] = res
	}
	if pagination != nil {
		if pagination.PageNum > 0 {
			m["pagination"] = pagination
		}
	}
	response.With200V2(w, "Success", m, platform)
}

// CheckWardWisesOldHoldingNoOfProperty :""
func (h *Handler) CheckWardWisesOldHoldingNoOfProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")
	ward := r.URL.Query().Get("ward")
	oldHoldingNo := r.URL.Query().Get("oldNo")

	if ward == "" {
		response.With400V2(w, "wardNo is missing", platform)
	}
	if oldHoldingNo == "" {
		response.With400V2(w, "oldholdingNo is missing", platform)
	}

	// property := new(models.RefProperty)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	_, err := h.Service.CheckWardWiseOldHoldingNoOfProperty(ctx, ward, oldHoldingNo)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			response.With500mV2(w, "failed no data for this id", platform)
			return
		}
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["oldHoldingNo"] = "success"
	// m["response"] = property
	response.With200V2(w, "Success", m, platform)
}

// UpdatePropertyUniqueID : ""
func (h *Handler) UpdatePropertyUniqueID(w http.ResponseWriter, r *http.Request) {

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
	err = h.Service.UpdatePropertyUniqueID(ctx, property)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["updatePropertyUniqueId"] = "success"
	response.With200V2(w, "Success", m, platform)
}

func (h *Handler) CreateUserChargeForProperty(w http.ResponseWriter, r *http.Request) {

	platform := r.URL.Query().Get("platform")

	propertyusercharge := new(models.PropertyUserCharge)
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())

	err := json.NewDecoder(r.Body).Decode(&propertyusercharge)
	defer r.Body.Close()
	if err != nil {
		response.With400V2(w, err.Error(), platform)
		return
	}
	if propertyusercharge.PropertyID == "" {
		response.With400V2(w, "id is missing", platform)
	}
	err = h.Service.CreateUserChargeForProperty(ctx, propertyusercharge)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["propertyusercharge"] = "success"
	response.With200V2(w, "Success", m, platform)
}
