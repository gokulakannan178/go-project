package services

import (
	"nicessm-api-service/models"
)

//GetSingleOnePageAttachment :""
func (s *Service) GetSingleOnePageAttachment(ctx *models.Context, UniqueID string) (string, error) {
	OnePageAttachment, err := s.Daos.GetSingleOnePageAttachment(ctx, UniqueID)
	if err != nil {
		return "", err
	}
	return OnePageAttachment, nil
}
