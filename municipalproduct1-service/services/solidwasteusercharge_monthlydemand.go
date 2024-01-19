package services

import (
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"
)

//CalcShopRentMonthlyDemandForParticulars : ""
func (s *Service) SolidWasteUserChargeDemandForParticulars(ctx *models.Context, filter *models.SolidWasteUserChargeDemandCalcQueryFilter) (*models.SolidWasteUserChargeDemand, error) {
	fmt.Println("server is starting")

	mtd := new(models.SolidWasteUserChargeDemand)
	mtd.RefSolidWasteUserCharge.UniqueID = filter.SolidWasteChargeID

	var dberr error
	mainpipeline, err := mtd.CalcQuery(filter)
	if err != nil {
		return nil, errors.New("Error in generating Query - " + err.Error())
	}
	mtd, dberr = s.Daos.CalcSolidWasteUserChargeDemand(ctx, mainpipeline)
	if dberr != nil {
		return nil, dberr
	}
	fmt.Println("mtd is clear")
	if mtd == nil {
		return nil, errors.New("demand is nil")
	}
	dberr = mtd.CalcDemand()
	if dberr != nil {
		return nil, errors.New("Error in calculating demand " + dberr.Error())
	}
	fmt.Println("dberr is clear")
	return mtd, nil
}

// SolidWasteUserChargeMonthlyDemandPDF : ""
func (s *Service) SolidWasteUserChargeMonthlyDemandPDF(ctx *models.Context, uniqueID string) ([]byte, error) {
	filter := new(models.SolidWasteUserChargeDemandCalcQueryFilter)
	filter.SolidWasteChargeID = uniqueID
	filter.OmitPayedMonths = true
	demand, err := s.SolidWasteUserChargeDemandForParticulars(ctx, filter)
	if err != nil {
		return nil, err
	}
	date := s.Shared.EndOfMonth(time.Now())
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["demand"] = demand
	m2["currentDate"] = time.Now()
	m["validityDate"] = date

	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "solidwaste_monthlydemand.html"
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, errors.New("Error in parcing template - " + err.Error())
	}
	return nil, nil
}
