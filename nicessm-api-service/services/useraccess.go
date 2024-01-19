package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUserAcl :""
func (s *Service) SaveUserAcl(ctx *models.Context, useracl *models.UserAcl) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	useracl.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERACL)

	useracl.Status = constants.USERACLSTATUSACTIVE
	//	useracl.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 useracl.created")
	useracl.Created = &created
	log.Println("b4 useracl.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveUserAcl(ctx, useracl)
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

//UpdateUserAcl : ""
func (s *Service) UpdateUserAcl(ctx *models.Context, useracl *models.UserAcl) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUserAcl(ctx, useracl)
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

//EnableUserAcl : ""
func (s *Service) EnableUserAcl(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableUserAcl(ctx, UniqueID)
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

//DisableUserAcl : ""
func (s *Service) DisableUserAcl(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableUserAcl(ctx, UniqueID)
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

//DeleteUserAcl : ""
func (s *Service) DeleteUserAcl(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteUserAcl(ctx, UniqueID)
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

//GetSingleUserAcl :""
func (s *Service) GetSingleUserAcl(ctx *models.Context, UniqueID string) (*models.RefUserAcl, error) {
	useracl, err := s.Daos.GetSingleUserAcl(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return useracl, nil
}

//FilterUserAcl :""
func (s *Service) FilterUserAcl(ctx *models.Context, useraclfilter *models.UserAclFilter, pagination *models.Pagination) (useracl []models.RefUserAcl, err error) {
	return s.Daos.FilterUserAcl(ctx, useraclfilter, pagination)
}
func (s *Service) SaveUserAclWithUpsert(ctx *models.Context, useracl *models.UserAcl) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	useracl.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERACL)

	useracl.Status = constants.USERACLSTATUSACTIVE
	//	useracl.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 useracl.created")
	useracl.Created = &created
	log.Println("b4 useracl.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveUserAclWithUpsert(ctx, useracl)
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
func (s *Service) GetSingleUserAclWithUserName(ctx *models.Context, UniqueID string) (*models.RefUserAcl, error) {
	useracl, err := s.Daos.GetSingleUserAclWithUserName(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return useracl, nil
}
