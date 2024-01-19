package services

import (
	"haritv2-service/models"
)

//ULBNearBy : ""
func (s *Service) ULBNearBy(ctx *models.Context, ulbnb *models.ULBNearBy, pagination *models.Pagination) ([]models.RefULB, error) {

	ulbs, err := s.Daos.ULBNearBy(ctx, ulbnb, pagination)
	if err != nil {
		return nil, err
	}

	return ulbs, nil
}

//UlbInTheState : ""
func (s *Service) UlbInTheState(ctx *models.Context, ulbnb string, pagination *models.Pagination) ([]models.RefULB, error) {

	ulbs, err := s.Daos.UlbInTheState(ctx, ulbnb, pagination)
	if err != nil {
		return nil, err
	}

	return ulbs, nil
}

//UlbInTheState : ""
func (s *Service) UlbInTheStateV2(ctx *models.Context, ulbnb string, sortBy string, sortorder int, pagination *models.Pagination) ([]models.ULBNearByResponse, error) {

	ulbs, err := s.Daos.UlbInTheStateV2(ctx, ulbnb, sortBy, sortorder, pagination)
	if err != nil {
		return nil, err
	}

	return ulbs, nil
}

//UlbInTheStateV3 : ""
func (s *Service) UlbInTheStateV3(ctx *models.Context, ULBStateIn *models.ULBStateIn, pagination *models.Pagination) ([]models.ULBNearByResponse, error) {

	ulbs, err := s.Daos.UlbInTheStateV3(ctx, ULBStateIn, pagination)
	if err != nil {
		return nil, err
	}

	return ulbs, nil
}

//UlbCompostInTheState : ""
func (s *Service) UlbCompostInTheState(ctx *models.Context, ulbnb string) (*models.CompostInStock, error) {

	ulbs, err := s.Daos.UlbCompostInTheState(ctx, ulbnb)
	if err != nil {
		return nil, err
	}

	return ulbs, nil
}
