package services

import (
	"errors"
	"log"
	"time"

	"ecommerce-service/constants"
	"ecommerce-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveScenario : ""
func (s *Service) SaveScenario(ctx *models.Context, Scenario *models.Scenario) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	Scenario.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSCENARIO)
	Scenario.Status = constants.SCENARIOSTATUSACTIVE
	t := time.Now()

	created := new(models.Created)
	created.On = &t
	created.By = constants.SYSTEM
	Scenario.Created = *created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveScenario(ctx, Scenario)
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

//GetSingleScenario :""
func (s *Service) GetSingleScenario(ctx *models.Context, UniqueID string) (*models.RefScenario, error) {
	Scenario, err := s.Daos.GetSingleScenario(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Scenario, nil
}

// UpdateScenario : ""
func (s *Service) UpdateScenario(ctx *models.Context, Scenario *models.Scenario) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateScenario(ctx, Scenario)
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

// EnableScenario : ""
func (s *Service) EnableScenario(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableScenario(ctx, UniqueID)
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

//DisableScenario : ""
func (s *Service) DisableScenario(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableScenario(ctx, UniqueID)
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

//DeleteScenario : ""
func (s *Service) DeleteScenario(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteScenario(ctx, UniqueID)
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

// FilterScenario : ""
func (s *Service) FilterScenario(ctx *models.Context, filter *models.ScenarioFilter, pagination *models.Pagination) ([]models.RefScenario, error) {
	return s.Daos.FilterScenario(ctx, filter, pagination)

}
