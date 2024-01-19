package services

import (
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
)

//FilterBulkPrint :""
func (s *Service) BulkPrintGetDetailsForProperty(ctx *models.Context, filter *models.BulkPrintFilter) (BulkPrint []models.BulkPrintDetail, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.BulkPrintGetDetailsForProperty(ctx, filter)
}
func (s *Service) BulkPrintReceiptsRequestForProperty(ctx *models.Context, bprr *models.BulkPrintReceiptsRequest) ([]byte, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	payments := models.PDFDataV2Arr{}
	pdfDta := new(models.PDFDataV2)
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	for _, v := range bprr.TnxIds {

		data, err := s.GetSinglePropertyPaymentTxtID(ctx, v)
		if err != nil {
			return nil, err
		}
		m2 := make(map[string]interface{})

		m2["arrearYears"] = ""
		m2["currentYear"] = ""
		if data != nil {
			var arrearYearTax, arrearYearRebate, arrearYearPenalty, arrearYearTotal, arrearYearOtherDemand float64

			if len(data.Fys) > 0 {
				if data.Fys[0].FY.IsCurrent {
					m2["currentYear"] = data.Fys[0].FY.Name
					m2["currentYearTax"] = data.Fys[0].FY.Tax + data.Fys[0].FY.VacantLandTax + data.Fys[0].FY.CompositeTax + data.Fys[0].FY.Ecess
					m2["currentYearRebate"] = data.Fys[0].FY.Rebate
					m2["currentYearPenalty"] = data.Fys[0].FY.Penalty
					m2["currentYearTotal"] = data.Fys[0].FY.TotalTax
					m2["currentYearOtherDemand"] = data.Fys[0].FY.OtherDemand

					if len(data.Fys) == 2 {
						m2["arrearYears"] = data.Fys[1].FY.Name
					} else if len(data.Fys) >= 3 {
						m2["arrearYears"] = data.Fys[len(data.Fys)-1].FY.Name + " to " + data.Fys[1].FY.Name
					}
				} else {
					if len(data.Fys) == 1 {
						m2["arrearYears"] = data.Fys[0].FY.Name
					} else if len(data.Fys) > 1 {
						m2["arrearYears"] = data.Fys[len(data.Fys)-1].FY.Name + " to" + data.Fys[0].FY.Name
					}
				}

			}
			for _, v := range data.Fys {
				if v.FY.IsCurrent {
					continue
				}
				arrearYearTax = arrearYearTax + v.FY.Tax + v.FY.VacantLandTax + v.FY.CompositeTax + v.FY.Ecess
				arrearYearRebate = arrearYearRebate + v.FY.Rebate
				arrearYearPenalty = arrearYearPenalty + v.FY.Penalty
				arrearYearTotal = arrearYearTotal + v.FY.TotalTax
				arrearYearOtherDemand = arrearYearOtherDemand + v.FY.OtherDemand

			}
			m2["arrearYearTax"] = arrearYearTax
			m2["arrearYearRebate"] = arrearYearRebate
			m2["arrearYearPenalty"] = arrearYearPenalty
			m2["arrearYearTotal"] = arrearYearTotal
			m2["arrearYearOtherDemand"] = arrearYearOtherDemand
		}
		remainingAmount := data.Demand.TotalTax - data.Details.Amount
		m := make(map[string]interface{})
		m["receipt"] = data
		m["remainingAmount"] = remainingAmount
		payments.Data = m
		payments.RefData = m2
		payments.Config = productConfig.ProductConfiguration

		pdfDta.ArrData = append(pdfDta.ArrData, payments)

	}

	pdfDta.Config = productConfig.ProductConfiguration
	fmt.Println("data", pdfDta)

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "Receipt_bulk_page_minimal.html"

	if err := r.ParseTemplate(templatePath, pdfDta); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err

	} else {
		return nil, err
	}

}
func (s *Service) BulkPrintGetDetailsForTradelicense(ctx *models.Context, filter *models.BulkPrintFilter) (BulkPrint []models.BulkPrintDetail, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.BulkPrintGetDetailsForTradelicense(ctx, filter)
}

func (s *Service) BulkPrintGetDetailsReceiptsForTradelicense(ctx *models.Context, bprr *models.BulkPrintReceiptsRequest) ([]byte, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	payments := models.PDFDataV2Arr{}
	pdfDta := new(models.PDFDataV2)
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	for _, v := range bprr.TnxIds {

		data, err := s.GetSingleTradelicensePaymentTxtID(ctx, v)
		if err != nil {
			return nil, err
		}
		m2 := make(map[string]interface{})

		m2["arrearYears"] = ""
		m2["currentYear"] = ""
		if data != nil {
			var arrearYearTax, arrearYearRebate, arrearYearPenalty, arrearYearTotal, arrearYearOtherDemand float64

			if len(data.FYs) > 0 {
				if data.FYs[0].FY.IsCurrent {
					m2["currentYear"] = data.FYs[0].FY.Name
					m2["currentYearTax"] = data.FYs[0].FY.Details.Tax
					m2["currentYearRebate"] = data.FYs[0].FY.Details.Rebate
					m2["currentYearPenalty"] = data.FYs[0].FY.Details.Penalty
					m2["currentYearTotal"] = data.FYs[0].FY.Details.TotalTaxAmount
					m2["currentYearOtherDemand"] = data.FYs[0].FY.Details.Other

					if len(data.FYs) == 2 {
						m2["arrearYears"] = data.FYs[1].FY.Name
					} else if len(data.FYs) >= 3 {
						m2["arrearYears"] = data.FYs[len(data.FYs)-1].FY.Name + " to " + data.FYs[1].FY.Name
					}
				} else {
					if len(data.FYs) == 1 {
						m2["arrearYears"] = data.FYs[0].FY.Name
					} else if len(data.FYs) > 1 {
						m2["arrearYears"] = data.FYs[len(data.FYs)-1].FY.Name + " to" + data.FYs[0].FY.Name
					}
				}

			}
			for _, v := range data.FYs {
				if v.FY.IsCurrent {
					continue
				}
				arrearYearTax = arrearYearTax + v.FY.Details.Tax
				arrearYearRebate = arrearYearRebate + v.FY.Details.Rebate
				arrearYearPenalty = arrearYearPenalty + v.FY.Details.Penalty
				arrearYearTotal = arrearYearTotal + v.FY.Details.TotalTaxAmount
				arrearYearOtherDemand = arrearYearOtherDemand + v.FY.Details.Other

			}
			m2["arrearYearTax"] = arrearYearTax
			m2["arrearYearRebate"] = arrearYearRebate
			m2["arrearYearPenalty"] = arrearYearPenalty
			m2["arrearYearTotal"] = arrearYearTotal
			m2["arrearYearOtherDemand"] = arrearYearOtherDemand
		}
		remainingAmount := data.Demand.Total.Total - data.Details.Amount
		m := make(map[string]interface{})
		m["receipt"] = data
		m["remainingAmount"] = remainingAmount
		payments.Data = m
		payments.RefData = m2
		payments.Config = productConfig.ProductConfiguration

		pdfDta.ArrData = append(pdfDta.ArrData, payments)

	}

	pdfDta.Config = productConfig.ProductConfiguration
	fmt.Println("data", pdfDta)

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "tradelicense_receipt.html"

	if err := r.ParseTemplate(templatePath, pdfDta); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err

	} else {
		return nil, err
	}

}
