package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePropertyType : ""
func (s *Service) SavePropertyType(ctx *models.Context, propertytype *models.PropertyType) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	propertytype.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYTYPE)
	propertytype.Status = constants.PROPERTYTYPESTATUSACTIVE
	t := time.Now()
	//	PropertyType.RegisterDate = &t
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 PropertyType.created")
	//PropertyType.Created = &created
	log.Println("b4 PropertyType.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePropertyType(ctx, propertytype)
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

// GetSinglePropertyType : ""
func (s *Service) GetSinglePropertyType(ctx *models.Context, UniqueID string) (*models.RefPropertyType, error) {
	propertytype, err := s.Daos.GetSinglePropertyType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return propertytype, nil
}

// func (s *Service) GetSinglePropertyTypeWithHoldingNumber(ctx *models.Context, holdingNumber string) (*models.RefPropertyType, error) {
// 	PropertyType, err := s.Daos.GetSinglePropertyTypeWithHoldingNumber(ctx, holdingNumber)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return PropertyType, nil
// }

//UpdatePropertyType : ""
func (s *Service) UpdatePropertyType(ctx *models.Context, propertytype *models.PropertyType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyType(ctx, propertytype)
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

// EnablePropertyType : ""
func (s *Service) EnablePropertyType(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnablePropertyType(ctx, uniqueID)
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

// DisablePropertyType : ""
func (s *Service) DisablePropertyType(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisablePropertyType(ctx, uniqueID)
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

//DeletePropertyType : ""
func (s *Service) DeletePropertyType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePropertyType(ctx, UniqueID)
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

// FilterPropertyType : ""
func (s *Service) FilterPropertyType(ctx *models.Context, propertytype *models.FilterPropertyType, pagination *models.Pagination) (PropertyTypes []models.RefPropertyType, err error) {
	return s.Daos.FilterPropertyType(ctx, propertytype, pagination)
}
