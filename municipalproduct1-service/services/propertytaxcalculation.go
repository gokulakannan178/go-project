package services

import (
	"municipalproduct1-service/models"
)

//GetPropertyTaxCalculation :""
func (s *Service) GetPropertyTaxCalculation(ctx *models.Context, UniqueID string) ([]models.PropertyTaxCalculation, error) {
	propertTaxCalculations, err := s.Daos.GetPropertyTaxCalculation(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return propertTaxCalculations, nil
}
