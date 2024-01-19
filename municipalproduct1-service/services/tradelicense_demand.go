package services

import (
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"
)

//CalcTradeLicenseDemand : ""
func (s *Service) CalcTradeLicenseDemand(ctx *models.Context, uniqueID string) (*models.TradeLicenseDemand, error) {
	mtd := new(models.TradeLicenseDemand)
	mtd.RefTradeLicense.UniqueID = uniqueID
	var dberr error
	mainpipeline, err := mtd.CalcQuery(nil)
	if err != nil {
		return nil, errors.New("Error in generating Query - " + err.Error())
	}
	mtd, dberr = s.Daos.CalcTradeLicenseDemand(ctx, mainpipeline)
	if dberr != nil {
		return nil, dberr
	}
	dberr = mtd.CalcDemand()
	if dberr != nil {
		return nil, errors.New("Error in calculating demand " + dberr.Error())
	}
	return mtd, nil
}

//CalcTradeLicenseDemand : ""
func (s *Service) CalcTradeLicenseDemandForParticulars(ctx *models.Context, filter *models.TradeLicenseCalcQueryFilter) (*models.TradeLicenseDemand, error) {
	mtd := new(models.TradeLicenseDemand)
	mtd.RefTradeLicense.UniqueID = filter.TradeLicenseID
	var dberr error
	mainpipeline, err := mtd.CalcQuery(filter)
	if err != nil {
		return nil, errors.New("Error in generating Query - " + err.Error())
	}
	mtd, dberr = s.Daos.CalcTradeLicenseDemand(ctx, mainpipeline)
	if dberr != nil {
		return nil, dberr
	}
	dberr = mtd.CalcDemand()
	if dberr != nil {
		return nil, errors.New("Error in calculating demand " + dberr.Error())
	}
	return mtd, nil
}

func (s *Service) ReCalcTradeLicenseDemandWithOutTransaction(ctx *models.Context, uniqueID string) (*models.TradeLicenseDemand, error) {
	mtd, dberr := s.CalcTradeLicenseDemand(ctx, uniqueID)
	if dberr != nil {
		return nil, dberr
	}
	mtd.TradeLicense.Collections = models.TradeLicenseTotalCollection{}
	mtd.TradeLicense.PendingCollections = models.TradeLicenseTotalCollection{}
	mtd.TradeLicense.OutStanding = models.TradeLicenseTotalOutStanding{}

	return mtd, nil
}

func (s *Service) CalcTradeLicenseDemandForParticularsPDF(ctx *models.Context, uniqueID string) ([]byte, error) {
	demand, err := s.CalcTradeLicenseDemand(ctx, uniqueID)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["demand"] = demand
	m2["currentDate"] = time.Now()
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "shoprent_demand.html"
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

func (s *Service) TradeLicenseGetOutstandingDemand(ctx *models.Context, uniqueID string) (*models.TradeLicenseDemand, error) {
	filter := new(models.TradeLicenseCalcQueryFilter)

	filter.TradeLicenseID = uniqueID
	omitFys, err := s.Daos.GetPayedFinancialYearsOfTradeLicense(ctx, uniqueID)
	if err != nil {
		return nil, err
	}
	if len(omitFys) > 0 {
		filter.OmitFy = append(filter.OmitFy, omitFys...)
	}
	fmt.Println("omity = >", filter.OmitFy)
	return s.CalcTradeLicenseDemandForParticulars(ctx, filter)
}

func (s *Service) TradeLicenseGetOutstandingDemandPDF(ctx *models.Context, uniqueID string) ([]byte, error) {
	demand, err := s.TradeLicenseGetOutstandingDemand(ctx, uniqueID)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["demand"] = demand
	m2["currentDate"] = time.Now()
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in geting product config" + err.Error())
	}

	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	pdfdata.Config = productConfig.ProductConfiguration
	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "tradelicense_demand.html"
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
