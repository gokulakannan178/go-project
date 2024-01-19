package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveUserLoginLog :""
func (s *Service) SaveUserLoginLog(ctx *models.Context, UserLoginLog *models.UserLoginLog) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	UserLoginLog.Status = constants.USERLOGINLOGSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 UserLoginLog.created")
	UserLoginLog.Created = &created
	log.Println("b4 UserLoginLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveUserLoginLog(ctx, UserLoginLog)
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

//UpdateUserLoginLog : ""
func (s *Service) UpdateUserLoginLog(ctx *models.Context, UserLoginLog *models.UserLoginLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUserLoginLog(ctx, UserLoginLog)
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

//EnableUserLoginLog : ""
func (s *Service) EnableUserLoginLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableUserLoginLog(ctx, UniqueID)
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

//DisableUserLoginLog : ""
func (s *Service) DisableUserLoginLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableUserLoginLog(ctx, UniqueID)
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

//DeleteUserLoginLog : ""
func (s *Service) DeleteUserLoginLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteUserLoginLog(ctx, UniqueID)
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

//GetSingleUserLoginLog :""
func (s *Service) GetSingleUserLoginLog(ctx *models.Context, UniqueID string) (*models.RefUserLoginLog, error) {
	UserLoginLog, err := s.Daos.GetSingleUserLoginLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return UserLoginLog, nil
}

//FilterUserLoginLog :""
func (s *Service) FilterUserLoginLog(ctx *models.Context, UserLoginLogfilter *models.UserLoginLogFilter, pagination *models.Pagination) (UserLoginLog []models.RefUserLoginLog, err error) {
	return s.Daos.FilterUserLoginLog(ctx, UserLoginLogfilter, pagination)
}
func (s *Service) UserLogin(ctx *models.Context, UserLoginLog *models.UserLoginLog) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	UserLoginLog.Status = constants.USERLOGINLOGSTATUSLOGIN
	t := time.Now()
	UserLoginLog.LoginTime = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 UserLoginLog.created")
	UserLoginLog.Created = &created
	log.Println("b4 UserLoginLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		userid, _ := s.Daos.GetSingleUserLoginLogWithStatus(ctx, UserLoginLog.UserId.Hex())
		if userid != nil {
			err := s.Daos.LogoutUserLoginLog(ctx, userid.ID.Hex())
			if err != nil {
				return err
			}
		}
		dberr := s.Daos.SaveUserLoginLog(ctx, UserLoginLog)
		if dberr != nil {
			return errors.New("Db Error" + dberr.Error())
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
func (s *Service) UpdateUserLogout(ctx *models.Context, UserLoginLog *models.UserLoginLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		userid, _ := s.Daos.GetSingleUserLoginLogWithStatus(ctx, UserLoginLog.UserId.Hex())
		if userid != nil {
			err := s.Daos.LogoutUserLoginLog(ctx, userid.ID.Hex())
			if err != nil {
				return err
			}
		}
		if userid == nil {
			return errors.New("pls login")
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
