package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePropertyPenalty : ""
func (s *Service) SavePropertyVisitLogRemarkType(ctx *models.Context, remark *models.PropertyVisitLogRemarkType) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	remark.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYVISITLOGREMARKTYPE)
	remark.Status = constants.PROPERTYVISITLOGREMARKTYPESTATUSACTIVE
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	remark.Created = created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SavePropertyVisitLogRemarkType(ctx, remark)
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

//GetSinglePropertyVisitLogRemarkType :""
func (s *Service) GetSinglePropertyVisitLogRemarkType(ctx *models.Context, UniqueID string) (*models.RefPropertyVisitLogRemarkType, error) {
	remark, err := s.Daos.GetSinglePropertyVisitLogRemarkType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return remark, nil
}

// UpdatePropertyVisitLogRemarkType : ""
func (s *Service) UpdatePropertyVisitLogRemarkType(ctx *models.Context, remark *models.PropertyVisitLogRemarkType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.UpdatePropertyVisitLogRemarkType(ctx, remark)
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

// EnablePropertyVisitLogRemarkType : ""
func (s *Service) EnablePropertyVisitLogRemarkType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.EnablePropertyVisitLogRemarkType(ctx, UniqueID)
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

//DisablePropertyVisitLogRemarkType : ""
func (s *Service) DisablePropertyVisitLogRemarkType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.DisablePropertyVisitLogRemarkType(ctx, UniqueID)
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

// DeletePropertyVisitLogRemarkType : ""
func (s *Service) DeletePropertyVisitLogRemarkType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		dberr := s.Daos.DeletePropertyVisitLogRemarkType(ctx, UniqueID)
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

// FilterPropertyPenalty : ""
func (s *Service) FilterPropertyVisitLogRemarkType(ctx *models.Context, filter *models.PropertyVisitLogRemarkTypeFilter, pagination *models.Pagination) ([]models.RefPropertyVisitLogRemarkType, error) {
	return s.Daos.FilterPropertyVisitLogRemarkType(ctx, filter, pagination)

}
