package services

import (
	"errors"
	"fmt"
	"log"
	"time"

	"ecommerce-service/constants"
	"ecommerce-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveOrderPayment : ""
func (s *Service) SaveOrderPayment(ctx *models.Context, OrderPayment *models.OrderPayment) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	OrderPayment.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORDERPAYMENT)
	OrderPayment.Status = constants.ORDERPAYMENTSTATUSACTIVE
	t := time.Now()

	created := new(models.Created)
	created.On = &t
	created.By = constants.SYSTEM
	OrderPayment.Created = *created
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOrderPayment(ctx, OrderPayment)
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

//GetSingleOrderPayment :""
func (s *Service) GetSingleOrderPayment(ctx *models.Context, UniqueID string) (*models.RefOrderPayment, error) {
	OrderPayment, err := s.Daos.GetSingleOrderPayment(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return OrderPayment, nil
}

// UpdateOrderPayment : ""
func (s *Service) UpdateOrderPayment(ctx *models.Context, OrderPayment *models.OrderPayment) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOrderPayment(ctx, OrderPayment)
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

// EnableOrderPayment : ""
func (s *Service) EnableOrderPayment(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableOrderPayment(ctx, UniqueID)
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

//DisableOrderPayment : ""
func (s *Service) DisableOrderPayment(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableOrderPayment(ctx, UniqueID)
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

//DeleteOrderPayment : ""
func (s *Service) DeleteOrderPayment(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOrderPayment(ctx, UniqueID)
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

// FilterOrderPayment : ""
func (s *Service) FilterOrderPayment(ctx *models.Context, filter *models.OrderPaymentFilter, pagination *models.Pagination) ([]models.RefOrderPayment, error) {
	return s.Daos.FilterOrderPayment(ctx, filter, pagination)

}

// MakePayment : ""
func (s *Service) MakePayment(ctx *models.Context, OrderPayment *models.OrderPayment) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	OrderPayment.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORDERPAYMENT)
	OrderPayment.Status = constants.ORDERPAYMENTSTATUSCOMPLETED
	t := time.Now()
	if OrderPayment.Date == nil {
		OrderPayment.Date = &t
	}
	created := new(models.Created)
	created.On = &t
	created.By = constants.SYSTEM
	OrderPayment.Created = *created
	OrderPayment.RecordDate = &t
	var totalamt float64
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if OrderPayment.OrderID == "" {
			return errors.New("orderid is empty")

		}
		orderpayments, err := s.Daos.GetCompletedOrderPayment(ctx, OrderPayment.OrderID)
		if err != nil {
			return err
		}
		var amt float64
		if len(orderpayments) > 0 {
			for _, v := range orderpayments {
				amt += v.Amount
				fmt.Println("amount", amt)

			}

		} else {
			fmt.Println("no orderpayments")
		}
		totalamt = OrderPayment.Amount + amt
		fmt.Println("total amount", totalamt)

		dberr := s.Daos.SaveOrderPayment(ctx, OrderPayment)
		if dberr != nil {
			return dberr
		}
		order, dberr := s.Daos.GetSingleOrder(ctx, OrderPayment.OrderID)
		if dberr != nil {
			return dberr
		}
		if order.Payment.Amount <= totalamt {
			err := s.Daos.UpdateOrderPaymentStatus(ctx, OrderPayment.OrderID, constants.ORDERSTATUSCOMPLETED, 0)
			if err != nil {
				return err
			}
		} else {
			pendingamount := order.Payment.Amount - totalamt
			err := s.Daos.UpdateOrderPaymentStatus(ctx, OrderPayment.OrderID, constants.ORDERSTATUSPENDING, pendingamount)
			if err != nil {
				return err
			}
		}
		for _, v := range order.Item {
			err := s.Daos.UpdateInventoryQuantityForSale(ctx, v.Item.UniqueID, v.Quantity)
			if err != nil {
				return err
			}
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
