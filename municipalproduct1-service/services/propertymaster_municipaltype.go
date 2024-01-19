package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveMunicipalType :""
func (s *Service) SaveMunicipalType(ctx *models.Context, municipalType *models.MunicipalType) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	municipalType.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONMUNICIPALTYPES)
	municipalType.Status = constants.MUNICIPALTYPESTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 municipalType.created")
	municipalType.Created = created
	log.Println("b4 municipalType.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveMunicipalType(ctx, municipalType)
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

//UpdateMunicipalType : ""
func (s *Service) UpdateMunicipalType(ctx *models.Context, municipalType *models.MunicipalType) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateMunicipalType(ctx, municipalType)
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

//EnableMunicipalType : ""
func (s *Service) EnableMunicipalType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableMunicipalType(ctx, UniqueID)
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

//DisableMunicipalType : ""
func (s *Service) DisableMunicipalType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableMunicipalType(ctx, UniqueID)
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

//DeleteMunicipalType : ""
func (s *Service) DeleteMunicipalType(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteMunicipalType(ctx, UniqueID)
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

//GetSingleMunicipalType :""
func (s *Service) GetSingleMunicipalType(ctx *models.Context, UniqueID string) (*models.RefMunicipalType, error) {
	municipalType, err := s.Daos.GetSingleMunicipalType(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return municipalType, nil
}

//FilterMunicipalType :""
func (s *Service) FilterMunicipalType(ctx *models.Context, municipalTypefilter *models.MunicipalTypeFilter, pagination *models.Pagination) (municipalType []models.RefMunicipalType, err error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.FilterMunicipalType(ctx, municipalTypefilter, pagination)
}

//GetSelectableMunicipalType : ""
func (s *Service) GetSelectableMunicipalType(ctx *models.Context) (*models.MunicipalType, error) {
	defer ctx.Session.EndSession(ctx.CTX)

	return s.Daos.GetSelectableMunicipalType(ctx)
}
