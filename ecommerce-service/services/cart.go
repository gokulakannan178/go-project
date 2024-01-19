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

// SaveCart : ""
func (s *Service) AddCart(ctx *models.Context, cart *models.Cart) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	t := time.Now()
	created := new(models.CreatedV2)
	created.On = &t
	created.By = constants.SYSTEM
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		err := s.AddCartWithoutTransaction(ctx, cart)
		if err != nil {
			return err
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
func (s *Service) AddCartWithoutTransaction(ctx *models.Context, cart *models.Cart) error {
	cart2, err := s.Daos.GetSingleCartAndVendor(ctx, cart.Customer.ID, cart.Company.ID)
	if err != nil {
		return err
	}

	if cart2 == nil {
		cart.Status = constants.CARTSTATUSACTIVE
		cart.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONCART)
		dberr := s.Daos.AddCart(ctx, cart)
		if dberr != nil {
			return dberr
		}
		return nil
	}
	fmt.Println("cart2===>", cart2)
	if cart2 != nil {
		if cart2.Company.ID != cart.Company.ID {
			return errors.New("ConflictVendor")
		}
		var products []models.CreateOrderProduct
		for _, v := range cart.Products {
			//	found := false
			var v2 models.CreateOrderProduct
			for _, v2 = range cart2.Products {
				if v.InventoryID == v2.InventoryID {
					v2.Quantity = v.Quantity
					products = append(products, v)
					//found = true
					break

				}
				// if v.InventoryID != v2.InventoryID {
				// 	products = append(products, v)
				// }
				if v.InventoryID != v2.InventoryID {
					products = append(products, v2)

				}

			}
			if v.InventoryID != v2.InventoryID {
				products = append(products, v)

			}
		}
		cart2.Products = products
		err = s.Daos.UpdateCart(ctx, &cart2.Cart)
		if err != nil {
			return err
		}
		cart.UniqueID = cart2.UniqueID
		cart.Status = cart2.Status
	}
	return err
}

//GetSingleCart :""
func (s *Service) GetSingleCart(ctx *models.Context, UniqueID string) (*models.RefCart, error) {
	cart, err := s.Daos.GetSingleCart(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// UpdateCart : ""
func (s *Service) UpdateCart(ctx *models.Context, cart *models.Cart) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateCart(ctx, cart)
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

// EnableCart : ""
func (s *Service) EnableCart(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableCart(ctx, UniqueID)
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

//DisableCart : ""
func (s *Service) DisableCart(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableCart(ctx, UniqueID)
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

//DeleteCart : ""
func (s *Service) DeleteCart(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteCart(ctx, UniqueID)
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

// FilterCart : ""
func (s *Service) FilterCart(ctx *models.Context, filter *models.CartFilter, pagination *models.Pagination) ([]models.RefCart, error) {
	return s.Daos.FilterCart(ctx, filter, pagination)

}

// UpdateCartItemQuanity : ""
func (s *Service) UpdateCartItemQuanity(ctx *models.Context, updatecart *models.UpdateCart) (cart *models.RefCart, err error) {
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}

	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		cart, err = s.Daos.GetSingleCartWithCustomerId(ctx, updatecart.CustomerId)
		if err != nil {
			return err
		}
		if cart == nil {
			newCart := new(models.Cart)
			newCart.Customer.ID = updatecart.CustomerId
			newCart.Customer.Type = updatecart.CustomerType
			newCart.Company.ID = updatecart.VendorID
			newCart.Company.Type = updatecart.VendorType
			product := models.CreateOrderProduct{}
			product.InventoryID = updatecart.InventoryID
			product.Quantity = updatecart.Quantity
			newCart.Products = append(newCart.Products, product)
			err = s.AddCartWithoutTransaction(ctx, newCart)
			if err != nil {
				return err
			}
			//	updatecart.CartUniqueID = newCart.UniqueID
			return nil

		}
		err := s.Daos.UpdateCartItemQuanity(ctx, updatecart)
		if err != nil {
			return errors.New("Db Error" + err.Error())
		}

		cart, err = s.Daos.GetSingleCartWithInventoryId(ctx, updatecart.CustomerId, updatecart.InventoryID)
		if err != nil {
			return err
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

	cart, err = s.Daos.GetSingleCartWithInventoryId(ctx, updatecart.CustomerId, updatecart.InventoryID)
	if err != nil {
		return nil, err
	}
	if cart != nil {
		err = s.Daos.CalculationCart(ctx, &cart.Cart)
		if err != nil {
			return nil, err
		}

	}
	return cart, nil
}
