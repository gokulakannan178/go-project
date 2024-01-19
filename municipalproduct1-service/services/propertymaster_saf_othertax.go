package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SavePropertyOtherTax :""
func (s *Service) SavePropertyOtherTax(ctx *models.Context, propertyOtherTax *models.PropertyOtherTax) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	propertyOtherTax.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYOTHERTAX)
	propertyOtherTax.Status = constants.PROPERTYOTHERSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 propertyOtherTax.created")
	propertyOtherTax.Created = created
	log.Println("b4 propertyOtherTax.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePropertyOtherTax(ctx, propertyOtherTax)
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

//UpdatePropertyOtherTax : ""
func (s *Service) UpdatePropertyOtherTax(ctx *models.Context, propertyOtherTax *models.PropertyOtherTax) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyOtherTax(ctx, propertyOtherTax)
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

//EnablePropertyOtherTax : ""
func (s *Service) EnablePropertyOtherTax(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePropertyOtherTax(ctx, UniqueID)
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

//DisablePropertyOtherTax : ""
func (s *Service) DisablePropertyOtherTax(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePropertyOtherTax(ctx, UniqueID)
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

//DeletePropertyOtherTax : ""
func (s *Service) DeletePropertyOtherTax(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePropertyOtherTax(ctx, UniqueID)
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

//GetSinglePropertyOtherTax :""
func (s *Service) GetSinglePropertyOtherTax(ctx *models.Context, UniqueID string) (*models.RefPropertyOtherTax, error) {
	propertyOtherTax, err := s.Daos.GetSinglePropertyOtherTax(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return propertyOtherTax, nil
}

//FilterPropertyOtherTax :""
func (s *Service) FilterPropertyOtherTax(ctx *models.Context, propertyOtherTaxfilter *models.PropertyOtherTaxFilter, pagination *models.Pagination) (propertyOtherTax []models.RefPropertyOtherTax, err error) {
	return s.Daos.FilterPropertyOtherTax(ctx, propertyOtherTaxfilter, pagination)
}
