package services

import (
	"errors"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveOrder :""
func (s *Service) SaveOrder(ctx *models.Context, order *models.Order) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

	order.Status = constants.ORDERSTATUSACTIVE
	order.ActiveStatus = true
	t := time.Now()
	created := models.Created{}
	created.On = &t
	created.By = constants.SYSTEM
	log.Println("b4 Order.created")
	order.Created = &created
	order.Date = t
	log.Println("b4 Order.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveOrder(ctx, order)
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

//UpdateOrder : ""
func (s *Service) UpdateOrder(ctx *models.Context, order *models.Order) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateOrder(ctx, order)
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

//EnableOrder : ""
func (s *Service) EnableOrder(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableOrder(ctx, UniqueID)
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

//DisableOrder : ""
func (s *Service) DisableOrder(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableOrder(ctx, UniqueID)
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

//DeleteOrder : ""
func (s *Service) DeleteOrder(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteOrder(ctx, UniqueID)
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

//GetSingleOrder :""
func (s *Service) GetSingleOrder(ctx *models.Context, UniqueID string) (*models.RefOrder, error) {
	order, err := s.Daos.GetSingleOrder(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

//FilterOrder :""
func (s *Service) FilterOrder(ctx *models.Context, filter *models.OrderFilter, pagination *models.Pagination) (Order []models.RefOrder, err error) {
	return s.Daos.FilterOrder(ctx, filter, pagination)
}
func (s *Service) CreateOrder(ctx *models.Context, order *models.Order) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		//organisation.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORGANISATION)

		order.Status = constants.ORDERSTATUSINIT
		order.ActiveStatus = true
		t := time.Now()
		created := models.Created{}
		created.On = &t
		created.By = constants.SYSTEM
		log.Println("b4 Order.created")
		order.Created = &created
		order.Date = t
		log.Println("b4 Order.created")
		switch order.From.Type {
		case constants.DEALER:
			dealer, err := s.Daos.GetSingleDealer(ctx, string(order.From.ID.Hex()))
			if err != nil {
				return err
			}
			if dealer == nil {
				return errors.New("dealer not available")
			}
			order.From.Name = dealer.Name
			order.From.Mobile = dealer.Mobile
			order.From.Email = dealer.Email
			order.From.Block = dealer.Block
			order.From.State = dealer.State
			order.From.Village = dealer.Village
			order.From.GramPanchayat = dealer.GramPanchayat
			order.From.District = dealer.District
			order.From.PinCode = dealer.PinCode

		case constants.FARMER:
			farmer, err := s.Daos.GetSingleFarmer(ctx, order.From.ID.Hex())
			if err != nil {
				return err
			}
			if farmer == nil {
				return errors.New("farmer not available")
			}
			order.To.Name = farmer.Name
			order.To.Email = farmer.Email
			order.To.Mobile = farmer.MobileNumber
			order.To.Block = farmer.Block
			order.To.State = farmer.State
			order.To.Village = farmer.Village
			order.To.GramPanchayat = farmer.GramPanchayat
			order.To.District = farmer.District
			order.To.PinCode = farmer.PinCode

		}
		if len(order.Items) < 0 {
			return errors.New("items required to payment.please check the items")
		}
		for k, v := range order.Items {
			product, err := s.Daos.GetSingleProduct(ctx, v.ProductID.Hex())
			if err != nil {
				return err
			}
			if product == nil {
				return errors.New("product not available")
			}
			order.Items[k].Name = product.Name
			order.Items[k].BuyingPrice = product.BuyingPrice
			order.Items[k].SellingPrice = product.SellingPrice
			order.Items[k].Amount = product.SellingPrice * order.Items[k].Quantity
			order.Subtotal = order.Subtotal + order.Items[k].Amount
		}
		order.Discount = 0
		order.Tax = 0
		order.Total = order.Subtotal + order.Discount + order.Tax
		dberr := s.Daos.CreateOrder(ctx, order)
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
