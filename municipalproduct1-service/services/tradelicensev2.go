package services

import (
	"municipalproduct1-service/models"
)

//GetSingleTradeLicenseV2 :""
func (s *Service) GetSingleTradeLicenseV2(ctx *models.Context, UniqueID string) (*models.RefTradeLicense, error) {
	tower, err := s.Daos.GetSingleTradeLicenseV2(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}

//GetSingleTradeLicensePDF :""
func (s *Service) GetSingleTradeLicensePDF(ctx *models.Context, UniqueID string) (*models.RefTradeLicense, error) {
	tower, err := s.Daos.GetSingleTradeLicenseV2(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return tower, nil
}
