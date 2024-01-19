package daos

import (
	"ecommerce-service/constants"
	"ecommerce-service/models"
)

// SaveManyInventoryMesh : ""
func (d *Daos) SaveManyInventoryMesh(ctx *models.Context, inMesh []models.InventoryMesh) error {
	d.Shared.BsonToJSONPrint(inMesh)
	data := []interface{}{}
	for _, v := range inMesh {
		data = append(data, v)
	}
	_, err := ctx.DB.Collection(constants.COLLECTIONINVENTORYMESH).InsertMany(ctx.CTX, data)
	return err
}
