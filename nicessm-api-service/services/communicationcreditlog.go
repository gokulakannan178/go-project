package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveCommunicationCreditLog :""
func (s *Service) SaveCommunicationCreditLog(ctx *models.Context, CommunicationCreditLog *models.CommunicationCreditLog) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	CommunicationCreditLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCOMMUNICATIONCREDITLOG)

	CommunicationCreditLog.Status = constants.COMMUNICATIONCREDITLOGSTATUSACTIVE

	t := time.Now()
	CommunicationCreditLog.CreateData = &t

	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 CommunicationCreditLog.created")
	CommunicationCreditLog.Created = &created
	log.Println("b4 CommunicationCreditLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveCommunicationCreditLog(ctx, CommunicationCreditLog)
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

//UpdateCommunicationCreditLog : ""
func (s *Service) UpdateCommunicationCreditLog(ctx *models.Context, CommunicationCreditLog *models.CommunicationCreditLog) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateCommunicationCreditLog(ctx, CommunicationCreditLog)
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

//EnableCommunicationCreditLog : ""
func (s *Service) EnableCommunicationCreditLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableCommunicationCreditLog(ctx, UniqueID)
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

//DisableCommunicationCreditLog : ""
func (s *Service) DisableCommunicationCreditLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableCommunicationCreditLog(ctx, UniqueID)
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

//DeleteCommunicationCreditLog : ""
func (s *Service) DeleteCommunicationCreditLog(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteCommunicationCreditLog(ctx, UniqueID)
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

//GetSingleCommunicationCreditLog :""
func (s *Service) GetSingleCommunicationCreditLog(ctx *models.Context, UniqueID string) (*models.RefCommunicationCreditLog, error) {
	CommunicationCreditLog, err := s.Daos.GetSingleCommunicationCreditLog(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return CommunicationCreditLog, nil
}

//FilterCommunicationCreditLog :""
func (s *Service) FilterCommunicationCreditLog(ctx *models.Context, CommunicationCreditLogfilter *models.CommunicationCreditLogFilter, pagination *models.Pagination) (CommunicationCreditLog []models.RefCommunicationCreditLog, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterCommunicationCreditLog(ctx, CommunicationCreditLogfilter, pagination)

}
func (s *Service) UpdateCommunicationCreditLogWithPostCredit(ctx *models.Context, CommunicationCreditLog *models.CommunicationCreditLog) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	CommunicationCreditLog.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCOMMUNICATIONCREDITLOG)

	CommunicationCreditLog.Status = constants.COMMUNICATIONCREDITLOGSTATUSACTIVE

	t := time.Now()
	CommunicationCreditLog.CreateData = &t
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 CommunicationCreditLog.created")
	CommunicationCreditLog.Created = &created
	log.Println("b4 CommunicationCreditLog.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		precredit, err := s.Daos.GetSingleCommunicationCreditLogWithMode(ctx, CommunicationCreditLog.CommunicationMode)
		if err != nil {
			return err
		}
		if precredit != nil {
			CommunicationCreditLog.PreCredit = precredit.PostCredit
			CommunicationCreditLog.PostCredit = precredit.PostCredit + CommunicationCreditLog.Credit
		}
		if precredit == nil {
			CommunicationCreditLog.PreCredit = 0
			CommunicationCreditLog.PostCredit = CommunicationCreditLog.Credit

		}
		err = s.Daos.UpdateCommunicationCreditWithBalance(ctx, CommunicationCreditLog.CommunicationMode, CommunicationCreditLog.Credit)
		if err != nil {
			return err
		}
		dberr := s.Daos.SaveCommunicationCreditLog(ctx, CommunicationCreditLog)
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
