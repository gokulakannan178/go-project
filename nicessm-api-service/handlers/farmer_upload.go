package handlers

import (
	"fmt"
	"net/http"
	"nicessm-api-service/app"
	"nicessm-api-service/models"
	"nicessm-api-service/response"
)

//FarmerLandUploadExcel : ""
func (h *Handler) FarmerLandUploadExcel(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	organisation := r.URL.Query().Get("org")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()
	if organisation == "MPKV" {
		farmerland := h.Service.FarmerLandUploadExcelV2(ctx, file, false)
		if err != nil {
			response.With500mV2(w, "failed - "+err.Error(), platform)
			return
		}
		m := make(map[string]interface{})
		fmt.Println("dddd", farmerland)
		m["FarmerAggregationUpload"] = farmerland
		response.With200V2(w, "Success", m, platform)
		return
	}
	err = h.Service.FarmerLandUploadExcel(ctx, file)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["FarmerLandUpload"] = "success"
	response.With200V2(w, "Success", m, platform)

}

// FarmerAggregationUploadExcel : ""
func (h *Handler) FarmerAggregationUploadExcel(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()

	farmer := h.Service.FarmerAggregationUploadExcel(ctx, file)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["FarmerAggregationUpload"] = farmer
	response.With200V2(w, "Success", m, platform)
}

// FarmerCasteUploadExcel : ""
func (h *Handler) FarmerCasteUploadExcel(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()

	uploaderrs, err := h.Service.FarmerCasteUploadExcel(ctx, file)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["FarmerCasteUpload"] = "success"
	m["uploadErros"] = uploaderrs
	response.With200V2(w, "Success", m, platform)
}

// FarmerSoilUploadExcel : ""
func (h *Handler) FarmerSoilUploadExcel(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()

	uploaderrs, err := h.Service.FarmerSoilUploadExcel(ctx, file)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["FarmerSoilUpload"] = "success"
	m["uploadErros"] = uploaderrs
	response.With200V2(w, "Success", m, platform)
}

// FarmerCropUploadExcel : ""
func (h *Handler) FarmerCropUploadExcel(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()

	uploaderrs, err := h.Service.FarmerCropUploadExcel(ctx, file)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["FarmerCropUpload"] = "success"
	m["uploadErros"] = uploaderrs
	response.With200V2(w, "Success", m, platform)
}

// FarmerAggregationUploadExcelWithNames : ""
func (h *Handler) FarmerAggregationUploadExcelWithNames(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	searchPolicy := r.URL.Query().Get("searchPolicy")
	//uploadVersion := r.URL.Query().Get("uploadVersion")
	version := r.URL.Query().Get("version")
	organisation := r.URL.Query().Get("org")

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()
	var farmer []models.FarmerUploadError
	if organisation != "" {
		if organisation == "JNKVV" {
			farmer = h.Service.FarmerAggregationUploadExcelWithNames(ctx, file, func() bool {
				if searchPolicy == "Yes" {
					return true
				}
				return false
			}(), version)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
		} else if organisation == "MPKV" {
			farmer = h.Service.FarmerAggregationUploadExcelWithNamesV2(ctx, file, func() bool {
				if searchPolicy == "Yes" {
					return true
				}
				return false
			}(), version)
			if err != nil {
				response.With500mV2(w, "failed - "+err.Error(), platform)
				return
			}
		}
	}

	m := make(map[string]interface{})
	m["FarmerAggregationUpload"] = farmer
	response.With200V2(w, "Success", m, platform)
}

// FarmerAggregationUploadExcelWithNames : ""
func (h *Handler) FarmerAggregationUploadExcelWithNamesV2(w http.ResponseWriter, r *http.Request) {
	platform := r.URL.Query().Get("platform")
	searchPolicy := r.URL.Query().Get("searchPolicy")
	//uploadVersion := r.URL.Query().Get("uploadVersion")
	version := r.URL.Query().Get("version")

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer file.Close()
	var ctx *models.Context
	ctx = app.GetApp(r.Context(), h.Service.Daos)
	defer ctx.Client.Disconnect(r.Context())
	defer r.Body.Close()

	farmer := h.Service.FarmerAggregationUploadExcelWithNamesV2(ctx, file, func() bool {
		if searchPolicy == "Yes" {
			return true
		}
		return false
	}(), version)
	if err != nil {
		response.With500mV2(w, "failed - "+err.Error(), platform)
		return
	}
	m := make(map[string]interface{})
	m["FarmerAggregationUpload"] = farmer
	response.With200V2(w, "Success", m, platform)
}
