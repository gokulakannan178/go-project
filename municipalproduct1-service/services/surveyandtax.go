package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveSurveyAndTax : ""
func (s *Service) SaveSurveyAndTax(ctx *models.Context, sat *models.SurveyAndTax) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	sat.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSURVEYANDTAX)
	sat.Status = constants.SURVEYANDTAXSTATUSACTIVE
	t := time.Now()
	sat.Created.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveSurveyAndTax(ctx, sat)
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

//GetSingleSurveyAndTax : ""
func (s *Service) GetSingleSurveyAndTax(ctx *models.Context, uniqueID string) (*models.RefSurveyAndTax, error) {
	return s.Daos.GetSingleSurveyAndTax(ctx, uniqueID)
}

// PushNotification : ""
func (s *Service) PushNotification(ctx *models.Context, uniqueID string) (*models.SurveyAndTax, error) {
	return nil, nil
}

//SurveyAndTaxFilter
func (s *Service) SurveyAndTaxFilter(ctx *models.Context, satf *models.SurveyAndTaxFilter, pagination *models.Pagination) ([]models.RefSurveyAndTax, error) {
	return s.Daos.SurveyAndTaxFilter(ctx, satf, pagination)
}
