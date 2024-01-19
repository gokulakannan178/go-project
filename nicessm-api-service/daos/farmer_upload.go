package daos

import (
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"time"
)

// FarmerLandUploadExcel :""
func (d *Daos) FarmerLandUploadExcel(ctx *models.Context, farmerLands []models.FarmerLand) error {
	var temFarmerLand []interface{}
	for _, v := range farmerLands {
		t := time.Now()
		v.Created.On = &t
		v.Created.By = constants.SYSTEM
		temFarmerLand = append(temFarmerLand, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMERLAND).InsertMany(ctx.CTX, temFarmerLand)
	return err
}

// FarmerAggregationUploadExcel :""
func (d *Daos) FarmerAggregationUploadExcel(ctx *models.Context, farmers []models.Farmer) error {
	var temFarmer []interface{}
	for _, v := range farmers {
		temFarmer = append(temFarmer, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONFARMER).InsertMany(ctx.CTX, temFarmer)
	return err
}
