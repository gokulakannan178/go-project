package daos

import (
	"errors"
	"fmt"
	"haritv2-service/constants"
	"haritv2-service/models"
)

// SaveFPOINVENTORY : ""
func (d *Daos) SaveBlockChain(ctx *models.Context, blockChain *models.Inventory) error {
	data := make(map[string]string)
	data["id"] = blockChain.ID
	fmt.Println("data===========", data)
	_, err := d.Shared.Post("http://localhost:9000/api/blockchain", data, blockChain)
	if err != nil {
		return errors.New("please check the blockchain")
	}
	_, err = ctx.DB.Collection(constants.COLLECTIONBLOCKCHAIN).InsertOne(ctx.CTX, blockChain)
	return err
}
