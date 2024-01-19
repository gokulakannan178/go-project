package service

import (
	"errors"
	"lgf-ccc-service/constants"
	"lgf-ccc-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveIdentityType : ""
func (s *Service) SaveIdentityType(ctx *models.Context, identitytype *models.IdentityType) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	identitytype.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONIDENTITYTYPE)
	identitytype.Status = constants.IDENTITYTYPESTATUSACTIVE
	t := time.Now()
	//	IdentityType.RegisterDate = &t
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 IdentityType.created")
	//IdentityType.Created = &created
	log.Println("b4 IdentityType.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveIdentityType(ctx, identitytype)
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

// GetSingleIdentityType : ""
func (s *Service) GetSingleIdentityType(ctx *models.Context, UniqueID string) (*models.RefIdentityType, error) {
	identitytype, err := s.Daos.GetSingleIdentityType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return identitytype, nil
}

// func (s *Service) GetSingleIdentityTypeWithHoldingNumber(ctx *models.Context, holdingNumber string) (*models.RefIdentityType, error) {
// 	IdentityType, err := s.Daos.GetSingleIdentityTypeWithHoldingNumber(ctx, holdingNumber)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return IdentityType, nil
// }

//UpdateIdentityType : ""
func (s *Service) UpdateIdentityType(ctx *models.Context, identitytype *models.IdentityType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateIdentityType(ctx, identitytype)
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

// EnableIdentityType : ""
func (s *Service) EnableIdentityType(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.EnableIdentityType(ctx, uniqueID)
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

// DisableIdentityType : ""
func (s *Service) DisableIdentityType(ctx *models.Context, uniqueID string) error {

	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		debrr := s.Daos.DisableIdentityType(ctx, uniqueID)
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

//DeleteIdentityType : ""
func (s *Service) DeleteIdentityType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteIdentityType(ctx, UniqueID)
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

// FilterIdentityType : ""
func (s *Service) FilterIdentityType(ctx *models.Context, identitytype *models.FilterIdentityType, pagination *models.Pagination) (IdentityTypes []models.RefIdentityType, err error) {
	return s.Daos.FilterIdentityType(ctx, identitytype, pagination)
}
