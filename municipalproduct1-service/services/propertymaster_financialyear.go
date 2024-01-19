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

//SaveFinancialYear :""
func (s *Service) SaveFinancialYear(ctx *models.Context, financialYear *models.FinancialYear) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	financialYear.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONFINANCIALYEAR)
	financialYear.Status = constants.FINANCIALYEARSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 financialYear.created")
	financialYear.Created = created
	log.Println("b4 financialYear.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveFinancialYear(ctx, financialYear)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//UpdateFinancialYear : ""
func (s *Service) UpdateFinancialYear(ctx *models.Context, financialYear *models.FinancialYear) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateFinancialYear(ctx, financialYear)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//EnableFinancialYear : ""
func (s *Service) EnableFinancialYear(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableFinancialYear(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DisableFinancialYear : ""
func (s *Service) DisableFinancialYear(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableFinancialYear(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//DeleteFinancialYear : ""
func (s *Service) DeleteFinancialYear(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteFinancialYear(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

//GetSingleFinancialYear :""
func (s *Service) GetSingleFinancialYear(ctx *models.Context, UniqueID string) (*models.RefFinancialYear, error) {
	financialYear, err := s.Daos.GetSingleFinancialYear(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return financialYear, nil
}

//FilterFinancialYear :""
func (s *Service) FilterFinancialYear(ctx *models.Context, financialYearfilter *models.FinancialYearFilter, pagination *models.Pagination) (financialYear []models.RefFinancialYear, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.FilterFinancialYear(ctx, financialYearfilter, pagination)
}

// /FilterFinancialYearPDF : ""
func (s *Service) FilterFinancialYearPDF(ctx *models.Context, filter *models.FinancialYearFilter, p *models.Pagination) ([]byte, error) {
	properties, err := s.Daos.FilterFinancialYear(ctx, filter, p)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m["demand"] = properties
	m2["currentDate"] = time.Now()
	m3["inc"] = func(a int) int {
		return a + 1
	}
	cfy, err := s.Daos.GetCurrentFinancialYear(ctx)
	if err != nil {
		return nil, errors.New("Error in getting current financial year " + err.Error())
	}
	m2["currentFy"] = cfy
	m2["inc"] = func(a int) int {
		return a + 1
	}
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
	templatePath := templatePathStart + "property_penaltyrate_report.html"
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

//MakeCurrentFinancialYear : ""
func (s *Service) MakeCurrentFinancialYear(ctx *models.Context, uniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.MakeCurrentFinancialYear(ctx, uniqueID)
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
			return errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
	}
	return nil
}

//GetCurrentFinancialYear : ""
func (s *Service) GetCurrentFinancialYear(ctx *models.Context) (*models.RefFinancialYear, error) {

	return s.Daos.GetCurrentFinancialYear(ctx)
}

// GetSingleFinancialYearUsingDate : ""
func (s *Service) GetSingleFinancialYearUsingDate(ctx *models.Context, Date *time.Time) (*models.RefFinancialYear, error) {
	financialYear, err := s.Daos.GetSingleFinancialYearUsingDate(ctx, Date)
	if err != nil {
		return nil, err
	}
	return financialYear, nil
}
