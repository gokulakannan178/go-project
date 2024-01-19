package services

import (
	"errors"
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"
)

//CalcShopRentDemand : ""
func (s *Service) CalcShopRentDemand(ctx *models.Context, uniqueID string) (*models.ShopRentDemand, error) {
	mtd := new(models.ShopRentDemand)
	mtd.RefShopRent.UniqueID = uniqueID
	var dberr error
	mainpipeline, err := mtd.CalcQuery(nil)
	if err != nil {
		return nil, errors.New("Error in generating Query - " + err.Error())
	}
	mtd, dberr = s.Daos.CalcShopRentDemand(ctx, mainpipeline)
	if dberr != nil {
		return nil, dberr
	}
	dberr = mtd.CalcDemand()
	if dberr != nil {
		return nil, errors.New("Error in calculating demand " + dberr.Error())
	}
	return mtd, nil
}

//CalcShopRentDemand : ""
func (s *Service) CalcShopRentDemandForParticulars(ctx *models.Context, filter *models.ShopRentCalcQueryFilter) (*models.ShopRentDemand, error) {
	mtd := new(models.ShopRentDemand)
	mtd.RefShopRent.UniqueID = filter.ShopRentID
	var dberr error
	mainpipeline, err := mtd.CalcQuery(filter)
	if err != nil {
		return nil, errors.New("Error in generating Query - " + err.Error())
	}
	mtd, dberr = s.Daos.CalcShopRentDemand(ctx, mainpipeline)
	if dberr != nil {
		return nil, dberr
	}
	dberr = mtd.CalcDemand()
	if dberr != nil {
		return nil, errors.New("Error in calculating demand " + dberr.Error())
	}
	return mtd, nil
}

func (s *Service) ReCalcShopRentDemandWithOutTransaction(ctx *models.Context, uniqueID string) (*models.ShopRentDemand, error) {
	mtd, dberr := s.CalcShopRentDemand(ctx, uniqueID)
	if dberr != nil {
		return nil, dberr
	}
	mtd.ShopRent.Collections = models.ShopRentTotalCollection{}
	mtd.ShopRent.PendingCollections = models.ShopRentTotalCollection{}
	mtd.ShopRent.OutStanding = models.ShopRentTotalOutStanding{}

	return mtd, nil
}

func (s *Service) CalcShopRentDemandForParticularsPDF(ctx *models.Context, uniqueID string) ([]byte, error) {
	demand, err := s.CalcShopRentDemand(ctx, uniqueID)
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

func (s *Service) ShopRentGetOutstandingDemand(ctx *models.Context, uniqueID string) (*models.ShopRentDemand, error) {
	filter := new(models.ShopRentCalcQueryFilter)

	filter.ShopRentID = uniqueID
	omitFys, err := s.Daos.GetPayedFinancialYearsOfShoprent(ctx, uniqueID)
	if err != nil {
		return nil, err
	}
	if len(omitFys) > 0 {
		filter.OmitFy = append(filter.OmitFy, omitFys...)
	}
	fmt.Println("omity = >", filter.OmitFy)
	return s.CalcShopRentDemandForParticulars(ctx, filter)
}

func (s *Service) SHopRentGetOutstandingDemandPDF(ctx *models.Context, uniqueID string) ([]byte, error) {
	demand, err := s.ShopRentGetOutstandingDemand(ctx, uniqueID)
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
