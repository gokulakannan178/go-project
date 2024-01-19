package services

import (
	"logikoof-echalan-service/models"
)

//PaymentWidget : ""
func (s *Service) PaymentWidget(ctx *models.Context, filter *models.PaymentWidgetFilter) (*models.PaymentWidget, error) {
	return s.Daos.PaymentWidget(ctx, filter)
}

//TodaysOffenceWidget : ""
func (s *Service) TodaysOffenceWidget(ctx *models.Context, filter *models.TodaysOffenceWidgetFilter) (*models.TodaysOffenceWidget, error) {
	return s.Daos.TodaysOffenceWidget(ctx, filter)
}

//TopOffencesWidget : ""
func (s *Service) TopOffencesWidget(ctx *models.Context, filter *models.TopOffencesWidgetFilter) ([]models.TopOffencesWidget, error) {
	return s.Daos.TopOffencesWidget(ctx, filter)
}
