package daos

import (
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
)

//SavePropertyFloors :""
func (d *Daos) SaveEstimatedFloors(ctx *models.Context, propertyFloors []models.PropertyFloor) error {
	insertdata := []interface{}{}
	for _, v := range propertyFloors {
		insertdata = append(insertdata, v)
	}
	result, err := ctx.DB.Collection(constants.COLLECTIONESTIMATEDPROPERTYFLOOR).InsertMany(ctx.SC, insertdata)
	fmt.Println("insert result =>", result)
	return err
}
