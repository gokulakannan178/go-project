package services

import (
	"errors"
	"log"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveShoprentPayeeNameChange :""
func (s *Service) SaveShoprentPayeeNameChange(ctx *models.Context, ppnc *models.ShoprentPayeeNameChange) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	ppnc.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSHOPRENTPAYEENAMEHANGE)
	ppnc.Status = constants.SHOPRENTPAYEENAMECHANGESTATUSPENDING
	t := time.Now()
	created := models.CreatedV2{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 ShoprentPayeeNameChange.created")
	ppnc.CreatedOn = created
	log.Println("b4 ShoprentPayeeNameChange.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		txnId, err := s.Daos.GetSingleShoprentPaymentWithTxtID(ctx, ppnc.TransactionId)
		if err != nil {
			return err
		}
		ppnc.PreviousPayeeName = txnId.Details.PayeeName
		ppnc.TransactionId = txnId.TnxID
		ppnc.ShoprentId = txnId.ShopRentID
		ppnc.ReceiptNo = txnId.ReciptNo
		dberr := s.Daos.SaveShoprentPayeeNameChange(ctx, ppnc)
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

// GetSingleShoprentPayeeNameChange :""
func (s *Service) GetSingleShoprentPayeeNameChange(ctx *models.Context, UniqueID string) (*models.RefShoprentPayeeNameChange, error) {
	ppnc, err := s.Daos.GetSingleShoprentPayeeNameChange(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return ppnc, nil
}

// UpdateShoprentPayeeNameChange : ""
func (s *Service) UpdateShoprentPayeeNameChange(ctx *models.Context, ppnc *models.ShoprentPayeeNameChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateShoprentPayeeNameChange(ctx, ppnc)
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

// EnableShoprentPayeeNameChange : ""
func (s *Service) EnableShoprentPayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableShoprentPayeeNameChange(ctx, UniqueID)
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

// DisableShoprentPayeeNameChange : ""
func (s *Service) DisableShoprentPayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableShoprentPayeeNameChange(ctx, UniqueID)
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

// DeleteShoprentPayeeNameChange : ""
func (s *Service) DeleteShoprentPayeeNameChange(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteShoprentPayeeNameChange(ctx, UniqueID)
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

// FilterShoprentPayeeNameChange :""
func (s *Service) FilterShoprentPayeeNameChange(ctx *models.Context, filter *models.ShoprentPayeeNameChangeFilter, pagination *models.Pagination) (ppnc []models.RefShoprentPayeeNameChange, err error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterShoprentPayeeNameChange(ctx, filter, pagination)
}

// ApproveTradeLicense : ""
func (s *Service) ApproveShoprentPayeeNameChange(ctx *models.Context, approve *models.ApproveShoprentPayeeNameChange) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()

		approve.On = &t
		ppnc, err := s.Daos.GetSingleShoprentPayeeNameChange(ctx, approve.UniqueID)
		if err != nil {
			return err
		}
		if ppnc != nil {
			err = s.Daos.UpdateShoprentPayeenamewithTxnId(ctx, ppnc.TransactionId, ppnc.ChangeData.Name)
			if err != nil {
				return err
			}
		}

		err = s.Daos.ApproveShoprentPayeeNameChange(ctx, approve)
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
func (s *Service) NotApproveShoprentPayeeNameChange(ctx *models.Context, notApprove *models.NotApproveTradeLicense) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		t := time.Now()
		notApprove.On = &t

		err := s.Daos.NotApproveShoprentPayeeNameChange(ctx, notApprove)
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
