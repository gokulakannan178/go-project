package services

import (
	"context"
	"errors"
	"fmt"
	"municipalproduct1-service/app"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"
)

//CalcShopRentMonthlyDemandForParticulars : ""
func (s *Service) CalcShopRentMonthlyDemandForParticulars(ctx *models.Context, filter *models.ShopRentMonthlyCalcQueryFilter) (*models.ShopRentMonthlyDemand, error) {
	fmt.Println("server is starting")

	mtd := new(models.ShopRentMonthlyDemand)
	mtd.RefShopRent.UniqueID = filter.ShopRentID

	var dberr error
	mainpipeline, err := mtd.CalcQuery(filter)
	if err != nil {
		return nil, errors.New("Error in generating Query - " + err.Error())
	}
	mtd, dberr = s.Daos.CalcShopRentMonthlyDemand(ctx, mainpipeline)
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

func (s *Service) SHopRentMonthlyGetOutstandingDemandPDF(ctx *models.Context, uniqueID string) ([]byte, error) {
	filter := new(models.ShopRentMonthlyCalcQueryFilter)
	filter.ShopRentID = uniqueID
	filter.OmitPayedMonths = true
	demand, err := s.CalcShopRentMonthlyDemandForParticulars(ctx, filter)
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
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, errors.New("Error in getting product config" + err.Error())
	}
	pdfdata.Config = productConfig.ProductConfiguration

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "shoprent_monthlydemand.html"
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

//CalcShopRentOverallMonthlyDemand : "Calculates overall shoprent demand"
/*
* By - solomon
* ctx *models.Context - context
* shoprentID - shoprent unique id
* doUpdate - true to update the demand with sgoprent collection or false to not update
* returns demand and error
 */
func (s *Service) CalcShopRentOverallMonthlyDemand(ctx *models.Context, shoprentID string, doUpdate bool) (*models.ShopRentMonthlyDemand, error) {
	filter := new(models.ShopRentMonthlyCalcQueryFilter)
	filter.ShopRentID = shoprentID
	mtd := new(models.ShopRentMonthlyDemand)
	mtd.RefShopRent.UniqueID = shoprentID
	demand, err := s.CalcShopRentMonthlyDemandForParticulars(ctx, filter)
	if err != nil {
		return nil, err
	}
	if demand == nil {
		return nil, errors.New("nil demand")
	}
	if doUpdate {
		fmt.Println("updating", shoprentID)
		currFy, err := s.Daos.GetCurrentFinancialYear(ctx)
		if err != nil {
			return nil, errors.New("Err in geting current fynancial yesr - " + err.Error())
		}
		if currFy == nil {
			return nil, errors.New("nil current fy")
		}
		srDemand := new(models.ShopRentTotalDemand)
		for _, v := range demand.FY {
			if v.FinancialYear.UniqueID == currFy.UniqueID {
				srDemand.Current.Tax = srDemand.Current.Tax + v.Details.Tax
				srDemand.Current.Penalty = srDemand.Current.Penalty + v.Details.Penalty
				srDemand.Current.Other = srDemand.Current.Other + v.Details.Other
				srDemand.Current.Total = srDemand.Current.Total + v.Details.TotalTaxAmount
				continue
			}
			srDemand.Arrear.Tax = srDemand.Arrear.Tax + v.Details.Tax
			srDemand.Arrear.Penalty = srDemand.Arrear.Penalty + v.Details.Penalty
			srDemand.Arrear.Other = srDemand.Arrear.Other + v.Details.Other
			srDemand.Arrear.Total = srDemand.Arrear.Total + v.Details.TotalTaxAmount
		}
		srDemand.Total.Tax = srDemand.Current.Tax + srDemand.Arrear.Tax
		srDemand.Total.Penalty = srDemand.Current.Penalty + srDemand.Arrear.Penalty
		srDemand.Total.Other = srDemand.Current.Other + srDemand.Arrear.Other
		srDemand.Total.Total = srDemand.Current.Total + srDemand.Arrear.Total

		if err := s.Daos.UpdateOverallShopRentDemand(ctx, shoprentID, srDemand); err != nil {
			return nil, errors.New("Err saving overall demand - " + err.Error())
		}
	}
	return demand, nil
}

// SaveShopRentDemandForAll : ""
func (s *Service) SaveShopRentDemandForAll(IDs []string) error {

	for k, v := range IDs {
		fmt.Println("k------>", k)
		c := context.TODO()
		ctx := app.GetApp(c, s.Daos)
		defer ctx.Client.Disconnect(c)

		_, err := s.CalcShopRentOverallMonthlyDemand(ctx, v, true)
		fmt.Println(v, err)
	}
	return nil
}
