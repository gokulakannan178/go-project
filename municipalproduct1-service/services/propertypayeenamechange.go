package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SavePropertyPayeeNameChange :""
func (s *Service) SavePropertyPayeeNameChange(ctx *models.Context, ppnc *models.PropertyPayeeNameChange) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	ppnc.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPROPERTYPAYEENAMEHANGE)
	ppnc.Status = constants.PROPERTYPAYEENAMECHANGESTATUSPENDING
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 PropertyPayeeNameChange.created")
	ppnc.CreatedOn = created
	log.Println("b4 PropertyPayeeNameChange.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		txnId, err := s.Daos.GetSinglePropertyPaymentTxtID(ctx, ppnc.TransactionId)
		if err != nil {
			return err
		}
		ppnc.PreviousPayeeName = txnId.Details.PayeeName
		ppnc.TransactionId = txnId.TnxID
		ppnc.PropertyId = txnId.PropertyID
		ppnc.ReceiptNo = txnId.ReciptNo
		dberr := s.Daos.SavePropertyPayeeNameChange(ctx, ppnc)
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

// GetSinglePropertyPayeeNameChange :""
func (s *Service) GetSinglePropertyPayeeNameChange(ctx *models.Context, UniqueID string) (*models.RefPropertyPayeeNameChange, error) {
	ppnc, err := s.Daos.GetSinglePropertyPayeeNameChange(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ppnc, nil
}

// UpdatePropertyPayeeNameChange : ""
func (s *Service) UpdatePropertyPayeeNameChange(ctx *models.Context, ppnc *models.PropertyPayeeNameChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdatePropertyPayeeNameChange(ctx, ppnc)
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

// EnablePropertyPayeeNameChange : ""
func (s *Service) EnablePropertyPayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnablePropertyPayeeNameChange(ctx, UniqueID)
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

// DisablePropertyPayeeNameChange : ""
func (s *Service) DisablePropertyPayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisablePropertyPayeeNameChange(ctx, UniqueID)
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

// DeletePropertyPayeeNameChange : ""
func (s *Service) DeletePropertyPayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeletePropertyPayeeNameChange(ctx, UniqueID)
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

// FilterPropertyPayeeNameChange :""
func (s *Service) FilterPropertyPayeeNameChange(ctx *models.Context, filter *models.PropertyPayeeNameChangeFilter, pagination *models.Pagination) (ppnc []models.RefPropertyPayeeNameChange, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterPropertyPayeeNameChange(ctx, filter, pagination)
}

// ApproveTradeLicense : ""
func (s *Service) ApprovePropertyPayeeNameChange(ctx *models.Context, approve *models.ApprovePropertyPayeeNameChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()

		approve.On = &t
		ppnc, err := s.Daos.GetSinglePropertyPayeeNameChange(ctx, approve.UniqueID)
		if err != nil {
			return err
		}
		if ppnc != nil {
			err = s.Daos.UpdatePropertyPayeenamewithTxnId(ctx, ppnc.TransactionId, ppnc.ChangeData.Name)
			if err != nil {
				return err
			}
		}

		err = s.Daos.ApprovePropertyPayeeNameChange(ctx, approve)
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
func (s *Service) NotApprovePropertyPayeeNameChange(ctx *models.Context, notApprove *models.NotApproveTradeLicense) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		notApprove.On = &t

		err := s.Daos.NotApprovePropertyPayeeNameChange(ctx, notApprove)
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
