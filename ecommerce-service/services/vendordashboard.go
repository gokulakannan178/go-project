package services

import (
	"ecommerce-service/models"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// MakePayment : ""
func (s *Service) VendorDashboard(ctx *models.Context, filter *models.VendorDashBoardFilter) (*models.VendorDashBoard, error) {
	vendordashboard := new(models.VendorDashBoard)
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	log.Println("b4 vendor.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		order, dberr := s.Daos.VendorDashboardOrder(ctx, filter)
		if dberr != nil {
			return dberr
		}
		product, dberr := s.Daos.VendorDashboardProduct(ctx, filter)
		if dberr != nil {
			return dberr
		}
		lowstock, dberr := s.Daos.VendorDashboardLowStock(ctx, filter)
		if dberr != nil {
			return dberr
		}

		vendordashboard.Order = *order
		vendordashboard.Product = *product
		vendordashboard.LowStock = *lowstock
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
	return vendordashboard, nil

}
