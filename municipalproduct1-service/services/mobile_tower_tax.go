package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// PreSaveMobileTower : ""
func (s *Service) PreSaveMobileTower(ctx *models.Context, mobileTower *models.PropertyMobileTower) {

	mobileTower.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWER)

	t := time.Now()
	// created := models.CreatedV2{}
	mobileTower.Created.On = &t
	// created.By = mobileTower.CreatedBy
	// created.ByType = mobileTower.CreatedType
	log.Println("b4 user.created")
	// mobileTower.Created = created
	mobileTower.Status = constants.MOBILETOWERSTATUSPENDING
	log.Println("b4 user.created")
	// for k := range mobileTower.MobileTowerDemandFY {
	// 	mobileTower.MobileTowerDemandFY[k].PropertyID = mobileTower.PropertyID
	// 	mobileTower.MobileTowerDemandFY[k].UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWERYEAR)
	// 	mobileTower.MobileTowerDemandFY[k].MobileTowerID = mobileTower.UniqueID
	// 	mobileTower.MobileTowerDemandFY[k].Status = constants.MOBILETOWERFYSTATUSACTIVE
	// }

}

// SaveMobileTowerTax : ""
func (s *Service) SaveMobileTowerTax(ctx *models.Context, mobile *models.MobileTowerTax) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	mobile.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMOBILETOWERTAX)
	mobile.Status = constants.MOBILETOWERTAXSTATUSACTIVE
	t := time.Now()
	mobile.Created.On = &t
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveMobileTowerTax(ctx, mobile)
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

//GetSingleMobileTowerTax :""
func (s *Service) GetSingleMobileTowerTax(ctx *models.Context, UniqueID string) (*models.MobileTowerTax, error) {
	tower, err := s.Daos.GetSingleMobileTowerTax(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateMobileTowerTax : ""
func (s *Service) UpdateMobileTowerTax(ctx *models.Context, mobile *models.MobileTowerTax) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateMobileTowerTax(ctx, mobile)
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

// EnableMobileTowerTax : ""
func (s *Service) EnableMobileTowerTax(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableMobileTowerTax(ctx, UniqueID)
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

//DisableMobileTowerTax : ""
func (s *Service) DisableMobileTowerTax(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableMobileTowerTax(ctx, UniqueID)
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

// DeleteMobileTowerTax : ""
func (s *Service) DeleteMobileTowerTax(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteMobileTowerTax(ctx, UniqueID)
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

// FilterMobileTowerTax : ""
func (s *Service) FilterMobileTowerTax(ctx *models.Context, filter *models.MobileTowerTaxFilter, pagination *models.Pagination) ([]models.RefMobileTowerTax, error) {
	return s.Daos.FilterMobileTowerTax(ctx, filter, pagination)

}
