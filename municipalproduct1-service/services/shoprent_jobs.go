package services

import (
	"fmt"
	"municipalproduct1-service/models"
)

func (s *Service) UpdateOverallDemandForAll(ctx *models.Context, status []string) error {
	data, err := s.Daos.GetShopRentForOverAllDemand(ctx, status)
	if err != nil {
		return err
	}
	boolValue := true
	for _, v := range data {
		fmt.Println(v.UniqueID)
		res, err := s.CalcShopRentOverallMonthlyDemand(ctx, v.UniqueID, boolValue)
		if err != nil {
			return err
		}
		fmt.Println(res)
	}

	return nil
}
