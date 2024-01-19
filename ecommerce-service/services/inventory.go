package services

import (
	"ecommerce-service/constants"
	"errors"
	"log"

	"ecommerce-service/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveInventory : ""
func (s *Service) SaveInventory(ctx *models.Context, block *models.Inventory) error {
	log.Println("transaction start")
	// Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	block.UniqueID = s.Daos.GetUniqueID(ctx, constants.COLLECTIONINVENTORY)
	block.Status = constants.INVENTORYSTATUSACTIVE
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.UpsertInventory(ctx, block)
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

//GetSingleInventory :""
func (s *Service) GetSingleInventory(ctx *models.Context, UniqueID string) (*models.RefInventory, error) {
	tower, err := s.Daos.GetSingleInventory(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

// UpdateInventory : ""
func (s *Service) UpdateInventory(ctx *models.Context, crop *models.Inventory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateInventory(ctx, crop)
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

// UpdateInventoryQuantityDetails : ""
func (s *Service) UpdateInventoryQuantityDetails(ctx *models.Context, crop *models.Inventory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		tower, err := s.Daos.GetSingleInventory(ctx, crop.UniqueID)
		if err != nil {
			return err
		}
		newquan := crop.Quantity
		crop.Quantity = newquan + tower.Quantity

		err = s.Daos.UpdateInventoryQuantityDetails(ctx, crop)
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

// EnableInventory : ""
func (s *Service) EnableInventory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableInventory(ctx, UniqueID)
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

//DisableInventory : ""
func (s *Service) DisableInventory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableInventory(ctx, UniqueID)
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

//DeleteInventory : ""
func (s *Service) DeleteInventory(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteInventory(ctx, UniqueID)
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

// FilterInventory : ""
func (s *Service) FilterInventory(ctx *models.Context, filter *models.InventoryFilter, pagination *models.Pagination) ([]models.RefInventory, error) {
	productFilter := new(models.ProductFilter)
	var productIds []string
	if len(filter.CategoryID) > 0 {
		productFilter.CategoryID = append(productFilter.CategoryID, filter.CategoryID...)
	}
	if len(filter.SubCategoryID) > 0 {
		productFilter.SubCategoryID = append(productFilter.SubCategoryID, filter.SubCategoryID...)
	}
	res, err := s.Daos.FilterProduct(ctx, productFilter, nil)
	if err != nil {
		return nil, err
	}
	if len(res) > 0 {
		for _, v := range res {
			productIds = append(productIds, v.UniqueID)
		}
		filter.ProductID = append(filter.ProductID, productIds...)
	} else {
		return make([]models.RefInventory, 0), err
	}
	log.Println("products", filter.ProductID)
	return s.Daos.FilterInventory(ctx, filter, pagination)

}

// CreateMesh : ""
// func (s *Service) CreateMesh(ctx *models.Context, imc *models.InventoryMeshCreate) (interface{}, error) {

// 	productVarients := []models.ProductVariant{}
// 	if imc == nil {
// 		return nil, errors.New("imc is nill")
// 	}
// 	for k, v := range imc.ProductVarient {
// 		for _, v2 := range v {
// 			var pv models.ProductVariant
// 			pv.VariantTypeID = k
// 			pv.Name = v2
// 			productVarients = append(productVarients, pv)

// 		}
// 	}
// 	fmt.Println("Product Varients\n", productVarients)
// 	var types []string
// 	for k, _ := range imc.ProductVarient {
// 		types = append(types, k)

// 	}
// 	for i := 0; i < len(types); i++ {
// 		for _, v := range productVarients {

// 			if v.VariantTypeID == types[i] {

// 				for _, v1 := range productVarients {
// 					if v1.VariantTypeID != v.VariantTypeID {
// 						fmt.Println(v1.Name)
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return types, nil
// }

// ImageInventory : ""
func (s *Service) ImageInventory(ctx *models.Context, crop *models.Inventory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.ImageInventory(ctx, crop)
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

// ImageInventory : ""
func (s *Service) ImagesInventory(ctx *models.Context, crop *models.Inventory) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.ImagesInventory(ctx, crop)
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

//GetbyBarcodeAndVendor :""
func (s *Service) GetbyBarcodeAndVendor(ctx *models.Context, UniqueID string, Vendor string) (*models.RefInventory, error) {
	tower, err := s.Daos.GetbyBarcodeAndVendor(ctx, UniqueID, Vendor)
	if err != nil {
		return nil, err
	}
	return tower, nil
}
func (s *Service) ChkUniqueness(ctx *models.Context, Barcode string, Vendor string) (*models.RefInventory, error) {
	tower, err := s.Daos.ChkUniqueness(ctx, Barcode, Vendor)
	if err != nil {
		return nil, err
	}
	return tower, nil
}
