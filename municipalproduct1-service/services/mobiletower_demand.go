package services

import (
	"errors"
	"fmt"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//CalcMobileTowerDemand : ""
func (s *Service) ReCalcMobileTowerDemandWithTransaction(ctx *models.Context, uniqueID string, filter *models.MobileTowerCalcQueryFilter) (*models.MobileTowerDemand, error) {
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	mtd := new(models.MobileTowerDemand)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		var err error
		mtd, err = s.ReCalcMobileTowerDemandWithOutTransaction(ctx, uniqueID, filter)
		if err != nil {
			return err
		}
		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return nil, errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return nil, err
	}
	return mtd, nil
}
func (s *Service) ReCalcMobileTowerDemandWithOutTransaction(ctx *models.Context, uniqueID string, filter *models.MobileTowerCalcQueryFilter) (*models.MobileTowerDemand, error) {
	mtd, dberr := s.CalcMobileTowerDemand(ctx, uniqueID, filter)
	if dberr != nil {
		return nil, dberr
	}
	mtd.PropertyMobileTower.Collections = models.MobileTowerTotalCollection{}
	mtd.PropertyMobileTower.PendingCollections = models.MobileTowerTotalCollection{}
	mtd.PropertyMobileTower.OutStanding = models.MobileTowerTotalOutStanding{}
	dberr = s.CalcMobileTowerCollections(ctx, mtd)
	if dberr != nil {
		return nil, dberr
	}
	dberr = s.CalcMobileTowerPendingCollections(ctx, mtd)
	if dberr != nil {
		return nil, dberr
	}
	if err := s.CalcMobileTowerOutStanding(ctx, mtd); err != nil {
		return nil, err
	}
	if err := s.Daos.UpdateMobileTowerCalc(ctx, mtd); err != nil {
		return nil, err
	}

	return mtd, nil
}

//CalcMobileTowerDemand : ""
func (s *Service) CalcMobileTowerDemand(ctx *models.Context, uniqueID string, filter *models.MobileTowerCalcQueryFilter) (*models.MobileTowerDemand, error) {
	mtd := new(models.MobileTowerDemand)
	mtd.RefPropertyMobileTower.UniqueID = uniqueID
	var dberr error
	mainpipeline, err := mtd.CalcQuery(filter)
	if err != nil {
		return nil, errors.New("Error in generating Query - " + err.Error())
	}
	mtd, dberr = s.Daos.CalcMobileTowerDemand(ctx, mainpipeline)
	if dberr != nil {
		return nil, dberr
	}
	dberr = mtd.CalcDemand()
	if dberr != nil {
		return nil, errors.New("Error in calculating demand " + dberr.Error())
	}
	var registration *models.RefMobileTowerRegistrationTax
	if mtd.IsRegPaid == 0 {
		registration, err = s.Daos.GetSingleDefaultMobileTowerRegistrationTax(ctx)
		if err != nil {
			return nil, err
		}
	}
	if registration == nil {
		registration = new(models.RefMobileTowerRegistrationTax)
	}
	mtd.Ref.MobileTowerRegistrationTax = *registration
	mtd.UnPaid = mtd.Ref.MobileTowerRegistrationTax.Value + mtd.Demand.Total.Total
	return mtd, nil
}

//
func (s *Service) MobileTowerDemandPdf(ctx *models.Context, uniqueID string, filter *models.MobileTowerCalcQueryFilter) ([]byte, error) {
	data, err := s.CalcMobileTowerDemand(ctx, uniqueID, filter)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m["demand"] = data
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
	templatePath := templatePathStart + "mobiletower_demand.html"
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

//CalcMobileTowerCollections : ""
func (s *Service) CalcMobileTowerCollections(ctx *models.Context, mtd *models.MobileTowerDemand) error {

	var dberr error
	mainpipeline, err := mtd.CalcCollectionQuery()
	if err != nil {
		return errors.New("Error in generating Query - " + err.Error())
	}
	mtp, dberr := s.Daos.CalcMobileTowerPaymens(ctx, mainpipeline)
	if dberr != nil {
		return dberr
	}
	dberr = mtd.CalcCollection(mtp)
	if dberr != nil {
		return errors.New("Error in calculating demand " + dberr.Error())
	}
	return nil
}

//CalcMobileTowerPendingCollections : ""
func (s *Service) CalcMobileTowerPendingCollections(ctx *models.Context, mtd *models.MobileTowerDemand) error {

	var dberr error
	mainpipeline, err := mtd.CalcPendingCollectionQuery()
	if err != nil {
		return errors.New("Error in generating Query - " + err.Error())
	}
	mtp, dberr := s.Daos.CalcMobileTowerPendingPaymens(ctx, mainpipeline)
	if dberr != nil {
		return dberr
	}
	dberr = mtd.CalcPendingCollection(mtp)
	if dberr != nil {
		return errors.New("Error in calculating demand " + dberr.Error())
	}
	return nil
}

//CalcMobileTowerPendingCollections : ""
func (s *Service) CalcMobileTowerOutStanding(ctx *models.Context, mtd *models.MobileTowerDemand) error {

	dberr := mtd.CalcOutStanding()
	if dberr != nil {
		return errors.New("Error in calculating demand " + dberr.Error())
	}
	return nil
}
