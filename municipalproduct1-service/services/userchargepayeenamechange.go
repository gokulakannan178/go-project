package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveUserchargePayeeNameChange :""
func (s *Service) SaveUserchargePayeeNameChange(ctx *models.Context, ppnc *models.UserchargePayeeNameChange) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	ppnc.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONUSERCHARGEPAYEENAMEHANGE)
	ppnc.Status = constants.USERCHARGEPAYEENAMECHANGESTATUSPENDING
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 UserchargePayeeNameChange.created")
	ppnc.CreatedOn = created
	log.Println("b4 UserchargePayeeNameChange.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		txnId, err := s.Daos.GetSingleUserChargePaymentWithTxtID(ctx, ppnc.TransactionId)
		if err != nil {
			return err
		}
		ppnc.PreviousPayeeName = txnId.Details.PayeeName
		ppnc.TransactionId = txnId.TnxID
		ppnc.UserchargeId = txnId.PropertyID
		ppnc.ReceiptNo = txnId.ReciptNo
		dberr := s.Daos.SaveUserchargePayeeNameChange(ctx, ppnc)
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

// GetSingleUserchargePayeeNameChange :""
func (s *Service) GetSingleUserchargePayeeNameChange(ctx *models.Context, UniqueID string) (*models.RefUserchargePayeeNameChange, error) {
	ppnc, err := s.Daos.GetSingleUserchargePayeeNameChange(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ppnc, nil
}

// UpdateUserchargePayeeNameChange : ""
func (s *Service) UpdateUserchargePayeeNameChange(ctx *models.Context, ppnc *models.UserchargePayeeNameChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateUserchargePayeeNameChange(ctx, ppnc)
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

// EnableUserchargePayeeNameChange : ""
func (s *Service) EnableUserchargePayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableUserchargePayeeNameChange(ctx, UniqueID)
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

// DisableUserchargePayeeNameChange : ""
func (s *Service) DisableUserchargePayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableUserchargePayeeNameChange(ctx, UniqueID)
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

// DeleteUserchargePayeeNameChange : ""
func (s *Service) DeleteUserchargePayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteUserchargePayeeNameChange(ctx, UniqueID)
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

// FilterUserchargePayeeNameChange :""
func (s *Service) FilterUserchargePayeeNameChange(ctx *models.Context, filter *models.UserchargePayeeNameChangeFilter, pagination *models.Pagination) (ppnc []models.RefUserchargePayeeNameChange, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterUserchargePayeeNameChange(ctx, filter, pagination)
}

// ApproveTradeLicense : ""
func (s *Service) ApproveUserchargePayeeNameChange(ctx *models.Context, approve *models.ApproveUserchargePayeeNameChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()

		approve.On = &t
		ppnc, err := s.Daos.GetSingleUserchargePayeeNameChange(ctx, approve.UniqueID)
		if err != nil {
			return err
		}
		if ppnc != nil {
			err = s.Daos.UpdateUserchargePayeenamewithTxnId(ctx, ppnc.TransactionId, ppnc.ChangeData.Name)
			if err != nil {
				return err
			}
		}

		err = s.Daos.ApproveUserchargePayeeNameChange(ctx, approve)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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

// NotApproveTradeLicense : ""
func (s *Service) NotApproveUserchargePayeeNameChange(ctx *models.Context, notApprove *models.NotApproveTradeLicense) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		notApprove.On = &t

		err := s.Daos.NotApproveUserchargePayeeNameChange(ctx, notApprove)
		if err != nil {
			return nil
		}
		if err = ctx.Session.CommitTransaction(sc); err != nil {
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
