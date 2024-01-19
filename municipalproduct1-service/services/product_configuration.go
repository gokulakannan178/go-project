package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveProductConfiguration : ""
func (s *Service) SaveProductConfiguration(ctx *models.Context, pc *models.ProductConfiguration) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	pc.UniqueID = "1"
	pc.Status = constants.PRODUCTCONFIGURATIONACTIVE
	t := time.Now()
	pc.Created.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.UpsertProductConfiguration(ctx, pc)
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

//GetSingleProductConfiguration : ""
func (s *Service) GetSingleProductConfiguration(ctx *models.Context) (*models.RefProductConfiguration, error) {
	uniqueID := "1"
	return s.Daos.GetSingleProductConfiguration(ctx, uniqueID)
}

//GetProductLogo : ""
func (s *Service) GetProductLogo(ctx *models.Context) (*models.Logo, error) {
	uniqueID := "2"
	return s.Daos.GetProductLogo(ctx, uniqueID)
}

//GetWatermarkLogo : ""
func (s *Service) GetWatermarkLogo(ctx *models.Context) (*models.WatermarkLogo, error) {
	uniqueID := "3"
	return s.Daos.GetWatermarkLogo(ctx, uniqueID)
}

// FilterProductConfiguration : ""
func (s *Service) FilterProductConfiguration(ctx *models.Context, filter *models.ProductConfigurationFilter, pagination *models.Pagination) ([]models.RefProductConfiguration, error) {
	return s.Daos.FilterProductConfiguration(ctx, filter, pagination)

}
