package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveOurService : ""
func (s *Service) SaveOurService(ctx *models.Context, scenario string, OurService *models.OurService) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	OurService.UniqueID = s.Daos.GetUniqueID(ctx, scenario)
	OurService.Status = constants.OURSERVICESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OurService.created")
	OurService.Created = &created
	log.Println("b4 OurService.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if scenario == "uploadpdf" {
			err := s.Daos.UpdateManyDisableOurService(ctx, scenario)
			if err != nil {
				return err
			}
		}
		dberr := s.Daos.SaveOurService(ctx, scenario, OurService)
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

// GetSingleOurService : ""
func (s *Service) GetSingleOurService(ctx *models.Context, scenario string, UniqueID string) (*models.RefOurService, error) {
	OurService, err := s.Daos.GetSingleOurService(ctx, scenario, UniqueID)
	if err != nil {
		return nil, err
	}
	return OurService, nil
}

// UpdateOurService : ""
func (s *Service) UpdateOurService(ctx *models.Context, scenario string, OurService *models.OurService) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOurService(ctx, scenario, OurService)
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

// EnableOurService : ""
func (s *Service) EnableOurService(ctx *models.Context, scenario string, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if scenario == "uploadpdf" {
			err := s.Daos.UpdateManyDisableOurService(ctx, scenario)
			if err != nil {
				return err
			}
		}
		dberr := s.Daos.EnableOurService(ctx, scenario, uniqueID)
		if dberr != nil {
			return dberr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Aborted with error" + err1.Error())
		}
		return err
	}

	return nil
}

// DisableOurService : ""
func (s *Service) DisableOurService(ctx *models.Context, scenario string, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableOurService(ctx, scenario, uniqueID)
		if debrr != nil {
			return debrr
		}
		if err := sc.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		if err1 := ctx.Session.AbortTransaction(ctx.CTX); err1 != nil {
			log.Println("err in abort")
			return errors.New("Transaction Abort with error" + err1.Error())
		}
		return err
	}
	return nil
}

//DeleteOurService : ""
func (s *Service) DeleteOurService(ctx *models.Context, scenario string, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOurService(ctx, scenario, UniqueID)
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

// FilterOurService : ""
func (s *Service) FilterOurService(ctx *models.Context, scenario string, OurService *models.FilterOurService, pagination *models.Pagination) (OurServices []models.RefOurService, err error) {
	return s.Daos.FilterOurService(ctx, scenario, OurService, pagination)
}
