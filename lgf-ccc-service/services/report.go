package service

import "lgf-ccc-service/models"

func (s *Service) DayWiseDumphistoryCount(ctx *models.Context, filter *models.FilterDumpHistory) ([]models.MonthWiseDumphistoryCount, error) {
	Report, err := s.Daos.DayWiseDumphistoryCount(ctx, filter)
	if err != nil {
		return nil, err
	}
	return Report, nil
}
func (s *Service) MonthWiseDumphistoryCount(ctx *models.Context, filter *models.FilterDumpHistory) ([]models.MonthWiseDumphistoryCount, error) {
	Report, err := s.Daos.MonthWiseDumphistoryCount(ctx, filter)
	if err != nil {
		return nil, err
	}
	return Report, nil
}
func (s *Service) CircleWiseHouseVisitedCount(ctx *models.Context, filter *models.FilterHouseVisited) ([]models.CircleWiseHouseVisitedv2, error) {
	Report, err := s.Daos.CircleWiseHouseVisitedCount(ctx, filter)
	if err != nil {
		return nil, err
	}
	return Report, nil
}
func (s *Service) DayWiseWardHouseVisitedCount(ctx *models.Context, filter *models.FilterHouseVisited) ([]models.CircleWiseHouseVisitedv2, error) {
	Report, err := s.Daos.DayWiseWardHouseVisitedCount(ctx, filter)
	if err != nil {
		return nil, err
	}
	return Report, nil
}
