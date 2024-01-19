package daos

import (
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
)

func (d *Daos) LoadBlockWeatherReport(ctx *models.Context) error {
	var lbwr []models.LoadBlockWeatherReport
	var Weather []interface{}
	for _, v := range lbwr {
		Weather = append(Weather, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONWEATHERDATA).InsertMany(ctx.CTX, Weather)
	return err
}
