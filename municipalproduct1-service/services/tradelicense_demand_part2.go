package services

import (
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"
)

// TradeLicenseDemandCalcPart2 : ""
func (s *Service) TradeLicenseDemandCalcPart2(ctx *models.Context, filter *models.TradeLicenseDemandPart2Filter) (*models.RefTradeLicenseDemandPart2, error) {
	resTlDemand, err := s.Daos.TradeLicenseDemandPart2(ctx, filter)
	if err != nil {
		return nil, errors.New("no data found" + err.Error())
	}
	resTlDemand.CalcDemand()
	fmt.Println("======> before for loop")
	for _, v := range resTlDemand.FYs {
		fmt.Println("======> inside for loop")

		if v.FinancialYear.IsCurrent == true {
			resTlDemand.Demand.Current.Tax = resTlDemand.Demand.Current.Tax + v.FYTax
			resTlDemand.Demand.Current.Penalty = resTlDemand.Demand.Current.Penalty + v.FYTLPenalty
			resTlDemand.Demand.Current.Rebate = resTlDemand.Demand.Current.Rebate + v.FYRebate
			resTlDemand.Demand.Current.Other = resTlDemand.Demand.Current.Other + v.FYOther
			resTlDemand.Demand.Current.Total = resTlDemand.Demand.Current.Total + v.FYTotal
		} else {
			resTlDemand.Demand.Arrear.Tax = resTlDemand.Demand.Arrear.Tax + v.FYTax
			resTlDemand.Demand.Arrear.Penalty = resTlDemand.Demand.Arrear.Penalty + v.FYTLPenalty
			resTlDemand.Demand.Arrear.Rebate = resTlDemand.Demand.Arrear.Rebate + v.FYRebate
			// resTlDemand.Demand.Arrear.Rebate = resTlDemand.Demand.Arrear.Rebate + v.FYRebateValue
			resTlDemand.Demand.Arrear.Other = resTlDemand.Demand.Arrear.Other + v.FYOther
			resTlDemand.Demand.Arrear.Total = resTlDemand.Demand.Arrear.Total + v.FYTotal
		}
	}
	fmt.Println("======> after for loop")

	resTlDemand.Demand.Total.Tax = resTlDemand.Demand.Current.Tax + resTlDemand.Demand.Arrear.Tax
	resTlDemand.Demand.Total.Total = resTlDemand.Demand.Current.Total + resTlDemand.Demand.Arrear.Total
	resTlDemand.Demand.Total.Penalty = resTlDemand.Demand.Current.Penalty + resTlDemand.Demand.Arrear.Penalty
	resTlDemand.Demand.Total.Other = resTlDemand.Demand.Current.Other + resTlDemand.Demand.Arrear.Other
	resTlDemand.Demand.Total.Rebate = resTlDemand.Demand.Current.Rebate + resTlDemand.Demand.Arrear.Rebate

	return resTlDemand, nil
}

// GetTradeLicensePaymentDemandPDFPart2 : ""
func (s *Service) GetTradeLicensePaymentDemandPDFPart2(ctx *models.Context, filter *models.TradeLicenseDemandPart2Filter) ([]byte, error) {
	r := NewRequestPdf("")
	data, err := s.TradeLicenseDemandCalcPart2(ctx, filter)
	if err != nil {
		return nil, err
	}
	// fmt.Println(data.ReciptNo)
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["Payment"] = data
	m2["currentDate"] = time.Now()
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	pdfdata.Config = productConfig.ProductConfiguration

	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "tradelicense_demand_report_v2.html"
	err = r.ParseTemplate(templatePath, pdfdata)
	if err != nil {
		return nil, err
	}
	ok, file, err := r.GeneratePDFAsFile()
	if err != nil {
		return nil, err
	}
	fmt.Println(ok, "pdf generated successfully")

	return file, nil
}
