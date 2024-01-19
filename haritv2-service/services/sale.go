package services

import (
	"errors"
	"haritv2-service/constants"
	"haritv2-service/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//SaveFarmerCart :""
func (s *Service) SaveSale(ctx *models.Context, farmerCart *models.Sale) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	farmerCart.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONSALE)
	farmerCart.Status = constants.FARMERCARTSTATUSACTIVE
	farmerCart.CompanyID = farmerCart.Company.ID
	t := time.Now()
	farmerCart.Date = &t
	// created := models.CreatedV2{}
	// created.On = &t
	// created.By = constants.SYSTEM
	// log.Println("b4 FarmerCart.created")
	// farmerCart.Created = &created
	log.Println("b4 FarmerCart.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveSale(ctx, farmerCart)
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
func (s *Service) FilterSale(ctx *models.Context, filter *models.SaleFilter, pagination *models.Pagination) ([]models.RefSale, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterSale(ctx, filter, pagination)

}
