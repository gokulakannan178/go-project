package services

import "municipalproduct1-service/models"

// FilterPropertyOtherDemandPartPayment :""
func (s *Service) FilterPropertyOtherDemandPartPayment(ctx *models.Context, filter *models.PropertyOtherDemandPartPaymentFilter, pagination *models.Pagination) ([]models.RefPropertyOtherDemandPartPayment, error) {
	defer ctx.Session.EndSession(ctx.CTX)
	return s.Daos.FilterPropertyOtherDemandPartPayment(ctx, filter, pagination)
}
