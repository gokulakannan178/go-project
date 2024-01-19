package service

import "lgf-ccc-service/models"

func (s *Service) GetCollectionCount(ctx *models.Context, UniqueID string) (*models.Dashboard, error) {
	Dashboard, err := s.Daos.GetCollectionCount(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Dashboard, nil
}
func (s *Service) GetDumbSiteCount(ctx *models.Context, UniqueID string) (*models.DumbSiteCount, error) {
	Dashboard, err := s.Daos.GetDumbSiteCount(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Dashboard, nil
}
func (s *Service) GetHousevisitedCount(ctx *models.Context, UniqueID string) (*models.HousevisitedCount, error) {
	Dashboard, err := s.Daos.GetHousevisitedCount(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Dashboard, nil
}
func (s *Service) GetvehicleCount(ctx *models.Context, UniqueID string) (*models.Dashboard, error) {
	Dashboard, err := s.Daos.GetvehicleCount(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Dashboard, nil
}
func (s *Service) GetUsertypeCount(ctx *models.Context, UniqueID string) (*models.UserTypeCount, error) {
	Dashboard, err := s.Daos.GetUsertypeCount(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return Dashboard, nil
}

func (s *Service) GetPropertyCount(ctx *models.Context, filter *models.FilterProperties) (*models.PropertyCount, error) {
	return s.Daos.GetPropertyCount(ctx, filter)
}

func (s *Service) GetGarbaggeCount(ctx *models.Context, filter *models.FilterHouseVisited) (*models.PropertyCount, error) {
	return s.Daos.GetGarbaggeCount(ctx, filter)
}
