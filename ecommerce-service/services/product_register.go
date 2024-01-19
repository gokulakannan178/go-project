package services

import (
	"ecommerce-service/constants"
	"ecommerce-service/models"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveRegisterProduct : ""
func (s *Service) SaveRegisterProduct(ctx *models.Context, rgProduct *models.RegisterProduct) error {

	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)

	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		if rgProduct.Product.UniqueID == "" {
			rgProduct.Product.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONPRODUCT)
			rgProduct.Product.Status = constants.PRODUCTSTATUSACTIVE
			dberr := s.Daos.SaveProduct(ctx, &rgProduct.Product)
			if dberr != nil {
				return dberr
			}
		}

		// for _, v := range rgProduct.Varients {
		// 	v.VendorID = rgProduct.Vendor.UniqueID
		// 	v.ProductID = rgProduct.Product.UniqueID
		// }
		// err := s.Daos.UpsertProductVarients(ctx, &rgProduct.Varients)
		// if err != nil {
		// 	return err
		// }

		for k, v := range rgProduct.InventoryData {
			v.Inventory.UniqueID = k + "_" + rgProduct.Vendor.UniqueID
			//	v.Inventory.PVCombination = k
			v.Inventory.ProductVarientID = rgProduct.Product.UniqueID
			v.Inventory.VendorID = rgProduct.Vendor.UniqueID
			v.Inventory.Status = constants.INVENTORYSTATUSACTIVE
			err := s.Daos.UpsertInventory(ctx, &v.Inventory)
			if err != nil {
				return err
			}
			inventoryMesh := []models.InventoryMesh{}
			for _, v2 := range v.InventoryMesh {
				v2.ProductID = rgProduct.Product.UniqueID
				v2.VendorID = rgProduct.Vendor.UniqueID
				v2.InventoryID = v.Inventory.UniqueID

				inventoryMesh = append(inventoryMesh, v2.InventoryMesh)
			}

			err = s.Daos.SaveManyInventoryMesh(ctx, inventoryMesh)
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
