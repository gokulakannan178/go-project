package services

import "municipalproduct1-service/models"

// BasicSolidWasteUpdateGetPaymentsToBeUpdated : ""
func (s *Service) BasicSolidWasteUpdateGetPaymentsToBeUpdated(ctx *models.Context, rbsrul *models.RefBasicSolidWasteUpdateLog) ([]models.RefSolidWasteChargeMonthlyPayments, error) {
	return s.Daos.BasicSolidWasteUpdateGetPaymentsToBeUpdated(ctx, rbsrul)
}
