package services

import (
	"nicessm-api-service/models"
)

func (s *Service) ContentManagerCount(ctx *models.Context, content *models.ContentFilter) ([]models.ContentCount, error) {

	//log.Println("showcontentcount")
	//Start Transaction
	err := s.ContentDataAccess(ctx, content)
	if err != nil {
		return nil, err
	}
	return s.Daos.ContentManagerCount(ctx, content)

}
