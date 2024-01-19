package services

import "municipalproduct1-service/models"

// BasicShopRentUpdateGetPaymentsToBeUpdated : ""
func (s *Service) BasicShopRentUpdateGetPaymentsToBeUpdated(ctx *models.Context, rbsrul *models.RefBasicShopRentUpdateLog) ([]models.RefShopRentPayments, error) {
	return s.Daos.BasicShopRentUpdateGetPaymentsToBeUpdated(ctx, rbsrul)
}
