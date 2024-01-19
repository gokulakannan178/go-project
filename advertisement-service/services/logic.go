package services

import (
	"ecommerce-service/models"
	"fmt"
	"strings"
)

//Logic :""
func (s *Service) Logic(ctx *models.Context, mlogic []models.VarientInputLogic) ([]models.VarientOutputLogic, error) {

	//fmt.Println(mlogic)
	var outlogic []models.VarientOutputLogic

	for i := 0; i < len(mlogic); i++ {
		var output models.VarientOutputLogic
		strs := strings.Split(mlogic[i].Varients, ",")
		output.VarientID = mlogic[i].VarientID
		output.VarientName = mlogic[i].VarientName
		output.Varients = append(output.Varients, strs...)
		outlogic = append(outlogic, output)
		strs = nil

	}

	fmt.Println(outlogic)

	return outlogic, nil

}
