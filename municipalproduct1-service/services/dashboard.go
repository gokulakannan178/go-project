package services

import "municipalproduct1-service/models"

func (s *Service) DashboardDemandAndCollection(ctx *models.Context, filter *models.DashboardDemandAndCollectionFilter) (*models.DashboardDemandAndCollection, error) {
	//defer ctx.Session.EndSession(ctx.CTX)

	data, err := s.Daos.DashboardDemandAndCollection(ctx, filter)
	if err != nil {
		return nil, err
	}
	if data != nil {
		// data.TotalCollectionCurrent = data.TotalCollectionCurrent -
		// data.TotalCollectionTax = data.TotalCollectionTax -
	}
	return data, nil
}

func (s *Service) DashboardDemandAndCollectionV2(ctx *models.Context, filter *models.DashboardDemandAndCollectionFilter) (*models.DashboardDemandAndCollection, error) {
	//defer ctx.Session.EndSession(ctx.CTX)

	data, err := s.Daos.DashboardDemandAndCollectionV2(ctx, filter)
	if err != nil {
		return nil, err
	}
	return data, nil
}
