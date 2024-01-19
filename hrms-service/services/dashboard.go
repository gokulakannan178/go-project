package services

import "hrms-services/models"

func (s *Service) GetCollectionCount(ctx *models.Context, UniqueID string) (*models.Dashboard, error) {
	Dashboard, err := s.Daos.GetCollectionCount(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Dashboard, nil
}
