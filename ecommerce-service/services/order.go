package services

import (
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) InitiateOrder(ctx *models.Context, co *models.CreateOrder) (*models.Order, error) {
	order := new(models.Order)
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		if order.UniqueID == "" {
			order.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONORDER)

		}
		order.Status = constants.ORDERSTATUSINIT
		t := time.Now()
		order.Date = &t
		switch co.Company.Type {
		case constants.MODULEVENDOR:
			vendor, err := s.Daos.GetSingleVendor(ctx, co.Company.ID)
			if err != nil {
				return errors.New("error getting company-vendor-" + err.Error())
			}
			order.From.ID = co.Company.ID
			order.From.Name = vendor.Name
			order.From.Email = vendor.EmailID
			order.From.Moblie = vendor.MobileNo
			order.From.Address = vendor.Address
			order.From.Logo = vendor.Logo
			order.From.Type = constants.MODULEVENDOR
		default:
			return errors.New("invalid company type")
		}
		switch co.Customer.Type {
		case constants.MODULEVENDOR:
			vendor, err := s.Daos.GetSingleVendor(ctx, co.Customer.ID)
			if err != nil {
				return errors.New("error getting customer-vendor-" + err.Error())
			}
			order.To.ID = co.Customer.ID
			order.To.Name = vendor.Name
			order.To.Email = vendor.EmailID
			order.To.Moblie = vendor.MobileNo
			order.To.Address = vendor.Address
			order.To.Logo = vendor.Logo
			order.To.Type = constants.MODULEVENDOR
			order.BillingAddress.Address = vendor.Address
			order.ShippingAddress.Address = vendor.Address
		case constants.MODULECUSTOMER:
			customer, err := s.Daos.GetSingleGetUsingMobileNumber(ctx, co.Customer.ID)
			if err != nil {
				return errors.New("error getting customer-customer-" + err.Error())
			}
			if customer == nil {
				tempcustomer := new(models.Customer)
				tempcustomer.Mobile = co.Customer.ID
				tempcustomer.Email = co.Customer.Email
				tempcustomer.Name = co.Customer.Name
				tempcustomer.Address = co.Customer.Address
				tempcustomer.BelongsTo = co.Company.ID
				tempcustomer.BelongsToType = constants.MODULEVENDOR

				err := s.SaveCustomerWithOutTransaction(ctx, tempcustomer)
				if err != nil {
					return errors.New("error in saving customer" + err.Error())
				}
				customer, err = s.Daos.GetSingleCustomer(ctx, tempcustomer.UniqueID)
				if err != nil {
					return errors.New("error in getting save customer" + err.Error())

				}
			}
			order.To.Name = customer.Name
			order.To.ID = customer.Mobile
			order.To.Email = customer.Email
			order.To.Moblie = customer.Mobile
			order.To.Address = customer.Address
			order.To.Logo = customer.Photo
			order.To.Type = constants.MODULECUSTOMER
			order.ShippingAddress.Address = customer.Address
			order.BillingAddress.Address = customer.Address
		}
		for _, v := range co.Products {
			var item models.SaleItem
			item.Quantity = v.Quantity
			inventory, err := s.Daos.GetSingleInventory(ctx, v.InventoryID)
			if err != nil {

			}
			if inventory == nil {
				order.InitateRemark = append(order.InitateRemark, fmt.Sprintf("%v-item not found", v.InventoryID))
				continue

			}
			item.Item = *inventory
			if inventory.Quantity == 0 {
				order.InitateRemark = append(order.InitateRemark, fmt.Sprintf("%v-insufficient stock", inventory.Ref.Product.Name))
				continue
			}
			if inventory.Quantity < v.Quantity {

				order.InitateRemark = append(order.InitateRemark, fmt.Sprintf("the quantity of %v is reduced to %v due to insufficient of stock", inventory.Ref.Product.Name, inventory.Quantity))
				item.Quantity = inventory.Quantity
			}
			item.Total = item.Quantity * inventory.Price.Selling
			order.Item = append(order.Item, item)
			order.SubTotal = order.SubTotal + item.Total
		}
		order.Discount.Total = 0
		order.Total = order.SubTotal - order.Discount.Total
		if err := s.Daos.UpsertOrder(ctx, order); err != nil {
			return errors.New("Error in saving order - " + err.Error())
		}

		if err := ctx.Session.CommitTransaction(sc); err != nil {
			return errors.New("Not able to commit - " + err.Error())
		}
		return nil
	}); err != nil {
		log.Println("Transaction start aborting")
		if abortError := ctx.Session.AbortTransaction(ctx.CTX); abortError != nil {
			return nil, errors.New("Error while aborting transaction" + abortError.Error())
		}
		log.Println("Transaction aborting completed successfully")
		return nil, err
	}
	return order, nil
}

// GetSingleOrder : ""
func (s *Service) GetSingleOrder(ctx *models.Context, UniqueID string) (*models.RefOrder, error) {
	order, err := s.Daos.GetSingleOrder(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

//FilterProduct :""
func (s *Service) FilterOrder(ctx *models.Context, filter *models.OrderFilter, pagination *models.Pagination) ([]models.RefOrder, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterOrder(ctx, filter, pagination)

}
