package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveTradelicensePayeeNameChange :""
func (s *Service) SaveTradelicensePayeeNameChange(ctx *models.Context, ppnc *models.TradelicensePayeeNameChange) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	ppnc.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONTRADELICENSEPAYEENAMEHANGE)
	ppnc.Status = constants.TRADELICENSEPAYEENAMECHANGESTATUSPENDING
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 TradelicensePayeeNameChange.created")
	ppnc.CreatedOn = created
	log.Println("b4 TradelicensePayeeNameChange.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		txnId, err := s.Daos.GetSingleTradeLicensePaymentWithTxtID(ctx, ppnc.TransactionId)
		if err != nil {
			return err
		}
		ppnc.PreviousPayeeName = txnId.Details.PayeeName
		ppnc.TransactionId = txnId.TnxID
		ppnc.TradelicenseId = txnId.PropertyID
		ppnc.ReceiptNo = txnId.ReciptNo
		dberr := s.Daos.SaveTradelicensePayeeNameChange(ctx, ppnc)
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

// GetSingleTradelicensePayeeNameChange :""
func (s *Service) GetSingleTradelicensePayeeNameChange(ctx *models.Context, UniqueID string) (*models.RefTradelicensePayeeNameChange, error) {
	ppnc, err := s.Daos.GetSingleTradelicensePayeeNameChange(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ppnc, nil
}

// UpdateTradelicensePayeeNameChange : ""
func (s *Service) UpdateTradelicensePayeeNameChange(ctx *models.Context, ppnc *models.TradelicensePayeeNameChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateTradelicensePayeeNameChange(ctx, ppnc)
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

// EnableTradelicensePayeeNameChange : ""
func (s *Service) EnableTradelicensePayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableTradelicensePayeeNameChange(ctx, UniqueID)
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

// DisableTradelicensePayeeNameChange : ""
func (s *Service) DisableTradelicensePayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableTradelicensePayeeNameChange(ctx, UniqueID)
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

// DeleteTradelicensePayeeNameChange : ""
func (s *Service) DeleteTradelicensePayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteTradelicensePayeeNameChange(ctx, UniqueID)
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

// FilterTradelicensePayeeNameChange :""
func (s *Service) FilterTradelicensePayeeNameChange(ctx *models.Context, filter *models.TradelicensePayeeNameChangeFilter, pagination *models.Pagination) (ppnc []models.RefTradelicensePayeeNameChange, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterTradelicensePayeeNameChange(ctx, filter, pagination)
}

// ApproveTradeLicense : ""
func (s *Service) ApproveTradelicensePayeeNameChange(ctx *models.Context, approve *models.ApproveTradelicensePayeeNameChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()

		approve.On = &t
		ppnc, err := s.Daos.GetSingleTradelicensePayeeNameChange(ctx, approve.UniqueID)
		if err != nil {
			return err
		}
		if ppnc != nil {
			err = s.Daos.UpdateTradelicensePayeenamewithTxnId(ctx, ppnc.TransactionId, ppnc.ChangeData.Name)
			if err != nil {
				return err
			}
		}

		err = s.Daos.ApproveTradelicensePayeeNameChange(ctx, approve)
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
func (s *Service) NotApproveTradelicensePayeeNameChange(ctx *models.Context, notApprove *models.NotApproveTradeLicense) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		notApprove.On = &t

		err := s.Daos.NotApproveTradelicensePayeeNameChange(ctx, notApprove)
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
