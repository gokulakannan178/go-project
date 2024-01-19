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

// SavePropertyPenalty : ""
func (s *Service) SaveLetterGenerate(ctx *models.Context, lg *models.LetterGenerate) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if lg.UniqueID == "" {
		lg.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLETTERGENERATE)
	}
	lg.Status = constants.LETTERGENERATESTATUSDRAFT

	t := time.Now()
	lg.Submitted.On = &t
	lg.Submitted.Action = constants.LETTERGENERATESTATUSDRAFT
	created := models.CreatedV2{}
	created.On = &t
	lg.Created = &created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveLetterGenerate(ctx, lg)
		if dberr != nil {
			return dberr
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
		return err
	}
	return nil
}

//GetSingleLetterGenerate :""
func (s *Service) GetSingleLetterGenerate(ctx *models.Context, UniqueID string) (*models.RefLetterGenerate, error) {
	lg, err := s.Daos.GetSingleLetterGenerate(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return lg, nil
}

// UpdateLetterGenerate : ""
func (s *Service) UpdateLetterGenerate(ctx *models.Context, lg *models.LetterGenerate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.UpdateLetterGenerate(ctx, lg)
		if dberr != nil {
			return dberr
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
		return err
	}
	return nil
}

// EnableLetterGenerate : ""
func (s *Service) EnableLetterGenerate(ctx *models.Context, lg *models.LetterGenerateAction) error {
	t := time.Now()
	lg.On = &t
	lg.Action = constants.LETTERGENERATESTATUSACTIVE

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.EnableLetterGenerate(ctx, lg)
		if dberr != nil {
			return dberr
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
		return err
	}
	return nil
}
func (s *Service) ApprovedLetterGenerate(ctx *models.Context, lg *models.LetterGenerateAction) error {
	t := time.Now()
	lg.On = &t
	lg.Action = constants.LETTERGENERATESTATUSAPPROVED

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.ApprovedLetterGenerate(ctx, lg)
		if dberr != nil {
			return dberr
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
		return err
	}
	return nil
}

//DisableLetterGenerate : ""
func (s *Service) DisableLetterGenerate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.DisableLetterGenerate(ctx, UniqueID)
		if dberr != nil {
			return dberr
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
		return err
	}
	return nil
}

// DeleteLetterGenerate : ""
func (s *Service) DeleteLetterGenerate(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.DeleteLetterGenerate(ctx, UniqueID)
		if dberr != nil {
			return dberr
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
		return err
	}
	return nil
}

// BlockedLetterGenerate : ""
func (s *Service) BlockedLetterGenerate(ctx *models.Context, lg *models.LetterGenerateAction) error {
	lg.Action = constants.LETTERGENERATESTATUSBLOCKED
	t := time.Now()
	lg.On = &t
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.BlockedLetterGenerate(ctx, lg)
		if dberr != nil {
			return dberr
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
		return err
	}
	return nil
}

// SubmittedLetterGenerate : ""
func (s *Service) SubmittedLetterGenerate(ctx *models.Context, lg *models.LetterGenerateAction) error {
	if lg.UniqueID == "" {
		lg.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONLETTERGENERATE)

	}

	lg.Action = constants.LETTERGENERATESTATUSSUBMITTED
	t := time.Now()
	lg.On = &t
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.SubmittedLetterGenerate(ctx, lg)
		if dberr != nil {
			return dberr
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
		return err
	}
	return nil
}
func (s *Service) UploadLetterGenerate(ctx *models.Context, lg *models.LetterGenerate) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.UploadLetterGenerate(ctx, lg)
		if dberr != nil {
			return dberr
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
		return err
	}
	return nil
}

// FilterLetterGenerate : ""
func (s *Service) FilterLetterGenerate(ctx *models.Context, filter *models.LetterGenerateFilter, pagination *models.Pagination) ([]models.RefLetterGenerate, error) {
	return s.Daos.FilterLetterGenerate(ctx, filter, pagination)

}

//LetterGenerateExecute :""
func (s *Service) LetterGenerateExecute(ctx *models.Context, UniqueID string) ([]byte, *models.RefLetterGenerate, error) {
	lg, err := s.Daos.GetSingleLetterGenerate(ctx, UniqueID)
	if err != nil {
		return nil, nil, err
	}

	r := NewRequestPdf("")
	templatePathStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.TEMPLATELOC)
	//html template path
	templatePath := templatePathStart + "letterhead.html"

	//path for download pdf
	// t := time.Now()
	// outputPath := fmt.Sprintf("storage/SampleTemplate%v.pdf", t.Unix())
	m := make(map[string]interface{})
	m["letter"] = lg
	m2 := make(map[string]interface{})
	m2["currentDate"] = time.Now()
	productConfigUniqueID := "1"
	productConfig, err := s.Daos.GetSingleProductConfiguration(ctx, productConfigUniqueID)
	if err != nil {
		return nil, nil, errors.New("Error in geting product config" + err.Error())
	}
	var pdfdata models.PDFData
	pdfdata.Data = m
	pdfdata.RefData = m2
	pdfdata.Config = productConfig.ProductConfiguration
	marginzero := uint(0)
	r.SetMargin(&marginzero, &marginzero, &marginzero, &marginzero)
	fmt.Println(m)
	if err := r.ParseTemplate(templatePath, pdfdata); err == nil {
		fmt.Println("start pdf generated successfully")

		ok, data, err := r.GeneratePDFAsFile()

		fmt.Println(ok, "pdf generated successfully")
		return data, lg, err
	} else {
		fmt.Println("Error in parcing template - " + err.Error())

		return nil, nil, errors.New("Error in parcing template - " + err.Error())
	}

}
