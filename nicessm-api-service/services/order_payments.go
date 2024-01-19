package services

import (
	"errors"
	"fmt"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveOrderPayments :""
func (s *Service) SaveOrderPayment(ctx *models.Context, payment *models.OrderPayment) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	payment.Status = constants.ORDERPAYMENTSTATUSACTIVE
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 OrderPayments.created")
	payment.Created = &created
	log.Println("b4 OrderPayments.created")
	var dberr error
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		order, err := s.Daos.GetSingleOrder(ctx, payment.OrderID.Hex())
		if err != nil {
			log.Println(err)
			return err
		}
		if order.Status == constants.ORDERSTATUSINIT {
			for _, v := range order.Items {
				s.Daos.UpdateProductQuantity(ctx, v.ProductID, v.Quantity)

			}
		}

		payments, err := s.Daos.GetpaymentbyOrderID(ctx, payment.OrderID.Hex())
		if err != nil {
			log.Println(err)
			return err
		}
		var paidAmts float64
		for _, v := range payments {
			paidAmts = paidAmts + v.Amount
		}
		fmt.Println((paidAmts + payment.Amount), ">", order.Total)
		//overpayed
		if (paidAmts + payment.Amount) > order.Total {
			return errors.New("overpayment")

		}
		dberr = s.Daos.SaveOrderPayment(ctx, payment)
		if dberr != nil {
			log.Println(err)
			return dberr
		}
		if (paidAmts + payment.Amount) == order.Total {
			err := s.Daos.UpdateOrderPaymentStatus(ctx, payment.OrderID.Hex(), constants.ORDERSTATUSCOMPLETED, 0)
			if err != nil {
				log.Println(err)
				return err
			}
		} else {
			pendingamount := order.Total - (paidAmts + payment.Amount)
			err := s.Daos.UpdateOrderPaymentStatus(ctx, payment.OrderID.Hex(), constants.ORDERSTATUSPENDING, pendingamount)
			if err != nil {
				log.Println(err)
				return err
			}

		}

		if dberr != nil {
			if err1 := ctx.Session.CommitTransaction(sc); err1 != nil {
				log.Println("err in commit")
				return errors.New("Not able to commit - " + err1.Error())
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

//UpdateOrderPayments : ""
func (s *Service) UpdateOrderPayment(ctx *models.Context, payment *models.OrderPayment) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOrderPayment(ctx, payment)
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

//EnableOrderPayments : ""
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

//DisableOrderPayments : ""
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

//DeleteOrderPayments : ""
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

//GetSingleOrderPayments :""
func (s *Service) GetSingleOrderPayment(ctx *models.Context, UniqueID string) (*models.RefOrderPayment, error) {
	payment, err := s.Daos.GetSingleOrderPayment(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

//FilterOrderPayments :""
func (s *Service) FilterOrderPayment(ctx *models.Context, filter *models.OrderPaymentFilter, pagination *models.Pagination) (OrderPayments []models.RefOrderPayment, err error) {
	return s.Daos.FilterOrderPayment(ctx, filter, pagination)
}
